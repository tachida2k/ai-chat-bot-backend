package entities

import "time"

type Action struct {
	ID         uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	IntentID   uint      `gorm:"column:intent_id;not null;index" json:"intent_id"`
	ActionType string    `gorm:"column:action_type;type:varchar(50);not null" json:"action_type"`
	Payload    string    `gorm:"column:payload;type:json" json:"payload"` // JSON string
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Action) TableName() string {
	return "actions"
}
