package entities

import "time"

type Intent struct {
	ID         uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MessageID  uint      `gorm:"column:message_id;not null;index" json:"message_id"`
	Intent     string    `gorm:"column:intent;type:varchar(100);not null" json:"intent"`
	Confidence float64   `gorm:"column:confidence;type:decimal(5,4);not null" json:"confidence"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Intent) TableName() string {
	return "intents"
}
