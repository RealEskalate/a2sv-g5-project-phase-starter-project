package passwordreset

import (
	"log"

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
	claims, err := h.jwtService.Decode(cmd.token)
	if err != nil {
		return false, er.NewUnauthorized(err.Error())
	}

	log.Println(claims)
	if isForReset, ok := claims["is_for_reset"].(bool); !ok || !isForReset {
		return false, er.NewUnauthorized("invalid token")
	}

	email, ok := claims["email"].(string)
	log.Println(email, ok)
	if !ok {
		return false, er.NewUnauthorized("invalid token")
	}

	user, err := h.userRepo.FindByEmail(email)
	log.Println(user)
	if err != nil {
		return false, er.NewUnauthorized(err.Error())
	}

	if err := user.UpdatePassword(cmd.NewPassword, h.hashService); err != nil {
		return false, er.NewUnauthorized(err.Error())
	}

	if err := h.userRepo.Save(user); err != nil {
		return false, er.NewUnauthorized(err.Error())
	}

	return true, nil
}
