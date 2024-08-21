package main

import (
	"fmt"
	"log"
	"time"

	"github.com/group13/blog/config"
	"github.com/group13/blog/delivery/common"
	blogcontroller "github.com/group13/blog/delivery/controller/blog"
	usercontroller "github.com/group13/blog/delivery/controller/user"
	"github.com/group13/blog/delivery/router"
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
	usercmd "github.com/group13/blog/usecase/user/command"
	userqry "github.com/group13/blog/usecase/user/query"
	"go.mongodb.org/mongo-driver/mongo"
)

// main is the entry point for the application.
// It initializes the MongoDB client, services, controllers, and starts the HTTP server.
func main() {
	cfg := config.Envs

	// Initialize MongoDB client and perform migrations
	mongoClient := initDB(cfg)

	// Initialize services
	userRepo, blogRepo, _, _ := initRepos(cfg, mongoClient)
	jwtService := jwt.New(
		jwt.Config{
			SecretKey: config.Envs.JWTSecret,
			Issuer:    config.Envs.ServerHost,
			ExpTime:   time.Duration(config.Envs.JWTExpirationInSeconds) * time.Second,
		})
	hashService := &hash.Service{}
	emailService := &email.MailTrapService{}

	// init controllers
	userController := initUserController(userRepo, hashService, jwtService, emailService)
	blogController := initBlogController(blogRepo)

	// Router configuration
	routerConfig := router.Config{
		Addr:        fmt.Sprintf(":%s", cfg.ServerPort),
		BaseURL:     "/api",
		Controllers: []common.IController{userController, blogController},
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

func initUserController(userRepo *userrepo.Repo, hashService *hash.Service, jwtService *jwt.Service, mailService *email.MailTrapService) *usercontroller.UserController {
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
	})
}

func initBlogController(blogRepo *blogrepo.Repo) *blogcontroller.Controller {
	addHandler := blogcmd.NewAddHandler(blogRepo)
	updateHandler := blogcmd.NewUpdateHandler(blogRepo)
	deleteHandler := blogcmd.NewDeleteHandler(blogRepo)
	getMultipleHandler := blogqry.NewGetMultipleHandler(blogRepo)

	return blogcontroller.New(blogcontroller.Config{
		AddHandler:         addHandler,
		UpdateHandler:      updateHandler,
		DeleteHandler:      deleteHandler,
		GetMultipleHandler: getMultipleHandler,
	})
}
