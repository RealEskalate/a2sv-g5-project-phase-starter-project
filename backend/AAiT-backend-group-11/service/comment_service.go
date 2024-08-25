package service

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"errors"
)

type commentService struct {
	commentRepository interfaces.CommentRepository
	blogRepository interfaces.BlogRepository
	userRepository interfaces.UserRepository
}

func NewCommentService(cr interfaces.CommentRepository, br interfaces.BlogRepository, ur interfaces.UserRepository) interfaces.CommentService {
	return &commentService{
		commentRepository: cr,
		blogRepository: br,
		userRepository: ur,
	}
}

func (cs *commentService) AddComment(comment *entities.Comment) (*entities.Comment, error) {
	// Check if the user exists by authorId
	userExists, err := cs.userRepository.FindUserById(comment.AuthorID.Hex())

	if userExists == nil || err != nil {
		return nil, errors.New("user not found")
	}

	// Check if the blog post exists by blogPostId
	blogExists, err := cs.blogRepository.GetBlogPostById(comment.BlogPostID.Hex())


	if blogExists == nil || err != nil{
		return nil, errors.New("blog post not found")
	}

	res,err :=  cs.commentRepository.AddComment(comment)
	if err != nil{
		return nil, err
	}
	cs.blogRepository.ChangeCommentCount(comment.BlogPostID.Hex(), 1)

	return res,nil
}


func (cs *commentService) DeleteComment( commentId string) error {
	comment,err := cs.commentRepository.GetCommentById(commentId) 
	
	err = cs.commentRepository.DeleteComment(commentId)
	if err != nil{
		return err
	}
	err = cs.blogRepository.ChangeCommentCount(comment.BlogPostID.Hex(), -1)
	if err != nil{
		return err
	}
	return nil

}

func (cs *commentService) GetCommentsByBlogPostId( blogPostId string) ([]entities.Comment, error) {
	return cs.commentRepository.GetCommentsByBlogPostId( blogPostId)
}

func (cs *commentService) UpdateComment( comment *entities.Comment) (*entities.Comment, error) {
	
	return cs.commentRepository.UpdateComment( comment)
}
