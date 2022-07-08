package auth

import "github.com/eriicafes/go-api-starter/models"

type AuthService interface {
	Profile(accountId string) (*models.User, error)
	SignIn(user models.User) *models.User
	SignOut(accountId string) error
}

type AuthRepository interface {
	FindOne(id int) (*models.User, error)
	FindByAccountId(accountId string) (*models.User, error)
	Create(user models.User) *models.User
	RemoveByAccountId(accountId string) error
}
