package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tachida2k/ai-chat-bot-backend/internal/cache"
	"github.com/tachida2k/ai-chat-bot-backend/internal/database/entities"
	"github.com/tachida2k/ai-chat-bot-backend/internal/database/repositories"
	"github.com/tachida2k/ai-chat-bot-backend/pkg/utils"
)

type AuthService struct {
	UserRepo    *repositories.UserRepository
	RedisClient *cache.RedisClientWrapper
}

func NewAuthService(userRepo *repositories.UserRepository, redis *cache.RedisClientWrapper) *AuthService {
	return &AuthService{
		UserRepo:    userRepo,
		RedisClient: redis,
	}
}

func (s *AuthService) HandleLogin(walletAddress, message, signature string) (sessionID string, user *entities.User, err error) {
	// Get expected nonce
	expectedNonce, err := s.RedisClient.GetNonce(walletAddress)
	if err != nil || expectedNonce == "" {
		return "", nil, errors.New("missing or expired nonce")
	}

	// Verify signature
	recoveredAddress, err := utils.RecoverAddressFromSignature(message, signature)
	if err != nil || !strings.EqualFold(recoveredAddress, walletAddress) {
		return "", nil, errors.New("invalid signature")
	}

	// Create user if not exists
	user, err = s.UserRepo.CreateIfNotExists(walletAddress)
	if err != nil {
		return "", nil, errors.New("failed to create user")
	}

	// Create session ID
	sessionID = uuid.New().String()

	// Store session
	if err := s.RedisClient.SetSession(sessionID, walletAddress, 12*time.Hour); err != nil {
		return "", nil, errors.New("failed to store session")
	}

	return sessionID, user, nil
}

func (s *AuthService) GetUserFromSession(sessionID string) (*entities.User, error) {
	walletAddress, err := s.RedisClient.GetWalletAddress(sessionID)
	if err != nil {
		return nil, fmt.Errorf("invalid or expired session")
	}

	user, err := s.UserRepo.GetByWalletAddress(walletAddress)
	if err != nil || user == nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *AuthService) GetUserByWalletAddress(address string) (*entities.User, error) {
	return s.UserRepo.GetByWalletAddress(address)
}

func (s *AuthService) GenerateAndStoreNonce(walletAddress string) (string, error) {
	nonce, err := utils.GenerateNonce()
	if err != nil {
		return "", err
	}

	err = s.RedisClient.SetNonce(walletAddress, nonce, 5*time.Minute)
	if err != nil {
		return "", err
	}

	return nonce, nil
}
