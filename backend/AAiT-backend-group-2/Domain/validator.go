package domain

type Validator[T any] interface {
	Validate(entity T) error
}