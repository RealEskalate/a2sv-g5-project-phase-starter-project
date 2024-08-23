package dtos

import "github.com/go-playground/validator"

func (c *CreateAccountRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func (c *PasswordResetRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func (c *SetUpPasswordRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func (c *LoginRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func (c *LogoutRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func (c *ProfileUpdateRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
