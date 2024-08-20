package main

import (
	"fmt"
	"log"
	"time"

	"github.com/group13/blog/config"
	common "github.com/group13/blog/delivery/common/icontroller"
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
	addcmd "github.com/group13/blog/usecase/blog/command/add"
	deletecmd "github.com/group13/blog/usecase/blog/command/delete"
	updatecmd "github.com/group13/blog/usecase/blog/command/update"
	blogqry "github.com/group13/blog/usecase/blog/query"
	logincommand "github.com/group13/blog/usecase/user/command/login"
	promotcmd "github.com/group13/blog/usecase/user/command/promote"
	signcommand "github.com/group13/blog/usecase/user/command/signup"
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

func initUserController(userRepo *userrepo.Repo, hashService *hash.Service, jwtService *jwt.Service, mailService *email.MailTrapService) *usercontroller.Controller {
	// 	ResetPasswordHandler  icmd.IHandler[*resetcodevalidate.Command, bool]
	// 	ForgotPasswordHandler icmd.IHandler[*forgotpassword.Command, bool]
	// 	validateEmailHander   icmd.IHandler[string, *result.ValidateEmailResult]
	//
	promotHandler := promotcmd.New(userRepo)
	loginHandler := logincommand.NewLoginHandler(logincommand.LoginConfig{
		UserRepo:     userRepo,
		HashService:  hashService,
		JwtService:   jwtService,
		EmailService: mailService,
	})
	signupHandler := signcommand.NewSignUpHandler(signcommand.SignUpConfig{
		UserRepo:     userRepo,
		HashService:  hashService,
		JwtService:   jwtService,
		EmailService: mailService,
	})

	return usercontroller.New(usercontroller.Config{
		PromotHandler: promotHandler,
		LoginHandler:  loginHandler,
		SignupHandler: signupHandler,
	})
}

func initBlogController(blogRepo *blogrepo.Repo) *blogcontroller.Controller {
	addHandler := addcmd.NewHandler(blogRepo)
	updateHandler := updatecmd.NewHandler(blogRepo)
	deleteHandler := deletecmd.New(blogRepo)
	getMultipleHandler := blogqry.NewGetMultipleHandler(blogRepo)

	return blogcontroller.New(blogcontroller.Config{
		AddHandler:         addHandler,
		UpdateHandler:      updateHandler,
		DeleteHandler:      deleteHandler,
		GetMultipleHandler: getMultipleHandler,
	})
}
