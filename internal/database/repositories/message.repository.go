package repositories

import (
	"github.com/tachida2k/ai-chat-bot-backend/internal/database/entities"
	"gorm.io/gorm"
)

type MessageRepository struct {
	DB *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{DB: db}
}

func (r *MessageRepository) GetByChatID(chatID uint) ([]entities.Message, error) {
	var messages []entities.Message
	err := r.DB.
		Where("chat_id = ?", chatID).
		Order("created_at ASC").
		Find(&messages).Error
	return messages, err
}

func (r *MessageRepository) AddMessage(chatID uint, sender string, content string) error {
	newMessage := &entities.Message{
		ChatID:  chatID,
		Sender:  sender,
		Content: content,
	}

	return r.DB.Create(newMessage).Error
}

func (r *MessageRepository) AddMessageAndTouchChat(chatID uint, sender string, content string) error {
	tx := r.DB.Begin()
	newMessage := &entities.Message{
		ChatID:  chatID,
		Sender:  sender,
		Content: content,
	}

	if err := r.DB.Create(newMessage).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&entities.Chat{}).
		Where("id = ?", chatID).
		Update("updated_at", gorm.Expr("NOW()")).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
