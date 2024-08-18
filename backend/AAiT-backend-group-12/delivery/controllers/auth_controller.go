package controllers
import (
	"blog_api/domain"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	usecase domain.UserUsecaseInterface
}

func GetHTTPErrorCode(err domain.CodedError) int {
	switch err.GetCode() {
	case domain.ERR_BAD_REQUEST:
		return 400
	case domain.ERR_UNAUTHORIZED:
		return 401
	case domain.ERR_FORBIDDEN:
		return 403
	case domain.ERR_NOT_FOUND:
		return 404
	case domain.ERR_CONFLICT:
		return 409
	default:
		return 500
	}
}

func NewAuthController(usecase domain.UserUsecaseInterface) *AuthController {
	return &AuthController{usecase: usecase}
}

func (controller *AuthController) HandleSignup(c *gin.Context) {
	var newUser domain.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(400, domain.Response{"error": "Invalid input"})
		return
	}

	err := controller.usecase.Signup(c, &newUser)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}

	c.JSON(201, domain.Response{"message": "User created"})
}

func (controller *AuthController) HandleLogin(c *gin.Context) {
	var newUser domain.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(400, domain.Response{"error": "Invalid input"})
		return
	}

	acK, rfK, err := controller.usecase.Login(c, &newUser)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}

	c.JSON(201, domain.Response{"accessToken": acK, "refreshToken": rfK})
}
