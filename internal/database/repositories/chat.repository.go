package repositories

import (
	"github.com/otosei-ai/otosei-ai-backend/internal/database/entities"
	"gorm.io/gorm"
)

type ChatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) *ChatRepository {
	return &ChatRepository{db: db}
}

func (r *ChatRepository) GetByChatID(chatID uint) (*entities.Chat, error) {
	var chat entities.Chat
	err := r.db.
		Where("id = ?", chatID).
		First(&chat).Error
	if err != nil {
		return nil, err
	}
	return &chat, nil
}

func (r *ChatRepository) CreateChat(userID uint, title string) (*entities.Chat, error) {
	chat := &entities.Chat{
		UserID: userID,
		Title:  title,
	}
	if err := r.db.Create(chat).Error; err != nil {
		return nil, err
	}
	return chat, nil
}
