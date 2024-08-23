package main

import (
	"context"
	"fmt"
	"log"
	
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/google/generative-ai-go/genai"
	"github.com/group13/blog/config"
	"github.com/group13/blog/delivery/common"
	blogcontroller "github.com/group13/blog/delivery/controller/blog"
	"github.com/group13/blog/delivery/controller/gemini"
	usercontroller "github.com/group13/blog/delivery/controller/user"
	reactioncontroller "github.com/group13/blog/delivery/controller/reaction"
	"github.com/group13/blog/delivery/router"
	cache "github.com/group13/blog/infrastructure/cache"
	db "github.com/group13/blog/infrastructure/database"
	"github.com/group13/blog/infrastructure/email"
	"github.com/group13/blog/infrastructure/hash"
	"github.com/group13/blog/infrastructure/jwt"
	blogrepo "github.com/group13/blog/infrastructure/repo/blog"
	commentrepo "github.com/group13/blog/infrastructure/repo/comment"
	reactionrepo "github.com/group13/blog/infrastructure/repo/reaction"
	userrepo "github.com/group13/blog/infrastructure/repo/user"
	blogcmd "github.com/group13/blog/usecase/blog/command"
	blogqry "github.com/group13/blog/usecase/blog/query"
	passwordreset "github.com/group13/blog/usecase/password_reset"
	geminiService "github.com/group13/blog/usecase/ai_recommendation/query"
	usercmd "github.com/group13/blog/usecase/user/command"
	reactioncmd "github.com/group13/blog/usecase/reaction/command"
	userqry "github.com/group13/blog/usecase/user/query"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/option"
)

// main is the entry point for the application.
// It initializes the MongoDB client, services, controllers, and starts the HTTP server.
func main() {
	cfg := config.Envs

	// Initialize MongoDB client and perform migrations
	mongoClient := initDB(cfg)
	cacheClient := initCache(cfg)
	geminiModel := initGeminiClient(cfg)

	// Initialize services
	userRepo, blogRepo, _, reactionRepo := initRepos(cfg, mongoClient)
	jwtService := jwt.New(
		jwt.Config{
			SecretKey: config.Envs.JWTSecret,
			Issuer:    config.Envs.ServerHost,
			ExpTime:   config.Envs.JWTExpirationInSeconds,
		})
	hashService := &hash.Service{}
	emailService := &email.MailTrapService{}

	// init controllers
	
	reactionController := initReactionController(reactionRepo, blogRepo)
	userController := initUserController(userRepo, hashService, jwtService, emailService)
	blogController := initBlogController(blogRepo, cacheClient)
	geminiController := initGeminiController(geminiService.NewReccomendationHandler(geminiModel))
	
	// Router configuration
	routerConfig := router.Config{
		Addr:        fmt.Sprintf(":%s", cfg.ServerPort),
		BaseURL:     "/api",
		Controllers: []common.IController{userController, blogController,geminiController, reactionController},
		JwtService:  jwtService,
	}
	r := router.NewRouter(routerConfig)

	// Start the server
	if err := r.Run(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// initDB initializes the MongoDB client and performs any necessary database migrations.
// It returns the MongoDB client instance.
func initDB(cfg config.Config) *mongo.Client {
	mongoClient := db.Connect(db.Config{
		ConnectString: cfg.DBConnectionString,
	})

	db.Migrate(mongoClient, cfg.DBName)

	return mongoClient
}

// initServices initializes the necessary services for the application.
// It returns the user repository, task repository, JWT service, and hash service.
func initRepos(cfg config.Config, mongoClient *mongo.Client) (*userrepo.Repo, *blogrepo.Repo, *commentrepo.Repo, *reactionrepo.Repo) {
	userRepo := userrepo.NewRepo(mongoClient, cfg.DBName, "users")
	blogRepo := blogrepo.New(mongoClient, cfg.DBName, "blogs")
	commentRepo := commentrepo.New(mongoClient, cfg.DBName, "comments")
	reactionRepo := reactionrepo.New(mongoClient, cfg.DBName, "reactions")

	return userRepo, blogRepo, commentRepo, reactionRepo
}

func initCache(cfg config.Config) *cache.RedisCache {
	host, err := strconv.ParseInt(cfg.Cache_port, 10, 32)
	if err != nil {
		log.Fatalf("Error parsing cache db: %v", err)
	}

	// Initialize the cache
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Cache_host, cfg.Cache_port),
		Password: "",
		DB:       int(host),
	})

	redisClient := cache.NewRedisCache(client, cfg.Blog_cache_expiry)
	return redisClient
}


func initGeminiClient(cfg config.Config) *genai.GenerativeModel{
	ctx := context.Background()
	key := cfg.Google_Api_Key
	if key == "" {
		log.Fatalf("Error: Google API Key not found")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(key))
	
	if err != nil {
		log.Printf("Error Gemini client not created: %v", err)
	}

	model := client.GenerativeModel("gemini-1.5-pro-latest")
	model.SetTemperature(0.9)
	model.SetTopP(0.5)
	model.SetTopK(20)
	model.SetMaxOutputTokens(100)
	return model


}
func initUserController(userRepo *userrepo.Repo, hashService *hash.Service, jwtService *jwt.Service, mailService *email.MailTrapService,) *usercontroller.UserController {
	promoteHandler := usercmd.NewPromoteHandler(userRepo)
	loginHandler := userqry.NewLoginHandler(userqry.LoginConfig{
		UserRepo:     userRepo,
		HashService:  hashService,
		JwtService:   jwtService,
		EmailService: mailService,
	})
	signupHandler := usercmd.NewSignUpHandler(usercmd.SignUpConfig{
		UserRepo:     userRepo,
		HashService:  hashService,
		JwtService:   jwtService,
		EmailService: mailService,
	})
	updateProfileHandler := usercmd.NewUpdateProfileHandler(userRepo, hashService, mailService)
	// aiController := gemini.NewAiController(geminiHandler)
	
	resetPasswordHandler := passwordreset.NewResetHandler(userRepo, hashService, jwtService)
	resetCodeSendHandler := passwordreset.NewSendcodeHandler(userRepo, mailService, hashService)
	validateCodeHandler := passwordreset.NewValidateCodeHandler(userRepo, jwtService, hashService)
	validateEmailHandler := usercmd.NewValidateEmailHandler(userRepo, hashService, jwtService)

	return usercontroller.New(usercontroller.Config{
		PromoteHandler:       promoteHandler,
		LoginHandler:         loginHandler,
		SignupHandler:        signupHandler,
		ResetPasswordHandler: resetPasswordHandler,
		ResetCodeSendHandler: resetCodeSendHandler,
		ValidateCodeHandler:  validateCodeHandler,
		ValidateEmailHandler: validateEmailHandler,
		UpdateProfileHandler: updateProfileHandler,
	})

}

func initBlogController(blogRepo *blogrepo.Repo, cacheService *cache.RedisCache) *blogcontroller.Controller {
	addHandler := blogcmd.NewAddHandler(blogRepo)
	updateHandler := blogcmd.NewUpdateHandler(blogRepo)
	deleteHandler := blogcmd.NewDeleteHandler(blogRepo)
	getMultipleHandler := blogqry.NewGetMultipleHandler(blogRepo, cacheService)
	getHandler := blogqry.NewGetHandler(blogRepo)

	return blogcontroller.New(blogcontroller.Config{
		AddHandler:         addHandler,
		UpdateHandler:      updateHandler,
		DeleteHandler:      deleteHandler,
		GetMultipleHandler: getMultipleHandler,
		GetHandler:         getHandler,
	})
}

func initGeminiController(geminiHandler *geminiService.RecomendationHandler) *gemini.Controller {
	return gemini.NewAiController(geminiHandler)
}

func initReactionController(reactionRepo *reactionrepo.Repo ,blogRepo  *blogrepo.Repo) *reactioncontroller.ReactionController {
	updateHandler := reactioncmd.NewUpdateHandler(reactionRepo,blogRepo)
	deleteHandler := reactioncmd.NewDeleteHandler(blogRepo, reactionRepo)

	return reactioncontroller.New(reactioncontroller.Config{
		UpdateReactionHandler: updateHandler,
		DeleteReactionHandler: deleteHandler,
	})
}



