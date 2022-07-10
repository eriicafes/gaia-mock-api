package buckets

import "github.com/eriicafes/go-api-starter/models"

type BucketsService interface {
	GetAll() []models.Bucket
	Get(accountId string) (*models.Bucket, error)
	Put(accountId string, bucketData models.Bucket) *models.Bucket
}

type BucketsRepository interface {
	FindAll() []models.Bucket
	FindOne(accountId string, id int) (*models.Bucket, error)
	Create(accountId string, bucket models.Bucket) *models.Bucket
	Update(accountId string, id int, bucket models.Bucket) (*models.Bucket, error)
	Remove(accountId string, id int) error
}
