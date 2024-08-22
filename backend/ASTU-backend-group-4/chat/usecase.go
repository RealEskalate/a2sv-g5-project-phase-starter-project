package chat

import (
	"context"
	"time"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/pkg/infrastructure"
	"github.com/go-playground/validator/v10"
)

type ChatUsecase struct {
	Repository Repository
	AIService  AIService
}

var validate = validator.New()

func NewUsecase(repository Repository, aiService AIService) *ChatUsecase {
	return &ChatUsecase{
		Repository: repository,
		AIService:  aiService,
	}
}

func (usecase *ChatUsecase) CreateChat(ctx context.Context, form CreateChatForm) (Chat, error) {
	err := infrastructure.Validate(validate, form)
	if err != nil{
		return Chat{}, err
	}

	newChat := Chat{
		Title: form.Title,
		UserID: form.UserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newChat, err = usecase.Repository.CreateChat(ctx, newChat)
	if err != nil{
		return Chat{}, err
	}

	return newChat, nil
}

func (usecase *ChatUsecase) DeleteChat(ctx context.Context, form DefaultChatForm) error{
	if err := infrastructure.Validate(validate, form); err != nil{
		return err
	}	

	chat, err := usecase.Repository.GetChat(ctx, form.ChatID)
	if err != nil{
		return err
	}

	if chat.UserID != form.UserID{
		return ErrChatNotFound
	}

	return usecase.Repository.DeleteChat(ctx, form.ChatID)
}

func (usecase *ChatUsecase) GenerateChatTitle(ctx context.Context, form TextForm) (string, error){
	if err := infrastructure.Validate(validate, form); err != nil{
		return "", err
	}

	return usecase.AIService.GenerateChatTitle(ctx, form.Text)
}

func (usecase *ChatUsecase) GetChat(ctx context.Context, form DefaultChatForm) (Chat, error){
	if err := infrastructure.Validate(validate, form); err != nil{
		return Chat{}, err
	}

	chat, err := usecase.Repository.GetChat(ctx, form.ChatID)
	if err != nil{
		return Chat{}, err
	}

	if chat.ID != form.UserID{
		return Chat{}, ErrChatNotFound
	}

	return chat, nil
}

func (usecase *ChatUsecase) GetChats(ctx context.Context, form UserIDForm, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[Chat], error){
	if err := infrastructure.Validate(validate, form); err != nil{
		return infrastructure.PaginationResponse[Chat]{}, err
	}

	if pagination.Limit == 0{
		pagination.Limit = 10
	}
	if pagination.Page == 0{
		pagination.Page = 1
	}
	
	return usecase.Repository.GetChats(ctx, form.UserID, pagination)
}

func (usecase *ChatUsecase) SendMessage(ctx context.Context, chatForm DefaultChatForm, textForm TextForm) (Message, error){
	if err := infrastructure.Validate(validate, chatForm); err != nil{
		return Message{}, err
	}

	if err := infrastructure.Validate(validate, textForm); err != nil{
		return Message{}, err
	}
	
	chat, err := usecase.Repository.GetChat(ctx, chatForm.ChatID)
	if err != nil{
		return Message{}, err
	}

	if chat.UserID != chatForm.UserID{
		return Message{}, ErrChatNotFound
	}

	message := Message{
		Text: textForm.Text,
		Role: "user",
		SentAt: time.Now(),
	}

	response, err := usecase.AIService.SendMessage(ctx, chat.History, message)
	if err != nil{
		return Message{}, err
	}

	if err := usecase.Repository.AddMessage(ctx, chatForm.ChatID, message); err != nil{
		return Message{}, err
	}

	if err := usecase.Repository.AddMessage(ctx, chatForm.ChatID, response); err != nil{
		return Message{}, err
	}

	return response, nil
}







