package domain

import "github.com/go-playground/validator/v10"

type StructValidator struct {
	validator *validator.Validate
}

type ValidateInterface interface {
	ValidateStruct(interface{}) error
}

func NewValidator() *StructValidator {
	return &StructValidator{
		validator: validator.New(),
	}
}

func (v *StructValidator) ValidateStruct(model interface{}) error {
	return v.validator.Struct(model)
}