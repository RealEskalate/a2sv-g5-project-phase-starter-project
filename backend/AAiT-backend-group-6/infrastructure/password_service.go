package infrastructure

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) string{
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err!=nil{
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) error{
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))

	return err
}