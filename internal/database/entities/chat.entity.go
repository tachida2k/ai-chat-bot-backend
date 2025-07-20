package entities

import "time"

type Chat struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"column:user_id;not null;index" json:"user_id"`
	Title     string    `gorm:"column:title;type:varchar(255)" json:"title"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Chat) TableName() string {
	return "chats"
}
