package usecase

import (
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/infrastructure"

	"github.com/google/generative-ai-go/genai"
)




type ChatUseCase struct {
	ChatRepository domain.ChatRepository
	LlmClient *infrastructure.LlmClient
}

func NewChatUseCase(cr domain.ChatRepository,Llc *infrastructure.LlmClient) *ChatUseCase {
	return &ChatUseCase{
		ChatRepository: cr,
		LlmClient: Llc,
	}
}

func (cu *ChatUseCase) CreateChat(userquestion string) (*domain.ChatContext,error) {
	response, err := cu.LlmClient.GenerateText(userquestion,[]genai.Part{})
	if err != nil {
		return nil,err
	}
	chat := domain.ChatContext{
		ChatMessages: []domain.ChatMessage{
			{ Content: userquestion, Role: "user"},
			{ Content: response, Role: "system"},
		},
	}
	return cu.ChatRepository.CreateChat(&chat)
}

func (cu *ChatUseCase) GetChat(id string) (*domain.ChatContext, error) {
	return cu.ChatRepository.GetChat(id)
}

func (cu *ChatUseCase) GetChats() ([]*domain.ChatContext, error) {
	return cu.ChatRepository.GetChats()
}

func (cu *ChatUseCase) UpdateChat(userquestion, id string) (*domain.ChatContext, error) {

	chat, err := cu.ChatRepository.GetChat(id)
	
	if err != nil {
		return nil,err
	}
	chatMessages := []genai.Part{}
	for _, chatMessage := range chat.ChatMessages {
		chatMess := chatMessage.Role + " : " + chatMessage.Content
		chatMessages = append(chatMessages, genai.Text(chatMess))
	}
	response, err := cu.LlmClient.GenerateText(userquestion, chatMessages)
	
	if err != nil {
		return nil,err
	}
	chat.ChatMessages = append(chat.ChatMessages, domain.ChatMessage{Content: userquestion, Role: "user"})
	chat.ChatMessages = append(chat.ChatMessages, domain.ChatMessage{Content: response, Role: "system"})

	return cu.ChatRepository.UpdateChat(chat.ChatMessages, id)
}

func (cu *ChatUseCase) DeleteChat(id string)error{
	return nil
}