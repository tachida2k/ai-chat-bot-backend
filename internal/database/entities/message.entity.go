package entities

import "time"

type Message struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ChatID    uint      `gorm:"column:chat_id;not null;index" json:"chat_id"`
	Sender    string    `gorm:"column:sender;type:varchar(20);not null" json:"sender"` // user | assistant
	Content   string    `gorm:"column:content;type:text;not null" json:"content"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Message) TableName() string {
	return "messages"
}
