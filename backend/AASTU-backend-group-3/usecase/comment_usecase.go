package usecase

import (
	"group3-blogApi/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type commentUsecase struct {
	commentRepo domain.CommentRepository
	userRepo   domain.UserRepository
}

func NewCommentUsecase(commentRepo domain.CommentRepository, userRepo domain.UserRepository) domain.CommentUsecase {
	return &commentUsecase{
		commentRepo: commentRepo,
		userRepo:   userRepo,
	}
}

func (u *commentUsecase) CreateComment(comment *domain.Comment) (*domain.Comment, *domain.CustomError) {
	if comment.PostID.IsZero() || comment.UserID.IsZero() || comment.Content == "" {
		return nil, domain.ErrMissingRequiredFields

	}
	
	isVerified, verr := u.userRepo.IsVerified(comment.UserID)
	if verr != nil {
		return nil, domain.ErrUserNotVerified
	}
	if !isVerified {
		return nil, domain.ErrUserNotVerified
	}

	createdComment, err := u.commentRepo.CreateComment(comment)
	if err != nil {
		return nil, domain.ErrFailedToCreateComment
	}
	return createdComment, &domain.CustomError{}
}

func (u *commentUsecase) UpdateComment(comment *domain.Comment, role_, userId string) (*domain.Comment, *domain.CustomError) {
	existingComment, err := u.GetCommentByID(comment.ID.Hex())
	if err != nil {
		return nil, err
	}
	if existingComment.UserID.Hex() != userId && role_ != "admin" {
		return nil, domain.ErrUnauthorized
	}
	if comment.ID.IsZero() {
		return nil, domain.ErrInvalidCommentID
	}
	
	isVerified, verr := u.userRepo.IsVerified(comment.UserID)
	if verr != nil {
		return nil, domain.ErrUserNotVerified
	}
	if !isVerified {
		return nil, domain.ErrUserNotVerified
	}

	updatedComment, rerr := u.commentRepo.UpdateComment(comment)
	if rerr != nil {
		return nil, domain.ErrFailedToUpdateComment
	}
	return updatedComment, &domain.CustomError{}
}

func (u *commentUsecase) DeleteComment(commentID, role_, userID string) (*domain.Comment, *domain.CustomError) {
	existingComment, cerr := u.GetCommentByID(commentID)
	if cerr != nil {
		return nil, cerr
	}
	if existingComment.UserID.Hex() != userID && role_ != "admin" {
		return nil, domain.ErrUnauthorized
	}

	objID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return nil, domain.ErrInvalidCommentID
	}
	deletedComment, err := u.commentRepo.DeleteComment(objID)
	if err != nil {
		return nil, domain.ErrFailedToDeleteComment
	}
	return deletedComment, &domain.CustomError{}
}

func (u *commentUsecase) GetCommentByID(commentID string) (*domain.Comment, *domain.CustomError) {
	objID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return nil, domain.ErrInvalidCommentID
	}
	comment, err := u.commentRepo.GetCommentByID(objID)
	if err != nil {
		return nil, domain.ErrFailedToGetComment
	}
	return comment, &domain.CustomError{}
}

func (u *commentUsecase) GetComments(postID string, page, limit int) ([]domain.Comment, *domain.CustomError) {
	if page <= 0 || limit <= 0 {
		return nil, domain.ErrInvalidPaginationParameters
	}
	comments, err := u.commentRepo.GetCommentsByPostID(postID, int64(page), int64(limit))
	if err != nil {
		return nil, domain.ErrFailedToGetComments
	}

	return convertComments(comments), &domain.CustomError{}

}

func convertComments(comments []*domain.Comment) []domain.Comment {
	result := make([]domain.Comment, len(comments))
	for i, c := range comments {
		result[i] = *c
	}
	return result
}



func (u *commentUsecase) CreateReply(reply *domain.Reply) (*domain.Reply, *domain.CustomError) {
	if reply.CommentID.IsZero() || reply.UserID == "" || reply.Content == "" {
		return nil, domain.ErrMissingRequiredFields
	}

	userIdObj, err := primitive.ObjectIDFromHex(reply.UserID)
	if err != nil {
		return nil, domain.ErrInvalidUserID
	}
	
	isVerified, verr := u.userRepo.IsVerified(userIdObj)
	if verr != nil {
		return nil, domain.ErrUserNotVerified
	}
	if !isVerified {
		return nil, domain.ErrUserNotVerified
	}

	createdReply, err := u.commentRepo.CreateReply(reply)
	if err != nil {
		return nil, domain.ErrFailedToCreateReply
	}
	return createdReply, &domain.CustomError{}
}

func (u *commentUsecase) UpdateReply(reply *domain.Reply, userID string) (*domain.Reply, *domain.CustomError) {
	existingReply, err := u.commentRepo.GetReplyByID(reply.ID)
	if err != nil {
		return nil, domain.ErrReplyNotFound
	}
	if existingReply.UserID != userID {
		return nil, domain.ErrUnauthorized
	}
	if reply.ID.IsZero() {
		return nil, domain.ErrInvalidReplyID
	}
	userIdObj, err := primitive.ObjectIDFromHex(reply.UserID)
	if err != nil {
		return nil, domain.ErrInvalidUserID
	}
	
	isVerified, verr := u.userRepo.IsVerified(userIdObj)
	if verr != nil {
		return nil, domain.ErrUserNotVerified
	}
	if !isVerified {
		return nil, domain.ErrUserNotVerified
	}
	updatedReply, err := u.commentRepo.UpdateReply(reply)
	if err != nil {
		return nil, domain.ErrFailedToUpdateReply
	}
	return updatedReply, &domain.CustomError{}
}

func (u *commentUsecase) DeleteReply(replyID, role_, userID string) (*domain.Reply, *domain.CustomError) {
	objID, err := primitive.ObjectIDFromHex(replyID)
	if err != nil {
		return nil, domain.ErrInvalidReplyID
	}
	existingReply, err := u.commentRepo.GetReplyByID(objID)
	if err != nil {
		return nil, domain.ErrReplyNotFound
	}
	if existingReply.UserID != userID && role_ != "admin" {
		return nil, domain.ErrUnauthorized
	}

	deletedReply, err := u.commentRepo.DeleteReply(objID)
	if err != nil {
		return nil, domain.ErrFailedToDeleteReply
	}
	return deletedReply, &domain.CustomError{}
}

func (u *commentUsecase) GetReplies(commentID string, page, limit int) ([]domain.Reply, *domain.CustomError) {
	if page <= 0 || limit <= 0 {
		return nil, domain.ErrInvalidPaginationParameters
	}
	replies, err := u.commentRepo.GetRepliesByCommentID(commentID, int64(page), int64(limit))
	if err != nil {
		return nil, domain.ErrFailedToGetReplies
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

func (u *commentUsecase) LikeComment(commentID, userID string) *domain.CustomError {
	objID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return domain.ErrInvalidCommentID
	}

	isVerified, verr := u.userRepo.IsVerified(objID)
	if verr != nil {
		return domain.ErrUserNotVerified
	}
	if !isVerified {
		return domain.ErrUserNotVerified
	}
	err = u.commentRepo.LikeComment(objID, userID)
	if err != nil {
		return domain.ErrFailedToLikeComment
	}
	return &domain.CustomError{}
}

func (u *commentUsecase) UnlikeComment(commentID, userID string) *domain.CustomError {
	objID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return domain.ErrInvalidCommentID
	}
	
	isVerified, verr := u.userRepo.IsVerified(objID)
	if verr != nil {
		return domain.ErrUserNotVerified
	}
	if !isVerified {
		return domain.ErrUserNotVerified
	}
	err = u.commentRepo.UnlikeComment(objID, userID)
	if err != nil {
		return domain.ErrFailedToUnlikeComment
	}
	return &domain.CustomError{}
}

func (u *commentUsecase) LikeReply(replyID, userID string) *domain.CustomError {
	objID, err := primitive.ObjectIDFromHex(replyID)
	if err != nil {
		return domain.ErrInvalidReplyID
	}
	
	isVerified, verr := u.userRepo.IsVerified(objID)
	if verr != nil {
		return domain.ErrUserNotVerified
	}
	if !isVerified {
		return domain.ErrUserNotVerified
	}
	err = u.commentRepo.LikeReply(objID, userID)
	if err != nil {
		return domain.ErrFailedToLikeReply
	}
	return &domain.CustomError{}
}

func (u *commentUsecase) UnlikeReply(replyID, userID string) *domain.CustomError {
	objID, err := primitive.ObjectIDFromHex(replyID)
	if err != nil {
		return domain.ErrInvalidReplyID
	}
	
	isVerified, verr := u.userRepo.IsVerified(objID)
	if verr != nil {
		return domain.ErrUserNotVerified
	}
	if !isVerified {
		return domain.ErrUserNotVerified
	}
	err = u.commentRepo.UnlikeReply(objID, userID)
	if err != nil {
		return domain.ErrFailedToUnlikeReply
	}
	return &domain.CustomError{}
}
