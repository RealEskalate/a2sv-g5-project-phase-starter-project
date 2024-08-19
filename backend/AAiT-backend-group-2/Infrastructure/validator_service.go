// validator_service.go
package infrastructure

import (
    "github.com/go-playground/validator/v10"
)


type ValidatorService struct {
    validate *validator.Validate
}

func NewValidatorService() *ValidatorService {
    return &ValidatorService{
        validate: validator.New(),
    }
}

func ValidateStruct[T any](v *ValidatorService, s T) error {
    return v.validate.Struct(s)
}