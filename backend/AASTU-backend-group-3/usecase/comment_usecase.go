package usecase

import (
	"errors"
	"group3-blogApi/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type commentUsecase struct {
	commentRepo domain.CommentRepository
}

func NewCommentUsecase(commentRepo domain.CommentRepository) domain.CommentUsecase {
	return &commentUsecase{
		commentRepo: commentRepo,
	}
}

func (u *commentUsecase) CreateComment(comment *domain.Comment) (*domain.Comment, error) {
	if comment.PostID.IsZero()  || comment.Content == "" {
		return nil, errors.New("missing required fields")
	}
	return u.commentRepo.CreateComment(comment)
}

func (u *commentUsecase) UpdateComment(comment *domain.Comment, role_, userId string) (*domain.Comment, error) {
	existingComment, err := u.GetCommentByID(comment.ID.Hex())
	if err != nil {
		return nil, err
	}
	if existingComment.UserID.Hex() != userId || role_ != "admin" {
		return nil, errors.New("unauthorized")
	}
	if comment.ID.IsZero() {
		return nil, errors.New("invalid comment ID")
	}
	return u.commentRepo.UpdateComment(comment)
}

func (u *commentUsecase) DeleteComment(commentID, role_, UserID string) (*domain.Comment, error) {
	existingComment, err := u.GetCommentByID(commentID)
	if err != nil {
		return nil, err
	}
	if existingComment.UserID.Hex() != UserID || role_ != "admin" {
		return nil, errors.New("unauthorized")
	}

	objID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return nil, errors.New("invalid comment ID")
	}
	return u.commentRepo.DeleteComment(objID)
}

func (u *commentUsecase) GetCommentByID(commentID string) (*domain.Comment, error) {
	objID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return nil, errors.New("invalid comment ID")
	}
	return u.commentRepo.GetCommentByID(objID)
}

func (u *commentUsecase) GetComments(postID string, page, limit int) ([]domain.Comment, error) {
	if page <= 0 || limit <= 0 {
		return nil, errors.New("invalid pagination parameters")
	}
	comments, err := u.commentRepo.GetCommentsByPostID(postID, int64(page), int64(limit))
	if err != nil {
		return nil, err
	}
	return convertComments(comments), nil
}

func convertComments(comments []*domain.Comment) []domain.Comment {
	result := make([]domain.Comment, len(comments))
	for i, c := range comments {
		result[i] = *c
	}
	return result
}

func (u *commentUsecase) CreateReply(reply *domain.Reply) (*domain.Reply, error) {
	if reply.CommentID.IsZero() || reply.UserID == "" || reply.Content == "" {
		return nil, errors.New("missing required fields")
	}
	return u.commentRepo.CreateReply(reply)
}

func (u *commentUsecase) UpdateReply(reply *domain.Reply, userID string) (*domain.Reply, error) {
	existingReply, err := u.commentRepo.GetReplyByID(reply.ID)
	if err != nil {
		return nil, err
	}
	if existingReply.UserID != userID {
		return nil, errors.New("unauthorized")
	}
	if reply.ID.IsZero() {
		return nil, errors.New("invalid reply ID")
	}
	return u.commentRepo.UpdateReply(reply)
}

func (u *commentUsecase) DeleteReply(replyID, role_, UserID string) (*domain.Reply, error) {
	objID, err := primitive.ObjectIDFromHex(replyID)
	if err != nil {
		return nil, errors.New("invalid reply ID")
	}
	existingReply, err := u.commentRepo.GetReplyByID(objID)
	if err != nil {
		return nil, err
	}
	if existingReply.UserID != UserID || role_ != "admin" {
		return nil, errors.New("unauthorized")
	}

	return u.commentRepo.DeleteReply(objID)
}

func (u *commentUsecase) GetReplies(commentID string, page, limit int) ([]domain.Reply, error) {
	if page <= 0 || limit <= 0 {
		return nil, errors.New("invalid pagination parameters")
	}
	replies, err := u.commentRepo.GetRepliesByCommentID(commentID, int64(page), int64(limit))
	if err != nil {
		return nil, err
	}
	return convertReplies(replies), nil
}

func convertReplies(replies []*domain.Reply) []domain.Reply {
	result := make([]domain.Reply, len(replies))
	for i, r := range replies {
		result[i] = *r
	}
	return result
}

func (u *commentUsecase) LikeComment(commentID, userID string) error {
	objID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return errors.New("invalid comment ID")
	}
	return u.commentRepo.LikeComment(objID, userID)
}

func (u *commentUsecase) UnlikeComment(commentID, userID string) error {
	objID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return errors.New("invalid comment ID")
	}
	return u.commentRepo.UnlikeComment(objID, userID)
}

func (u *commentUsecase) LikeReply(replyID, userID string) error {
	objID, err := primitive.ObjectIDFromHex(replyID)
	if err != nil {
		return errors.New("invalid reply ID")
	}
	return u.commentRepo.LikeReply(objID, userID)
}

func (u *commentUsecase) UnlikeReply(replyID, userID string) error {
	objID, err := primitive.ObjectIDFromHex(replyID)
	if err != nil {
		return errors.New("invalid reply ID")
	}
	return u.commentRepo.UnlikeReply(objID, userID)
}
