package usercommand

import (
	"strings"
	"time"

	"github.com/google/uuid"
	er "github.com/group13/blog/domain/errors"
	ihash "github.com/group13/blog/domain/i_hash"
	result "github.com/group13/blog/usecases_sof/user/result"
	irepository "github.com/group13/blog/usecases_sof/utils/i_repo"
	ijwt "github.com/group13/blog/usecases_sof/utils/i_jwt"
)



type ValidateEmailHandler struct {
	repo         irepository.UserRepository
	hashService 	  ihash.Service
	jwtService 			ijwt.Services

}


func (h *ValidateEmailHandler) Handle(encryptedValue string) (*result.ValidateEmailResult,  error) {
	
	decodedUsername, decodedExpiry, decodeduserId := h.decodesecret(encryptedValue)
	userID, err := uuid.Parse(decodeduserId)
	if err != nil {
		return nil, er.NewBadRequest("invalid user ID")
	}

	// Retrieve the user from the repository
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
	
	if decodedUsername != user.Username()  || expiryTime.Before(time.Now()) {
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
		Token: token,
		Refreshtoken: refreshtoken,		
	}, nil
}

func(h *ValidateEmailHandler)  decodesecret(secret string) (string, string, string) {
	// Decode the secret value
	decoded, err := h.hashService.Decode(secret)
	if err != nil {
		return "", "", ""
	}

	// Split the decoded value
	values := strings.Split(decoded, "|")
	if len(values) != 3 {
		return "", "", ""
	}

	return values[0], values[1], values[2]
}
