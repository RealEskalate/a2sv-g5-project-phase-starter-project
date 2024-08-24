package utils

import (
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

// WordCountValidator is a custom validator for checking the number of words in a string
func WordCountValidator(fl validator.FieldLevel) bool {
	words := strings.Fields(fl.Field().String())
	minWords := fl.Param()
	minWordCount, _ := strconv.Atoi(minWords)
	return len(words) >= minWordCount
}
