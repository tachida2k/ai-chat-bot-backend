package entities

import "time"

type User struct {
	ID            uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	WalletAddress string    `gorm:"column:wallet_address;type:varchar(100);uniqueIndex;not null" json:"wallet_address"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
