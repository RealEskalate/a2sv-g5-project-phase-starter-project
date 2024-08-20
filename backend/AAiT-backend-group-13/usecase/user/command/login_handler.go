package usercmd

import (
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
var _ icmd.IHandler[*LoginCommand, *result.LoginInResult] = &LoginHandler{}

// Handle processes the login command and returns the login result with tokens.
func (h *LoginHandler) Handle(command *LoginCommand) (*result.LoginInResult, error) {
	// Find user by username
	user, err := h.repo.FindByUsername(command.username)
	if err != nil {
		return nil, er.NewNotFound("user not found")
	}

	// Verify password
	ok, err := h.hashService.Match(command.password, user.PasswordHash())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, er.NewValidation("password is incorrect")
	}

	// Mark user as active
	user.MakeActive()

	// Generate tokens
	token, err := h.jwtService.Generate(user, ijwt.Access)
	if err != nil {
		return nil, err
	}

	refreshToken, err := h.jwtService.Generate(user, ijwt.Refresh)
	if err != nil {
		return nil, err
	}

	// Return login result
	return &result.LoginInResult{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
