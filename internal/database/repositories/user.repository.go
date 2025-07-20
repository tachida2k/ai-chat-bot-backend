package repositories

import (
	"database/sql"

	"github.com/otosei-ai/otosei-ai-backend/internal/database/entities"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetByWalletAddress(walletAddress string) (*entities.User, error) {
	var user entities.User
	err := r.DB.QueryRow(`
		SELECT id, wallet_address, created_at, updated_at 
		FROM users WHERE wallet_address = $1
	`, walletAddress).Scan(&user.ID, &user.WalletAddress, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepository) CreateIfNotExists(walletAddress string) (*entities.User, error) {
	existing, err := r.GetByWalletAddress(walletAddress)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return existing, nil
	}

	var user entities.User
	err = r.DB.QueryRow(`
		INSERT INTO users (wallet_address) VALUES ($1)
		RETURNING id, wallet_address, created_at, updated_at
	`, walletAddress).Scan(&user.ID, &user.WalletAddress, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
