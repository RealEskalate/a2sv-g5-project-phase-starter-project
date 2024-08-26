package validation

import "github.com/badoux/checkmail"

func ValidateEmail(email string) error {
	// Validate email format
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return err
	}

	// Validate email host (domain)
	err = checkmail.ValidateHost(email)
	if err != nil {
		return err
	}

	// Validate the existence of the email user on the mail server
	err = checkmail.ValidateHostAndUser("", "", email)
	if err != nil {
		return err
	}

	return nil
}
