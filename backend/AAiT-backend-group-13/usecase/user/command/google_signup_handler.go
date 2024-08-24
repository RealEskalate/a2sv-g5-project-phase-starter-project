package usercmd

import (
	er "github.com/group13/blog/domain/errors"
	"github.com/group13/blog/domain/models"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// GoogleSignInHandler handles user sgnin using google.
type GoogleSignUpHandler struct {
	repo irepo.UserRepository
}

// LoginConfig holds the configuration for creating a LoginHandler.
type GoogleSignUpConfig struct {
	UserRepo irepo.UserRepository
}

// NewLoginHandler creates a new GoogleSignInHandler with the provided configuration.
func NewGoogleSignupHandler(config GoogleSignUpConfig) *GoogleSignUpHandler {
	return &GoogleSignUpHandler{
		repo: config.UserRepo,
	}
}

// Ensure LoginHandler implements icmd.IHandler
var _ icmd.IHandler[GoogleSignupCommand, bool] = &GoogleSignUpHandler{}

// Handle processes the login command and returns the login result with tokens.
func (h *GoogleSignUpHandler) Handle(command GoogleSignupCommand) (bool, error) {
	if !command.isVerified {
		return false, er.NewBadRequest("email is not verified")
	}

	user, err := h.repo.FindByEmail(command.email)

	if user != nil {
		return false, er.NewConflict("User already exists")
	}
	if err != nil {
		return false, err
	}

	cfg := models.UserConfig{
		Email:     command.email,
		FirstName: command.firstName,
		LastName:  command.lastName,
		IsAdmin:   false,
	}

	user = models.NewFederatedUser(cfg)
	err = h.repo.Save(user)

	if err != nil {
		return false, err
	}
	return true, nil

}
