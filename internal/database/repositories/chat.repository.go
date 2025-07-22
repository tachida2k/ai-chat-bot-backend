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
