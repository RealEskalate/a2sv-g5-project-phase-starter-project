package usecase

import (
	domain "AAiT-backend-group-8/Domain"
	infrastructure "AAiT-backend-group-8/Infrastructure"
	repository "AAiT-backend-group-8/Infrastructure/mongodb"
	interfaces "AAiT-backend-group-8/Interfaces"
	"errors"
)

type CommentUseCase struct {
	repository     *repository.CommentRepository
	infrastructure infrastructure.Infrastructure
	tokenService   interfaces.ITokenService
}

func NewCommentUseCase(commentRepository *repository.CommentRepository, infrastructure infrastructure.Infrastructure, tokenService interfaces.ITokenService) *CommentUseCase {
	return &CommentUseCase{
		repository:     commentRepository,
		infrastructure: infrastructure,
		tokenService:   tokenService,
	}
}

func (uc *CommentUseCase) CreateComment(comment *domain.Comment, blogID string) error {
	comment.CreatedAt = uc.infrastructure.GetCurrentTime()
	primitiveID, err := uc.infrastructure.ConvertToPrimitiveObjectID(blogID)
	if err != nil {
		return errors.New("invalid blog id")
	}
	comment.BlogID = primitiveID
	err = uc.repository.CreateComment(comment)
	if err != nil {
		return err
	}
	return nil
}

func (uc *CommentUseCase) GetComments(blogID string) ([]domain.Comment, error) {
	primitiveID, err := uc.infrastructure.ConvertToPrimitiveObjectID(blogID)
	if err != nil {
		return nil, errors.New("invalid blog id")
	}
	comments, err := uc.repository.GetComments(primitiveID)
	if err != nil {
		return nil, err
	}
	return comments, nil

}

func (uc *CommentUseCase) DeleteComment(commentID string) (string, error) {
	primitiveID, err := uc.infrastructure.ConvertToPrimitiveObjectID(commentID)
	if err != nil {
		return "", errors.New("invalid comment id")
	}
	blogID, err := uc.repository.DeleteComment(primitiveID)
	if err != nil {
		return blogID, err
	}
	return blogID, nil
}

func (uc *CommentUseCase) UpdateComment(comment *domain.Comment, commentID string) (string, error) {

	primitive, err := uc.infrastructure.ConvertToPrimitiveObjectID(commentID)

	if err != nil {
		return "", errors.New("invalid comment id")
	}

	comment.Id = primitive

	res, err := uc.repository.UpdateComment(comment)

	if err != nil {
		return "", err
	}
	return res, nil
}

func (uc *CommentUseCase) DeleteCommentsOfBlog(blogID string) error {
	primitiveID, err := uc.infrastructure.ConvertToPrimitiveObjectID(blogID)
	if err != nil {

		return errors.New("invalid blog id")
	}
	err = uc.repository.DeleteCommentsOfBlog(primitiveID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *CommentUseCase) DecodeToken(tokenStr string) (string, string, error) {
	// Parse the token
	myMap, err := uc.tokenService.GetClaimsOfToken(tokenStr)
	if err != nil {
		return "", "", err
	}
	if myMap == nil {
		return "", "", errors.New("invalid token - from decode token")
	}

	// Assert the types of the values to string
	id, ok := myMap["id"].(string)
	if !ok {
		return "", "", errors.New("invalid token - id claim is not a string")
	}

	name, ok := myMap["name"].(string)
	if !ok {
		return "", "", errors.New("invalid token - name claim is not a string")
	}

	return id, name, nil
}

func (uc *CommentUseCase) GetCommentByID(commentID string) (*domain.Comment, error) {

	primitiveID, err := uc.infrastructure.ConvertToPrimitiveObjectID(commentID)

	if err != nil {
		return nil, errors.New("invalid comment id")
	}

	comment, err1 := uc.repository.GetCommentByID(primitiveID)

	if err1 != nil {
		return nil, err1
	}

	return comment, nil
}
