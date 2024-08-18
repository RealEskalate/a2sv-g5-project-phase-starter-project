package usecases

import (
    // "errors"
    "meleket/domain"
)

type TokenUsecase struct {
    userRepo domain.TokenRepositoryInterface
    // jwtSvc   domain.JWTService
}

// func NewTokenRepository(ur repository.UserRepositoryInterface, js domain.JWTService) *UserUsecase {
func NewTokenUsecase(ur domain.TokenRepositoryInterface) *TokenUsecase{
    return &TokenUsecase{
        userRepo: ur,
        // jwtSvc:   js,
    }
}

// // RefreshToken refreshes a user's JWT token
// func (u *UserUsecase) RefreshToken(refreshToken string) (string, error) {
//     storedToken, err := u.userRepo.FindRefreshToken(refreshToken)
//     if err != nil {
//         return "", errors.New("invalid refresh token")
//     }

//     // Check if the refresh token is expired
//     if storedToken.ExpiresAt.Before(time.Now()) {
//         return "", errors.New("refresh token expired")
//     }

//     // Generate a new JWT token
//     newToken, err := u.jwtSvc.GenerateToken(storedToken.UserID.Hex(), storedToken.ExpiresAt.String())
//     if err != nil {
//         return "", err
//     }

//     return newToken, nil
// }
