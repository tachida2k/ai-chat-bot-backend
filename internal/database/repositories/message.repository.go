package repositories

import (
	"database/sql"

	"github.com/otosei-ai/otosei-ai-backend/internal/database/entities"
)

type MessageRepository struct {
	DB *sql.DB
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{DB: db}
}

func (r *MessageRepository) GetByChatID(chatID int) ([]entities.Message, error) {
	rows, err := r.DB.Query(`
		SELECT id, chat_id, sender, content, created_at, updated_at
		FROM messages
		WHERE chat_id = $1
		ORDER BY created_at ASC
	`, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []entities.Message
	for rows.Next() {
		var m entities.Message
		if err := rows.Scan(&m.ID, &m.ChatID, &m.Sender, &m.Content, &m.CreatedAt, &m.UpdatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}
	return messages, nil
}

func (r *MessageRepository) Insert(chatID int, sender string, content string) error {
	_, err := r.DB.Exec(`
		INSERT INTO messages (chat_id, sender, content)
		VALUES ($1, $2, $3)
	`, chatID, sender, content)
	return err
}
