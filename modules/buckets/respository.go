package buckets

import (
	"github.com/eriicafes/filedb"
	"github.com/eriicafes/go-api-starter/models"
)

type bucketsRepository struct {
	db *filedb.Database
}

func NewBucketsRepository(db *filedb.Database) *bucketsRepository {
	return &bucketsRepository{
		db: db,
	}
}

func (r *bucketsRepository) FindAll() []models.Bucket {
	model := models.NewBucketModel(r.db)

	return model.FindManyBuckets(nil)
}

func (r *bucketsRepository) FindOne(accountId string, id int) (*models.Bucket, error) {
	model := models.NewBucketModel(r.db)

	return model.FindOneBucket(&models.BucketQuery{
		AccountID: accountId,
		ID:        id,
	})
}

func (r *bucketsRepository) Create(accountId string, bucket models.Bucket) *models.Bucket {
	model := models.NewBucketModel(r.db)

	bucket.AccountID = &accountId

	return model.CreateBucket(bucket)
}

func (r *bucketsRepository) Update(accountId string, id int, bucket models.Bucket) (*models.Bucket, error) {
	model := models.NewBucketModel(r.db)

	return model.UpdateBucket(&models.BucketQuery{
		AccountID: accountId,
		ID:        id,
	}, bucket)
}

func (r *bucketsRepository) Remove(accountId string, id int) error {
	model := models.NewBucketModel(r.db)

	return model.RemoveOneBucket(&models.BucketQuery{
		AccountID: accountId,
		ID:        id,
	})
}
