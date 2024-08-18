package infrastructure

import (
	"context"
	"Blog_Starter/domain"
)

//implement te functions in the tokenutil.go
type NewTokenMangaer struct{}

func (m *NewTokenMangaer) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	panic("implement me")
}

func (m *NewTokenMangaer) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	panic("implement me")
}

func (m *NewTokenMangaer) InvalidateAccessToken(ctx context.Context, token string) error {
	panic("implement me")
}

func (m *NewTokenMangaer) InvalidateRefreshToken(ctx context.Context, token string) error {
	panic("implement me")
}
