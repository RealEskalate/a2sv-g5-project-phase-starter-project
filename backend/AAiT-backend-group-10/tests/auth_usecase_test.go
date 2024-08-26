package tests

import (
	"net/http"
	"testing"

	"aait.backend.g10/domain"
	"aait.backend.g10/tests/mocks"
	"aait.backend.g10/usecases"
	"aait.backend.g10/usecases/dto"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AuthUseCaseSuite struct {
	suite.Suite
	userRepo     *mocks.IUserRepository
	jwtService   *mocks.IJwtService
	pwdService   *mocks.IHashingService
	emailService *mocks.IEmailService
	authUsecase  usecases.IAuthUsecase
}

func (suite *AuthUseCaseSuite) SetupTest() {
	suite.userRepo = new(mocks.IUserRepository)
	suite.jwtService = new(mocks.IJwtService)
	suite.emailService = new(mocks.IEmailService)
	suite.pwdService = new(mocks.IHashingService)
	suite.authUsecase = usecases.NewAuthUsecase(
		suite.userRepo,
		suite.jwtService,
		suite.pwdService,
		suite.emailService,
	)
}

func (suite *AuthUseCaseSuite) TearDownTest() {
	suite.userRepo = nil
	suite.jwtService = nil
	suite.emailService = nil
	suite.pwdService = nil
	suite.authUsecase = nil
}
func (suite *AuthUseCaseSuite) TestRegisterUser_NoExistingUser_Positive() {
	// Arrange
	mockRegisterDto := &dto.RegisterUserDTO{
		FullName: "John Doe",
		Email:    "johndoe@example.com",
		Password: "password123",
		Bio:      "A bio",
	}

	hashedPassword := "hashedPassword"
	suite.userRepo.On("GetUserByEmail", mockRegisterDto.Email).Return(nil, nil)
	suite.pwdService.On("HashPassword", mockRegisterDto.Password).Return(hashedPassword, nil)
    suite.userRepo.On("Count").Return(int64(0), nil)
	suite.userRepo.On("CreateUser", mock.AnythingOfType("*domain.User")).Return(nil)

	// Act
	result, err := suite.authUsecase.RegisterUser(mockRegisterDto)

	// Assert
	suite.Nil(err)
	suite.NotNil(result)
}
func (suite *AuthUseCaseSuite) TestRegisterUser_UserCreationFailed_Negative() {
	// Arrange
	mockRegisterDto := &dto.RegisterUserDTO{
		FullName: "John Doe",
		Email:    "johndoe@example.com",
		Password: "password123",
		Bio:      "A bio",
	}

	hashedPassword := "hashedPassword"
	suite.userRepo.On("GetUserByEmail", mockRegisterDto.Email).Return(nil, nil)
	suite.pwdService.On("HashPassword", mockRegisterDto.Password).Return(hashedPassword, nil)
    suite.userRepo.On("Count").Return(int64(1), nil)
	suite.userRepo.On("CreateUser", mock.AnythingOfType("*domain.User")).Return(domain.ErrUserCreationFailed)

	// Act
	result, err := suite.authUsecase.RegisterUser(mockRegisterDto)

	// Assert
	suite.Nil(result)
	suite.Equal(err, domain.ErrUserCreationFailed)
}

func (suite *AuthUseCaseSuite) TestRegisterUser_ExistingUser_GoogleSignin_Positive() {
	// Arrange
	existingUser := &domain.User{
		ID: uuid.New(),
		GoogleSignIn: true,
		ImageURL: "",
	}
	mockRegisterDto := &dto.RegisterUserDTO{
		FullName: "John Doe",
		Email:    "johndoe@example.com",
		Password: "password123",
		Bio:      "A bio",
	}

	hashedPassword := "hashedPassword"
	suite.userRepo.On("GetUserByEmail", mockRegisterDto.Email).Return(existingUser, nil)
	suite.pwdService.On("HashPassword", mockRegisterDto.Password).Return(hashedPassword, nil)
	suite.userRepo.On("UpdateUserToken", existingUser).Return(nil)
	suite.userRepo.On("CreateUser", mock.AnythingOfType("*domain.User")).Return(nil)

	// Act
	result, err := suite.authUsecase.RegisterUser(mockRegisterDto)

	// Assert
	suite.Nil(err)
	suite.NotNil(result)
}
func (suite *AuthUseCaseSuite) TestRegisterUser_ExistingUser_GoogleSignin_UpdateTokenFailed_Negative() {
	// Arrange
	existingUser := &domain.User{
		ID: uuid.New(),
		GoogleSignIn: true,
		ImageURL: "",
	}
	mockRegisterDto := &dto.RegisterUserDTO{
		FullName: "John Doe",
		Email:    "johndoe@example.com",
		Password: "password123",
		Bio:      "A bio",
	}

	hashedPassword := "hashedPassword"
	suite.userRepo.On("GetUserByEmail", mockRegisterDto.Email).Return(existingUser, nil)
	suite.pwdService.On("HashPassword", mockRegisterDto.Password).Return(hashedPassword, nil)
	suite.userRepo.On("UpdateUserToken", existingUser).Return(domain.NewCustomError("Update token failed",  http.StatusBadRequest))

	// Act
	result, err := suite.authUsecase.RegisterUser(mockRegisterDto)

	// Assert
	suite.Nil(result)
	suite.Equal(err, domain.NewCustomError("Update token failed",  http.StatusBadRequest))
}
func (suite *AuthUseCaseSuite) TestRegisterUser_EmailExists_Negative() {
	// Arrange
	existingUser := &domain.User{
		ID: uuid.New(),
		GoogleSignIn: false,
		ImageURL: "",
	}
	mockRegisterDto := &dto.RegisterUserDTO{
		FullName: "John Doe",
		Email:    "johndoe@example.com",
		Password: "password123",
		Bio:      "A bio",
	}
	suite.userRepo.On("GetUserByEmail", mockRegisterDto.Email).Return(existingUser, nil)

	// Act
	result, err := suite.authUsecase.RegisterUser(mockRegisterDto)

	// Assert
	suite.Nil(result)
	suite.Equal(err, domain.ErrUserEmailExists)
}

func (suite *AuthUseCaseSuite) TestRegisterUser_HashingFailure_Negative() {
	// Arrange
	mockRegisterDto := &dto.RegisterUserDTO{
		FullName: "John Doe",
		Email:    "johndoe@example.com",
		Password: "password123",
		Bio:      "A bio",
	}

	suite.userRepo.On("GetUserByEmail", mockRegisterDto.Email).Return(nil, nil)
	suite.pwdService.On("HashPassword", mockRegisterDto.Password).Return("", domain.ErrPasswordHashingFailed)

	// Act
	result, err := suite.authUsecase.RegisterUser(mockRegisterDto)

	// Assert
	suite.Nil(result)
	suite.Equal(err, domain.ErrPasswordHashingFailed)
}
// Test for RegisterUser when user already exists
func (suite *AuthUseCaseSuite) TestLoginUser_ValidCredentials_Positive() {
    // Arrange
    loginUser := &dto.LoginUserDTO{
        Email:    "johndoe@example.com",
        Password: "password123",
    }
    user := &domain.User{
        ID:           uuid.New(),
        Email:        loginUser.Email,
        Password:     "hashedPassword",
        GoogleSignIn: false,
    }
    suite.userRepo.On("GetUserByEmail", loginUser.Email).Return(user, nil)
    suite.pwdService.On("CheckPasswordHash", loginUser.Password, user.Password).Return(true)
    suite.jwtService.On("GenerateToken", user).Return("accessToken", "refreshToken", nil)
    suite.userRepo.On("UpdateUserToken", user).Return(nil)

    // Act
    result, err := suite.authUsecase.LoginUser(loginUser)

    // Assert
    suite.Nil(err)
    suite.NotNil(result)
    suite.Equal(result.AccessToken, "accessToken")
    suite.Equal(result.RefreshToken, "refreshToken")
}

func (suite *AuthUseCaseSuite) TestLoginUser_InvalidCredentials_Negative() {
    // Arrange
    loginUser := &dto.LoginUserDTO{
        Email:    "johndoe@example.com",
        Password: "password123",
    }
    suite.userRepo.On("GetUserByEmail", loginUser.Email).Return(nil, nil)

    // Act
    result, err := suite.authUsecase.LoginUser(loginUser)

    // Assert
    suite.Nil(result)
    suite.Equal(err, domain.ErrInvalidCredentials)
}

func (suite *AuthUseCaseSuite) TestLoginUser_GoogleSignin_Negative() {
    // Arrange
    loginUser := &dto.LoginUserDTO{
        Email:    "johndoe@example.com",
        Password: "password123",
    }
    user := &domain.User{
        ID:           uuid.New(),
        Email:        loginUser.Email,
        Password:     "hashedPassword",
        GoogleSignIn: true,
    }
    suite.userRepo.On("GetUserByEmail", loginUser.Email).Return(user, nil)

    // Act
    result, err := suite.authUsecase.LoginUser(loginUser)

    // Assert
    suite.Nil(result)
    suite.Equal(err.Message, "please create account")
    suite.Equal(err.StatusCode, 400)
}

func (suite *AuthUseCaseSuite) TestLoginUser_PasswordCheckFailed_Negative() {
    // Arrange
    loginUser := &dto.LoginUserDTO{
        Email:    "johndoe@example.com",
        Password: "password123",
    }
    user := &domain.User{
        ID:           uuid.New(),
        Email:        loginUser.Email,
        Password:     "hashedPassword",
        GoogleSignIn: false,
    }
    suite.userRepo.On("GetUserByEmail", loginUser.Email).Return(user, nil)
    suite.pwdService.On("CheckPasswordHash", loginUser.Password, user.Password).Return(false)

    // Act
    result, err := suite.authUsecase.LoginUser(loginUser)

    // Assert
    suite.Nil(result)
    suite.Equal(err, domain.ErrInvalidCredentials)
}

func (suite *AuthUseCaseSuite) TestRefreshTokens_ValidToken_Positive() {
    // Arrange
    refreshToken := "validRefreshToken"
    token := &jwt.Token{
        Valid: true,
        Claims: jwt.MapClaims{
            "email": "johndoe@example.com",
        },
    }
    user := &domain.User{
        ID:           uuid.New(),
        Email:        "johndoe@example.com",
        RefreshToken: refreshToken,
    }
    suite.jwtService.On("ValidateToken", refreshToken).Return(token, nil)
    suite.userRepo.On("GetUserByEmail", "johndoe@example.com").Return(user, nil)
    suite.jwtService.On("GenerateToken", user).Return("newAccessToken", "newRefreshToken", nil)
    suite.userRepo.On("UpdateUserToken", user).Return(nil)

    // Act
    result, err := suite.authUsecase.RefreshTokens(refreshToken)

    // Assert
    suite.Nil(err)
    suite.NotNil(result)
    suite.Equal(result.AccessToken, "newAccessToken")
    suite.Equal(result.RefreshToken, "newRefreshToken")
}

func (suite *AuthUseCaseSuite) TestRefreshTokens_InvalidToken_Negative() {
    // Arrange
    refreshToken := "invalidRefreshToken"
    suite.jwtService.On("ValidateToken", refreshToken).Return(nil, domain.ErrInvalidToken)

    // Act
    result, err := suite.authUsecase.RefreshTokens(refreshToken)

    // Assert
    suite.Nil(result)
    suite.Equal(err, domain.ErrInvalidToken)
}

func (suite *AuthUseCaseSuite) TestRefreshTokens_UserNotFound_Negative() {
    // Arrange
    refreshToken := "validRefreshToken"
    token := &jwt.Token{
        Valid: true,
        Claims: jwt.MapClaims{
            "email": "johndoe@example.com",
        },
    }
    suite.jwtService.On("ValidateToken", refreshToken).Return(token, nil)
    suite.userRepo.On("GetUserByEmail", "johndoe@example.com").Return(nil, domain.ErrUserNotFound)

    // Act
    result, err := suite.authUsecase.RefreshTokens(refreshToken)

    // Assert
    suite.Nil(result)
    suite.Equal(err, domain.ErrUserNotFound)
}

func (suite *AuthUseCaseSuite) TestRefreshTokens_InvalidStoredRefreshToken_Negative() {
    // Arrange
    refreshToken := "validRefreshToken"
    token := &jwt.Token{
        Valid: true,
        Claims: jwt.MapClaims{
            "email": "johndoe@example.com",
        },
    }
    user := &domain.User{
        ID:           uuid.New(),
        Email:        "johndoe@example.com",
        RefreshToken: "differentStoredRefreshToken",
    }
    suite.jwtService.On("ValidateToken", refreshToken).Return(token, nil)
    suite.userRepo.On("GetUserByEmail", "johndoe@example.com").Return(user, nil)
	suite.jwtService.On("ValidateToken", user.RefreshToken).Return(nil, domain.ErrInvalidToken)

    // Act
    result, err := suite.authUsecase.RefreshTokens(refreshToken)

    // Assert
    suite.Nil(result)
    suite.Equal(err, domain.ErrInvalidToken)
}

func (suite *AuthUseCaseSuite) TestForgotPassword_ValidEmail_Positive() {
    // Arrange
    forgotPasswordDTO := &dto.ForgotPasswordRequestDTO{
        Email: "johndoe@example.com",
    }
    user := &domain.User{
        ID:    uuid.New(),
        Email: forgotPasswordDTO.Email,
    }
    suite.userRepo.On("GetUserByEmail", forgotPasswordDTO.Email).Return(user, nil)
    suite.jwtService.On("GenerateResetToken", user.Email, mock.AnythingOfType("int64")).Return("resetToken", nil)
    suite.userRepo.On("UpdateUserToken", user).Return(nil)
    suite.emailService.On("SendResetEmail", user.Email, "resetToken").Return(nil)

    // Act
    err := suite.authUsecase.ForgotPassword(forgotPasswordDTO)

    // Assert
    suite.Nil(err)
}

func (suite *AuthUseCaseSuite) TestForgotPassword_UserNotFound_Negative() {
    // Arrange
    forgotPasswordDTO := &dto.ForgotPasswordRequestDTO{
        Email: "johndoe@example.com",
    }
    suite.userRepo.On("GetUserByEmail", forgotPasswordDTO.Email).Return(nil, domain.ErrUserNotFound)

    // Act
    err := suite.authUsecase.ForgotPassword(forgotPasswordDTO)

    // Assert
    suite.Equal(err, domain.ErrUserNotFound)
}

func (suite *AuthUseCaseSuite) TestForgotPassword_TokenGenerationFailed_Negative() {
    // Arrange
    forgotPasswordDTO := &dto.ForgotPasswordRequestDTO{
        Email: "johndoe@example.com",
    }
    user := &domain.User{
        ID:    uuid.New(),
        Email: forgotPasswordDTO.Email,
    }
    suite.userRepo.On("GetUserByEmail", forgotPasswordDTO.Email).Return(user, nil)
    suite.jwtService.On("GenerateResetToken", user.Email, mock.AnythingOfType("int64")).Return("", domain.ErrTokenGenerationFailed)

    // Act
    err := suite.authUsecase.ForgotPassword(forgotPasswordDTO)

    // Assert
    suite.Equal(err, domain.ErrTokenGenerationFailed)
}

func (suite *AuthUseCaseSuite) TestResetPassword_ValidToken_Positive() {
    // Arrange
    resetPasswordDTO := &dto.ResetPasswordRequestDTO{
        Token:       "validToken",
        NewPassword: "newPassword",
    }
    token := &jwt.Token{
        Valid: true,
        Claims: jwt.MapClaims{
            "email": "johndoe@example.com",
            "code":  12345.0,
        },
    }
    user := &domain.User{
        ID:        uuid.New(),
        Email:     "johndoe@example.com",
        ResetCode: 12345,
        ResetToken: "validToken",
    }
    suite.jwtService.On("ValidateToken", resetPasswordDTO.Token).Return(token, nil)
    suite.userRepo.On("GetUserByEmail", "johndoe@example.com").Return(user, nil)
    suite.pwdService.On("HashPassword", resetPasswordDTO.NewPassword).Return("hashedPassword", nil)
    suite.userRepo.On("UpdateUserToken", user).Return(nil)

    // Act
    err := suite.authUsecase.ResetPassword(resetPasswordDTO)

    // Assert
    suite.Nil(err)
}

func (suite *AuthUseCaseSuite) TestResetPassword_InvalidToken_Negative() {
    // Arrange
    resetPasswordDTO := &dto.ResetPasswordRequestDTO{
        Token:       "invalidToken",
        NewPassword: "newPassword",
    }
    suite.jwtService.On("ValidateToken", resetPasswordDTO.Token).Return(nil, domain.ErrUserNotFound)

    // Act
    err := suite.authUsecase.ResetPassword(resetPasswordDTO)

    // Assert
    suite.Equal(err, domain.ErrUserNotFound)
}

func (suite *AuthUseCaseSuite) TestResetPassword_WrongResetCode_Negative() {
    // Arrange
    resetPasswordDTO := &dto.ResetPasswordRequestDTO{
        Token:       "validToken",
        NewPassword: "newPassword",
    }
    token := &jwt.Token{
        Valid: true,
        Claims: jwt.MapClaims{
            "email": "johndoe@example.com",
            "code":  99999.0,
        },
    }
    user := &domain.User{
        ID:        uuid.New(),
        Email:     "johndoe@example.com",
        ResetCode: 12345,
        ResetToken: "validToken",
    }
    suite.jwtService.On("ValidateToken", resetPasswordDTO.Token).Return(token, nil)
    suite.userRepo.On("GetUserByEmail", "johndoe@example.com").Return(user, nil)

    // Act
    err := suite.authUsecase.ResetPassword(resetPasswordDTO)

    // Assert
    suite.Equal(err, domain.ErrInvalidResetCode)
}

func (suite *AuthUseCaseSuite) TestHandleGoogleCallback_ExistingUser_Positive() {
    // Arrange
    userDto := &domain.User{
        Email:        "johndoe@gmail.com",
        GoogleSignIn: true,
    }
    user := &domain.User{
        ID:           uuid.New(),
        Email:        "johndoe@gmail.com",
        GoogleSignIn: true,
    }
    suite.userRepo.On("GetUserByEmail", userDto.Email).Return(user, nil)
    suite.jwtService.On("GenerateToken", user).Return("accessToken", "refreshToken", nil)
    suite.userRepo.On("UpdateUserToken", user).Return(nil)

    // Act
    accessToken, refreshToken, err := suite.authUsecase.HandleGoogleCallback(userDto)

    // Assert
    suite.Nil(err)
    suite.Equal(accessToken, "accessToken")
    suite.Equal(refreshToken, "refreshToken")
}

func (suite *AuthUseCaseSuite) TestHandleGoogleCallback_NewUser_Positive() {
    // Arrange
    userDto := &domain.User{
        Email:        "johndoe@gmail.com",
        GoogleSignIn: true,
    }
    suite.userRepo.On("GetUserByEmail", userDto.Email).Return(nil, domain.ErrUserNotFound)
    suite.jwtService.On("GenerateToken", userDto).Return("accessToken", "refreshToken", nil)
    suite.userRepo.On("CreateUser", userDto).Return(nil)

    // Act
    accessToken, refreshToken, err := suite.authUsecase.HandleGoogleCallback(userDto)

    // Assert
    suite.Nil(err)
    suite.Equal(accessToken, "accessToken")
    suite.Equal(refreshToken, "refreshToken")
}

func (suite *AuthUseCaseSuite) TestHandleGoogleCallback_LoginRequired_Negative() {
    // Arrange
    userDto := &domain.User{
        Email:        "johndoe@gmail.com",
        GoogleSignIn: false,
    }
    user := &domain.User{
        ID:           uuid.New(),
        Email:        "johndoe@gmail.com",
        GoogleSignIn: false,
    }
    suite.userRepo.On("GetUserByEmail", userDto.Email).Return(user, nil)

    // Act
    accessToken, refreshToken, err := suite.authUsecase.HandleGoogleCallback(userDto)

    // Assert
    suite.NotNil(err)
    suite.Equal(err.Error(), "login required")
    suite.Equal(accessToken, "")
    suite.Equal(refreshToken, "")
}

func (suite *AuthUseCaseSuite) TestHandleGoogleCallback_UserCreationFailed_Negative() {
    // Arrange
    userDto := &domain.User{
        Email:        "johndoe@gmail.com",
        GoogleSignIn: true,
    }
    suite.userRepo.On("GetUserByEmail", userDto.Email).Return(nil, domain.ErrUserNotFound)
    suite.jwtService.On("GenerateToken", userDto).Return("accessToken", "refreshToken", nil)
    suite.userRepo.On("CreateUser", userDto).Return(domain.ErrUserCreationFailed)

    // Act
    accessToken, refreshToken, err := suite.authUsecase.HandleGoogleCallback(userDto)

    // Assert
    suite.NotNil(err)
    suite.Equal(err.Error(), "failed to register user")
    suite.Equal(accessToken, "")
    suite.Equal(refreshToken, "")
}

func TestAuthCaseSuite(t *testing.T) {
	suite.Run(t, new(AuthUseCaseSuite))
}
