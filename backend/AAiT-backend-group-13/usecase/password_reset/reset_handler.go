package passwordreset

import (
	"github.com/google/uuid"
	er "github.com/group13/blog/domain/errors"
	ihash "github.com/group13/blog/domain/i_hash"
	ijwt "github.com/group13/blog/usecase/common/i_jwt"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// ResetHandler handles the logic for resetting a user's password.
type ResetHandler struct {
	userRepo    irepo.UserRepository
	hashService ihash.Service
	jwtService  ijwt.Service
}

// NewResetHandler creates a new instance of ResetHandler with the provided user repository,
// hash service, and JWT service.
func NewResetHandler(userRepo irepo.UserRepository, hashService ihash.Service, jwtService ijwt.Service) *ResetHandler {
	return &ResetHandler{
		userRepo:    userRepo,
		hashService: hashService,
		jwtService:  jwtService,
	}
}

// Handle processes the reset password command by validating the JWT token,
// verifying the user's identity, updating their password, and saving the changes.
// It returns a boolean indicating success or an error if any step fails.
func (h *ResetHandler) Handle(cmd *ResetCommand) (bool, error) {
	// Decode the provided JWT token
	claims, err := h.jwtService.Decode(cmd.token)
	if err != nil {
		return false, er.NewUnauthorized(err.Error())
	}

	// Validate that the token is for a password reset
	if isForReset, ok := claims["is_for_reset"].(bool); !ok || !isForReset {
		return false, er.NewUnauthorized("invalid token")
	}

	// Extract user ID from claims
	id, ok := claims["id"].(uuid.UUID)
	if !ok {
		return false, er.NewUnauthorized("invalid token")
	}

	// Find user by ID
	user, err := h.userRepo.FindById(id)
	if err != nil {
		return false, er.NewUnauthorized(err.Error())
	}

	// Update the user's password
	if err := user.UpdatePassword(cmd.NewPassword, h.hashService); err != nil {
		return false, er.NewUnauthorized(err.Error())
	}

	// Save the updated user
	if err := h.userRepo.Save(user); err != nil {
		return false, er.NewUnauthorized(err.Error())
	}

	return true, nil
}
