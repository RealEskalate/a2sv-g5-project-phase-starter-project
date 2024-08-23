package usercmd

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
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
		repo:         repo,
		hashService:  hashService,
		emailService: emailService,
	}
}

func (h *UpdateProfileHandler) Handle(command *UpdateProfileCommand) (*result.UpdateProfileResult, error) {
	log.Println("Handling UpdateProfileCommand")

	emailUpdated := false 
	// Parse the user ID from the command
	userID, err := uuid.Parse(command.userid)
	if err != nil {
		log.Printf("Error parsing user ID: %v", err)
		return &result.UpdateProfileResult{}, er.NewUnexpected("failed to parse user ID")
	}

	// Fetch the user by ID
	user, err := h.repo.FindById(userID)
	if err != nil {
		log.Printf("Error finding user by ID: %v", err)
		return &result.UpdateProfileResult{}, err
	}

	if user == nil {
		log.Println("User not found")
		return &result.UpdateProfileResult{}, er.UserNotFound
	}

	// Check if the new email is available
	if command.Email != "" {
		log.Println("Updating email")
		user.UpdateEmail(command.Email)
		emailUpdated = true
		// Generate a validation link
		
	} 

	// Check if the new username is available
	if command.Username != "" {
		log.Println("Updating username")
		user.UpdateUsername(command.Username)
	}

	// Update other fields
	if command.FirstName != "" {
		log.Println("Updating first name")
		user.UpdateFirstName(command.FirstName)
	} 
	

	if command.LastName != "" {
		log.Println("Updating last name")
		user.UpdateLastName(command.LastName)
	}

	if command.Password != "" {
		log.Println("Updating password")
		user.UpdatePassword(command.Password, h.hashService)
	}

	// Save the updated user profile
	err = h.repo.Save(user)
	if err != nil {
		log.Printf("Error saving user: %v", err)
		return &result.UpdateProfileResult{}, err
	}
	
	if emailUpdated{
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

	log.Println("Profile updated successfully")
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
