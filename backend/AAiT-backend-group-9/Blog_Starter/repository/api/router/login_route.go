package router

// import (
//     "time"

//     "Blog_Starter/api/controller"
//     "Blog_Starter/domain"
//     "Blog_Starter/repository"
//     "Blog_Starter/usecase"

//     "github.com/gin-gonic/gin"
//     "go.mongodb.org/mongo-driver/mongo"
// )

// func LoginRouter(timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
//     lr := repository.NewLoginRepository(db, domain.CollectionOTP)
//     ur := repository.NewUserRepository(db, domain.CollectionUser)
//     sc := controller.NewSignUpController(
//         usecase.NewSignUpUsecase(ur, timeout),
//         usecase.NewOtpUsecase(or, timeout),
//     )
//     group.POST("/signup", sc.SignUp)
//     group.POST("/verify-email", sc.VerifyEmail)
//     group.POST("/resend-otp", sc.ResendOTP)
// }
