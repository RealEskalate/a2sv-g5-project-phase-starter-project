package tests

import (
	"blog_api/infrastructure/utils"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
)

type UtilsTestSuite struct {
	suite.Suite
}

func (suite *UtilsTestSuite) TestGenerateToken() {
	tests := []int{12, 16, 32, 64}

	for _, v := range tests {
		tw, err := utils.GenerateToken(v)
		suite.NoError(err)
		suite.GreaterOrEqual(len(tw), v)
	}
}

func (suite *UtilsTestSuite) TestWordCountValidator() {
	validate := validator.New()
	err := validate.RegisterValidation("wordCount", utils.WordCountValidator)

	suite.NoError(err)

	type Blog struct {
		Title string `validate:"required,wordCount=5"`
	}

	green_tests := []string{"this string has more than five words", "this string has five words", "the red fox jumps over the lazy dog"}
	for _, v := range green_tests {
		blog := Blog{
			Title: v,
		}

		err = validate.Struct(blog)
		suite.NoError(err)
	}

	red_tests := []string{"", "this", "this string has four", "w a aa  a              "}
	for _, v := range red_tests {
		blog := Blog{
			Title: v,
		}

		err = validate.Struct(blog)
		suite.Error(err)
	}
}

func TestUtils(t *testing.T) {
	suite.Run(t, new(UtilsTestSuite))
}
