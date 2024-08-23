package userqry

import (
	"log"

	er "github.com/group13/blog/domain/errors"
	ihash "github.com/group13/blog/domain/i_hash"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	iemail "github.com/group13/blog/usecase/common/i_email"
	ijwt "github.com/group13/blog/usecase/common/i_jwt"
	irepo "github.com/group13/blog/usecase/common/i_repo"
	result "github.com/group13/blog/usecase/user/result"
)

// LoginHandler handles user login requests.
type LoginHandler struct {
	repo         irepo.UserRepository
	jwtService   ijwt.Service
	hashService  ihash.Service
	emailService iemail.Service
}

// LoginConfig holds the configuration for creating a LoginHandler.
type LoginConfig struct {
	UserRepo     irepo.UserRepository
	JwtService   ijwt.Service
	HashService  ihash.Service
	EmailService iemail.Service
}

// NewLoginHandler creates a new instance of LoginHandler with the provided configuration.
func NewLoginHandler(config LoginConfig) *LoginHandler {
	return &LoginHandler{
		repo:         config.UserRepo,
		jwtService:   config.JwtService,
		hashService:  config.HashService,
		emailService: config.EmailService,
	}
}

// Ensure LoginHandler implements icmd.IHandler
var _ icmd.IHandler[*LoginQuery, *result.LoginInResult] = &LoginHandler{}

// Handle processes the login command and returns the login result with tokens.
func (h *LoginHandler) Handle(command *LoginQuery) (*result.LoginInResult, error) {
	log.Printf("Finding user by username: %s -- LoginHandler", command.username)
	user, err := h.repo.FindByUsername(command.username)
	if  err != nil {
		return nil, er.NewUnauthorized("user not found")
	} else if !user.IsActive() {
		return nil, er.NewUnauthorized("user is not activated")
	}

	log.Println(user, "found this user", user.PasswordHash(), "now checking password")
	ok, err := h.hashService.Match(user.PasswordHash(), command.password)
	if err != nil {
		log.Printf("Error matching password %v", err)
		return nil, err
	}
	if !ok {
		log.Println("password didnot match")
		return nil, er.NewUnauthorized("password is incorrect")
	}

	token, err := h.jwtService.Generate(user, ijwt.Access)
	if err != nil {
		return nil, err
	}

	refreshToken, err := h.jwtService.Generate(user, ijwt.Refresh)
	if err != nil {
		return nil, err
	}

	return result.NewLoginInResult(user, token, refreshToken), nil
}
