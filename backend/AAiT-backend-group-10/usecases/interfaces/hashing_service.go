package interfaces

import "aait.backend.g10/domain"

type IHashingService interface {
	HashPassword(password string) (string, *domain.CustomError)
	CheckPasswordHash(password string, hash string) bool
}