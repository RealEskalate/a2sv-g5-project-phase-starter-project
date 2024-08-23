package commentusecase

import (
	domain "AAiT-backend-group-2/Domain"
	"context"
)

type commentUsecase struct {
	repo domain.CommentRepository
}

func NewCommentUsecase(repo domain.CommentRepository) domain.CommentUsecase {
	return &commentUsecase{
		repo: repo,
	}
}

func (u *commentUsecase) GetAllComments(ctx context.Context) ([]domain.Comment, error) {
	return u.repo.FindAll(ctx)
}

func (u *commentUsecase) GetCommentByID(ctx context.Context, id string) (*domain.Comment, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *commentUsecase) GetCommentsByBlogID(ctx context.Context, blogID string) ([]domain.Comment, error) {
	return u.repo.FindByBlogID(ctx, blogID)
}

func (u *commentUsecase) CreateComment(ctx context.Context, comment *domain.Comment) error {
	return u.repo.Save(ctx, comment)
}

func (u *commentUsecase) UpdateComment(ctx context.Context, comment *domain.Comment) error {
	return u.repo.Update(ctx, comment)
}

func (u *commentUsecase) DeleteComment(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}
