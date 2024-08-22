package controller

import (
	"net/http"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.User_Usecase_interface
}

func NewUserController(usecase domain.User_Usecase_interface) *UserController {
	return &UserController{UserUsecase: usecase}
}

func (controller *UserController) GetOneUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		user,err := controller.UserUsecase.GetOneUser(id)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound , gin.H{"error" : "user not found"})
			return
		}

		ctx.IndentedJSON(http.StatusOK , gin.H{"data" : user})

	}
}

func (controller *UserController) GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users,err := controller.UserUsecase.GetUsers()

		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound , gin.H{"error" : "not found"})
			return
		}
		ctx.IndentedJSON(http.StatusOK , gin.H{"data" : users})
	}
}

func (controller *UserController) UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		update_user := domain.UpdateUser{}
		id := ctx.Param("id")
		err := ctx.BindJSON(&update_user)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest , gin.H{"error" : "not able to update"})
			return
		}

		user,err := controller.UserUsecase.UpdateUser(id , update_user)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest , gin.H{"error" : "error updating user"})
			return
		}

		ctx.IndentedJSON(http.StatusOK , gin.H{"data" : user})
	}
}

func (controller *UserController)UpdatePassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var update_password domain.UpdatePassword
		id := ctx.Param("id")
		if err := ctx.BindJSON(&update_password); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
			return
		}

		response_user,err := controller.UserUsecase.UpdatePassword(id ,update_password)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusAccepted , gin.H{"data" : response_user})
	}
}

func (controller *UserController) DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		err := controller.UserUsecase.DeleteUser(id)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest , gin.H{"error" : "couldn't delete"})
			return
		}
		ctx.IndentedJSON(http.StatusAccepted , gin.H{"message" : "accepted!"})
	}
}

// func (controller *UserController) LogIn() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		log_in_model := domain.LogINUser{}

// 		err := ctx.BindJSON(&log_in_model)
// 		if err != nil {
// 			ctx.IndentedJSON(http.StatusBadRequest , gin.H{"error" : "invalid form data"})
// 			return
// 		}

// 		user,err := controller.UserUsecase.LogIn(log_in_model)
// 		if err != nil {
// 			ctx.IndentedJSON(http.StatusBadRequest , gin.H{"error" : "failed to login"})
// 			return
// 		}
// 		ctx.IndentedJSON(http.StatusOK , gin.H{"data" : user})
// 	}
// }

// func (controller *UserController) Register() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		register_user := domain.RegisterUser{}

// 		err := ctx.BindJSON(&register_user)
// 		if err != nil {
// 			ctx.IndentedJSON(http.StatusBadRequest , gin.H{"error" : "invalid form data"})
// 			return
// 		}

// 		user,err := controller.UserUsecase.Register(register_user)
// 		if err != nil {
// 			ctx.IndentedJSON(http.StatusBadRequest , gin.H{"error" : "failed to register"})
// 			return
// 		}
// 		ctx.IndentedJSON(http.StatusOK , gin.H{"data" : user})
// 	}
// }

func (controller *UserController) FilterUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		filter := make(map[string]string)
		for key,value := range query {
			filter[key] = value[0]
		}

		users,err := controller.UserUsecase.FilterUser(filter)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound , gin.H{"error" : "not found"})
			return
		}
		ctx.IndentedJSON(http.StatusOK , gin.H{"data" : users})
	}
}

func (controller *UserController) PromoteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		user,err := controller.UserUsecase.PromoteUser(id)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest , gin.H{"error" : "error updating user"})
			return
		}

		ctx.IndentedJSON(http.StatusOK , gin.H{"data" : user})
	}
}

func (controller *UserController) DemoteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		user,err := controller.UserUsecase.DemoteUser(id)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest , gin.H{"error" : "error updating user"})
			return
		}

		ctx.IndentedJSON(http.StatusOK , gin.H{"data" : user})
	}
}