package repositories

import (
	"strings"

	"github.com/tachida2k/ai-chat-bot-backend/internal/database/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetByWalletAddress(walletAddress string) (*entities.User, error) {
	normWalletAddress := strings.ToLower(walletAddress)
	var user entities.User
	err := r.DB.Where("wallet_address = ?", normWalletAddress).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepository) CreateIfNotExists(walletAddress string) (*entities.User, error) {
	normWalletAddress := strings.ToLower(walletAddress)
	user, err := r.GetByWalletAddress(normWalletAddress)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}

	newUser := &entities.User{
		WalletAddress: normWalletAddress,
		UserType:      "user",
	}
	err = r.DB.Create(newUser).Error
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
