package buckets

import "github.com/eriicafes/go-api-starter/models"

type BucketsService interface {
	Get(userId int) (*models.Bucket, error)
	Put(userId int, bucket models.Bucket) *models.Bucket
}

type BucketsRepository interface {
	FindAll(userId int) []models.Bucket
	FindOne(userId int, id int) (*models.Bucket, error)
	Create(userId int, bucket models.Bucket) *models.Bucket
	Update(userId int, id int, bucket models.Bucket) (*models.Bucket, error)
	Remove(userId int, id int) error
}
