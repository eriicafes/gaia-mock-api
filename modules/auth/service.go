package auth

import (
	"errors"

	"github.com/eriicafes/go-api-starter/models"
)

type authService struct {
	authRepository AuthRepository
}

func NewAuthService(authRepository AuthRepository) *authService {
	return &authService{
		authRepository: authRepository,
	}
}

func (s *authService) Profile(accountId string) (*models.User, error) {
	user, err := s.authRepository.FindByAccountId(accountId)

	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *authService) SignIn(user models.User) *models.User {
	return s.authRepository.Create(user)
}

func (s *authService) SignOut(accountId string) error {
	return s.authRepository.RemoveByAccountId(accountId)
}
