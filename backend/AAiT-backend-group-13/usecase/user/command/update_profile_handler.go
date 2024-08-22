package usercmd

import (
	"fmt"
	"log"
	"time"

	er "github.com/group13/blog/domain/errors"
	ihash "github.com/group13/blog/domain/i_hash"
	"github.com/group13/blog/domain/models"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	iemail "github.com/group13/blog/usecase/common/i_email"
	irepo "github.com/group13/blog/usecase/common/i_repo"
	result "github.com/group13/blog/usecase/user/result"
)

// UpdateProfileHandler handles user profile update logic.
type UpdateProfileHandler struct {
	repo         irepo.UserRepository
	hashService  ihash.Service
	emailService iemail.Service
}


var _ icmd.IHandler[*UpdateProfileCommand, *result.UpdateProfileResult] = &UpdateProfileHandler{}

// NewUpdateProfileHandler creates a new UpdateProfileHandler instance.
func NewUpdateProfileHandler(repo irepo.UserRepository, hashService ihash.Service, emailService iemail.Service) *UpdateProfileHandler {
	return &UpdateProfileHandler{
		repo: repo,
		hashService: hashService,
		emailService: emailService,
	}
}

// Handle handles the UpdateProfileCommand and returns the result.
func (h *UpdateProfileHandler) Handle(command *UpdateProfileCommand) (*result.UpdateProfileResult, error) {
	user, err := h.repo.FindByUsername(command.Username)
	if err != nil {
		return &result.UpdateProfileResult{}, err
	}

	if user == nil {
		return &result.UpdateProfileResult{}, er.UserNotFound
	}

	if user.Email() != command.Email {
		return &result.UpdateProfileResult{}, er.NewConflict("Another acc exists with this email Please use the other account")
	}

	user, err = h.repo.FindByEmail(command.Email)
	if err != nil {
		return &result.UpdateProfileResult{}, err
	}
	
	if user == nil {
		return &result.UpdateProfileResult{}, er.UserNotFound
	}

	if user.Username() != command.Username {
		return &result.UpdateProfileResult{}, er.NewConflict("Username taken")
	}
	if command.FirstName != ""{
		user.UpdateFirstName(command.FirstName)
	}

	if command.LastName != ""{
		user.UpdateLastName(command.LastName)
	}

	if command.Password != ""{
		user.UpdatePassword(command.Password, h.hashService)
	}
	if command.Username != ""{
		user.UpdateUsername(command.Username)
	}
	if command.Email != ""{

		user.UpdateEmail(command.Email)
		// Generate a validation link
		validationLink, err := h.generateValidationLink(*user)
		if err != nil {
			log.Printf("Error generating validation link: %v", err)
			return &result.UpdateProfileResult{}, err
		}
		log.Println("Validation link generated")

		// Send the sign-up email
		mails := []string{user.Email()}
		mail := iemail.NewSignUpEmail("", mails, validationLink)
		if err := h.emailService.Send(mail); err != nil {
			log.Printf("Error sending sign-up email: %v", err)
			return &result.UpdateProfileResult{}, err
		}
		log.Println("Sign-up email sent")
	}
	user.MakeInactive()
	
	err  = h.repo.Save(user)
	

	if err != nil {
		return &result.UpdateProfileResult{}, err
	}
	message := "Profile updated successfully"
	res := result.NewUpdateProfileResult(message)
	return &res, nil
}


func (h *UpdateProfileHandler) generateValidationLink(user models.User) (string, error) {
	log.Printf("Generating validation link for user %s", user.Username())

	userID := user.ID().String()
	expiryDay := time.Now().Add(time.Hour * 24).Format(time.RFC3339)
	username := user.Username()
	value := userID + "|" + expiryDay + "|" + username

	encryptedValue, err := encrypt(value)
	if err != nil {
		log.Printf("Error encrypting validation link: %v", err)
		return "", er.NewUnexpected("failed to encrypt value")
	}

	validationLink := fmt.Sprintf("http://localhost:8080/api/v1/auth/validateEmail?secret?=%s", encryptedValue)
	log.Printf("Validation link generated: %s", validationLink)
	return validationLink, nil
}






