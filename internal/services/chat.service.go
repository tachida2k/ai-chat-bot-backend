package services

import (
	"strings"

	"github.com/otosei-ai/otosei-ai-backend/internal/database/entities"
	"github.com/otosei-ai/otosei-ai-backend/internal/database/repositories"
	"github.com/otosei-ai/otosei-ai-backend/internal/llm/openrouter"
)

type ChatService struct {
	ChatRepo         *repositories.ChatRepository
	MessageRepo      *repositories.MessageRepository
	OpenRouterClient *openrouter.OpenRouterClient
}

func NewChatService(
	chatRepo *repositories.ChatRepository,
	messageRepo *repositories.MessageRepository,
	openrouterClient *openrouter.OpenRouterClient) *ChatService {
	return &ChatService{
		ChatRepo:         chatRepo,
		MessageRepo:      messageRepo,
		OpenRouterClient: openrouterClient,
	}
}

func (s *ChatService) CreateChatIfNotExists(userID uint, chatID uint, userMessage string) (*entities.Chat, error) {
	existingChat, err := s.ChatRepo.GetByChatID(chatID)
	if err == nil {
		return existingChat, nil
	}

	title, err := s.OpenRouterClient.GenerateChatTitle(userMessage)
	if err != nil || strings.TrimSpace(title) == "" {
		title = "Untitled"
	}

	return s.ChatRepo.CreateChat(userID, strings.TrimSpace(title))
}

func (s *ChatService) HandleUserChat(userID uint, chatID *uint, userMessage string) (string, error) {
	var resolvedChatID uint
	if chatID != nil {
		resolvedChatID = *chatID
	}
	chat, err := s.CreateChatIfNotExists(userID, resolvedChatID, userMessage)
	if err != nil {
		return "", err
	}

	err = s.MessageRepo.AddMessage(chat.ID, "user", userMessage)
	if err != nil {
		return "", err
	}

	reply, err := s.OpenRouterClient.Chat([]openrouter.Message{
		{Role: "user", Content: userMessage},
	})
	if err != nil || strings.TrimSpace(reply) == "" {
		return "", err
	}

	err = s.MessageRepo.AddMessageAndTouchChat(chat.ID, "assistant", reply)
	if err != nil {
		return "", err
	}

	return reply, nil
}
