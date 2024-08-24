package user_usecase

import "context"

func (u *userUsecase) Logout(ctx context.Context, userID string) error {
	// Delete the refresh token associated with the user
	return u.authService.DeleteRefreshToken(ctx, userID)
}
