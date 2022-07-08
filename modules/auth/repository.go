package auth

import (
	"github.com/eriicafes/filedb"
	"github.com/eriicafes/go-api-starter/models"
)

type authRepository struct {
	db *filedb.Database
}

func NewAuthRepository(db *filedb.Database) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) FindOne(id int) (*models.User, error) {
	model := models.NewUserModel(r.db)

	return model.FindOneUser(&models.UserQuery{
		ID: id,
	})
}

func (r *authRepository) FindByAccountId(accountId string) (*models.User, error) {
	model := models.NewUserModel(r.db)

	return model.FindOneUser(&models.UserQuery{
		AccountID: accountId,
	})
}

func (r *authRepository) Create(user models.User) *models.User {
	model := models.NewUserModel(r.db)

	return model.CreateUser(user)
}

func (r *authRepository) RemoveByAccountId(accountId string) error {
	model := models.NewUserModel(r.db)

	return model.RemoveOneUser(&models.UserQuery{
		AccountID: accountId,
	})
}
