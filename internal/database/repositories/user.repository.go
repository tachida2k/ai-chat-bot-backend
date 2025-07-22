package repositories

import (
	"github.com/otosei-ai/otosei-ai-backend/internal/database/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetByWalletAddress(walletAddress string) (*entities.User, error) {
	var user entities.User
	err := r.DB.Where("wallet_address = ?", walletAddress).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepository) CreateIfNotExists(walletAddress string) (*entities.User, error) {
	user, err := r.GetByWalletAddress(walletAddress)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}

	newUser := &entities.User{
		WalletAddress: walletAddress,
	}
	err = r.DB.Create(newUser).Error
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
