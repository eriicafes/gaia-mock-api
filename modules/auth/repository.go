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
	bucketModel := models.NewBucketModel(r.db)

	user, err := model.FindOneUser(&models.UserQuery{
		AccountID: accountId,
	})

	if err != nil {
		return err
	}

	model.RemoveOneUser(&models.UserQuery{
		AccountID: user.AccountID,
	})

	err = bucketModel.RemoveOneBucket(&models.BucketQuery{
		UserID: int(user.ID),
	})

	if err != nil {
		return err
	}

	return nil
}
