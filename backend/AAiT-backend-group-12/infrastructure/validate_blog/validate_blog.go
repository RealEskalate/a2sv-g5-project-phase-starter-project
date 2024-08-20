package validateblog

import (
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func WordCountValidator(fl validator.FieldLevel) bool {
	words := strings.Fields(fl.Field().String()) // Count words
	minWords := fl.Param()                       // Get min word count from the tag
	minWordCount, _ := strconv.Atoi(minWords) // Convert minWords to an integer
	return len(words) >= minWordCount         // Compare word count to minWordCount
}