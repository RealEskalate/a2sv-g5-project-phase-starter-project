


// RefreshToken refreshes a user's JWT token
func (u *UserUsecase) RefreshToken(refreshToken *domain.RefreshToken) (string, error) {
	storedToken, err := u.userRepo.FindRefreshToken(refreshToken.UserID)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	// Assuming that the storedToken contains userID and Role, you would generate a new token
	newToken, err := u.jwtSvc.GenerateToken(storedToken.UserID, storedToken.ExpiresAt.String())
	if err != nil {
		return "", err
	}

	return newToken, nil
}