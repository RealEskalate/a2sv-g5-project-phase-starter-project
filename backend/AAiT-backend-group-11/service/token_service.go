package service

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type tokenService struct {
	accessTokenSecret, refreshTokenSecret string
	tokenRepository                       interfaces.RefreshTokenRepository
	userRepo                              interfaces.UserRepository
}

func NewTokenService(accessTokenSecret, refreshTokenSecret string, token_repo interfaces.RefreshTokenRepository, user_repo interfaces.UserRepository) interfaces.TokenService {
	return &tokenService{
		accessTokenSecret:  accessTokenSecret,
		refreshTokenSecret: refreshTokenSecret,
		tokenRepository:    token_repo,
		userRepo:           user_repo,
	}
}

func (service *tokenService) InvalidateAccessToken(token string) (string, error) {

	//extract the claims from the token
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return "", err
	}

	claims["exp"] = 0
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	expiredToken, err := newToken.SignedString([]byte(service.accessTokenSecret))
	if err != nil {
		return "", err
	}
	return expiredToken, nil
}

func (service *tokenService) InvalidateRefreshToken(token string) (string, error) {

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return "", err
	}

	claims["exp"] = 0
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	expiredToken, err := newToken.SignedString([]byte(service.refreshTokenSecret))
	if err != nil {
		return "", err
	}
	return expiredToken, nil
}

func (service *tokenService) GenerateAccessToken(user *entities.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userName": user.Username,
			"userId": user.ID.Hex(),
			"email":  user.Email,
			"role":   user.Role,
			"exp":    time.Now().Add(time.Hour * 300).Unix(),
		})
	accessToken, err := token.SignedString([]byte(service.accessTokenSecret))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (service *tokenService) GenerateRefreshToken(user *entities.User) (*entities.RefreshToken, error) {
	err:=service.tokenRepository.DeleteRefreshTokenByUserId(user.ID.Hex())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userName": user.Username,
			"userId": user.ID.Hex(),
			"email": user.Email,
			"role":  user.Role,
			"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
		})
	refreshToken, err := token.SignedString([]byte(service.refreshTokenSecret))
	if err != nil {
		return &entities.RefreshToken{}, err
	}



	// this may be moved to other layers
	refresh_token := entities.RefreshToken{
		UserID:    user.ID.Hex(),
		Token:     refreshToken,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30),
	}
	service.tokenRepository.CreateRefreshToken(&refresh_token)
	return &refresh_token, nil
}


func (service *tokenService) VerifyAccessToken(token string) error {

	accessToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		//if the token is not signed with HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(service.accessTokenSecret), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := accessToken.Claims.(jwt.MapClaims); ok && accessToken.Valid {
		// if token is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			if err != nil {
				return errors.New("token is expired login again")
			}
		}
	}
	return nil
}

func (service *tokenService) VerifyRefreshToken(token string) error {
	fmt.Println("from token service",token)
	refreshToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(service.refreshTokenSecret), nil
	})
	if err != nil {
		fmt.Println("from token service 1st",err)
		return err
	}
	
	if claims, ok := refreshToken.Claims.(jwt.MapClaims); ok && refreshToken.Valid {
		userId:=claims["userId"].(string)
		storedToken,err:=service.tokenRepository.FindRefreshTokenByUserId(userId)
		if err!=nil{
			fmt.Println("from token service",err)

			return err
		}
		if storedToken.Token!=token{
			fmt.Println("from token service","token is not e")
			return errors.New("token is not valid")
		}
		// if token is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			service.tokenRepository.DeleteRefreshTokenByUserId(userId)
			return errors.New("token is expired login again")
		}
	
		return nil
	}
	return nil
}

func (service *tokenService) GetClaimsFromAccessToken(token string) map[string]string {
	
	Token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(service.accessTokenSecret), nil
	})

	if err != nil {
		return map[string]string{}
	}


	if claims, ok := Token.Claims.(jwt.MapClaims); ok && Token.Valid {
		resp := make(map[string]string)
		for key, value := range claims {
			resp[key] = fmt.Sprintf("%v", value)
		}
		return resp
	}

	return map[string]string{}
}

func (service *tokenService) GetClaimsFromRefreshToken(token string) map[string]string {
	
	Token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(service.refreshTokenSecret), nil
	})

	if err != nil {
		return map[string]string{}
	}
	if claims, ok := Token.Claims.(jwt.MapClaims); ok && Token.Valid {
		resp := make(map[string]string)
		for key, value := range claims {
			resp[key] = fmt.Sprintf("%v", value)
		}
		return resp
	}

	return map[string]string{}
}


func (service *tokenService) RefreshAccessToken(refreshToken string) (string, error) {
	err:=service.VerifyRefreshToken(refreshToken)

	if err!=nil{
		return "",err
	}
	claims:=service.GetClaimsFromRefreshToken(refreshToken)
	if len(claims)==0{
		return "",errors.New("no claims found,invalid token")
	}
	userId:=claims["userId"]
	user, err := service.userRepo.FindUserById(userId)
	if err != nil {
		return "", err
	}
	newToken, _ := service.GenerateAccessToken(user)
	return newToken, nil
}

func (service *tokenService) CreateRefreshToken(refreshToken *entities.RefreshToken) (*entities.RefreshToken, error) {
	return service.tokenRepository.CreateRefreshToken(refreshToken)
}

func (service *tokenService) DeleteRefreshTokenByUserId(userId string) error {
	return service.tokenRepository.DeleteRefreshTokenByUserId(userId)
}

func (service *tokenService) FindRefreshTokenByUserId(userId string) (*entities.RefreshToken, error) {
	return service.tokenRepository.FindRefreshTokenByUserId(userId)
}