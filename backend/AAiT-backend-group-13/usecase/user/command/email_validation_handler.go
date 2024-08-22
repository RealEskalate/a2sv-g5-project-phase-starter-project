package usercmd

import (
	"strings"
	"time"

	"github.com/google/uuid"
	er "github.com/group13/blog/domain/errors"
	ihash "github.com/group13/blog/domain/i_hash"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	ijwt "github.com/group13/blog/usecase/common/i_jwt"
	irepository "github.com/group13/blog/usecase/common/i_repo"
	result "github.com/group13/blog/usecase/user/result"
)

// ValidateEmailHandler handles the email validation process for users.
type ValidateEmailHandler struct {
	repo        irepository.UserRepository
	hashService ihash.Service
	jwtService  ijwt.Service
}

// Ensure ValidateEmailHandler implements the expected interface
var _ icmd.IHandler[string, *result.ValidateEmailResult] = &ValidateEmailHandler{}

// NewValidateEmailHandler return pointer to a newly created ValidateEmailHandler.
func NewValidateEmailHandler(repo irepository.UserRepository, hashService ihash.Service, jwtService ijwt.Service) *ValidateEmailHandler {
	return &ValidateEmailHandler{
		repo:        repo,
		hashService: hashService,
		jwtService:  jwtService,
	}
}

// Handle processes the email validation using the provided encrypted value.
// It decodes the value, verifies the user, checks expiry, and generates JWT tokens.
func (h *ValidateEmailHandler) Handle(encryptedValue string) (*result.ValidateEmailResult, error) {
	decodedUserId, decodedExpiry, decodedUsername := h.decodesecret(encryptedValue)
	userID, err := uuid.Parse(decodedUserId)
	if err != nil {
		return nil, er.NewBadRequest("invalid user ID")
	}

	user, err := h.repo.FindById(userID)
	if err != nil {
		return nil, er.NewNotFound("user not found")
	}
	if user == nil {
		return nil, er.NewNotFound("user not found")
	}

	expiryTime, err := time.Parse(time.RFC3339, decodedExpiry)
	if err != nil {
		return nil, er.NewBadRequest("invalid expiry time")
	}
	if decodedUsername != user.Username() || expiryTime.Before(time.Now()) {
		return nil, er.NewBadRequest("invalid secret")
	}

	user.MakeActive()
	if err := h.repo.Save(user); err != nil {
		return nil, err
	}

	token, err := h.jwtService.Generate(user, ijwt.Access)
	if err != nil {
		return nil, err
	}
	refreshtoken, err := h.jwtService.Generate(user, ijwt.Refresh)
	if err != nil {
		return nil, err
	}

	return &result.ValidateEmailResult{
		Token:        token,
		Refreshtoken: refreshtoken,
	}, nil
}

// decodesecret decodes the encrypted secret to extract the username, expiry time, and user ID.
func (h *ValidateEmailHandler) decodesecret(secret string) (string, string, string) {
	decoded, err := h.hashService.Decode(secret)
	if err != nil {
		return "", "", ""
	}

	values := strings.Split(decoded, "|")
	if len(values) != 3 {
		return "", "", ""
	}

	return values[0], values[1], values[2]
}
