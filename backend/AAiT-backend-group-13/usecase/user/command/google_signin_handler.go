package usercmd

import (
	er "github.com/group13/blog/domain/errors"
	"github.com/group13/blog/infrastructure/jwt"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	ijwt "github.com/group13/blog/usecase/common/i_jwt"
	irepo "github.com/group13/blog/usecase/common/i_repo"
	result "github.com/group13/blog/usecase/user/result"
	"go.mongodb.org/mongo-driver/mongo"
)

// GoogleSignInHandler handles user sgnin using google.
type GoogleSignInHandler struct {
	repo       irepo.UserRepository
	jwtService jwt.Service
}

// LoginConfig holds the configuration for creating a LoginHandler.
type GoogleSignInConfig struct {
	UserRepo   irepo.UserRepository
	JwtService jwt.Service
}

// NewLoginHandler creates a new GoogleSignInHandler with the provided configuration.
func NewGoogleSigninHandler(config GoogleSignInConfig) *GoogleSignInHandler {
	return &GoogleSignInHandler{
		repo:       config.UserRepo,
		jwtService: config.JwtService,
	}
}

// Ensure LoginHandler implements icmd.IHandler
var _ icmd.IHandler[GoogleSigninCommand, *result.LoginInResult] = &GoogleSignInHandler{}

// Handle processes the login command and returns the login result with tokens.
func (h *GoogleSignInHandler) Handle(command GoogleSigninCommand) (*result.LoginInResult, error) {
	if !command.isVerified {
		return nil, er.NewBadRequest("email is not verified")
	}
	user, err := h.repo.FindByEmail(command.email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, er.NewNotFound("no user with given email")
		}
		return nil, err
	}

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
