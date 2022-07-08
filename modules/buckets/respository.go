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

func (r *bucketsRepository) FindAll(userId int) []models.Bucket {
	model := models.NewBucketModel(r.db)

	return model.FindManyBuckets(&models.BucketQuery{
		UserID: userId,
	})
}

func (r *bucketsRepository) FindOne(userId int, id int) (*models.Bucket, error) {
	model := models.NewBucketModel(r.db)

	return model.FindOneBucket(&models.BucketQuery{
		UserID: userId,
		ID:     id,
	})
}

func (r *bucketsRepository) Create(userId int, bucket models.Bucket) *models.Bucket {
	model := models.NewBucketModel(r.db)

	foreignId := filedb.ID(userId)
	bucket.UserID = &foreignId

	return model.CreateBucket(bucket)
}

func (r *bucketsRepository) Update(userId int, id int, bucket models.Bucket) (*models.Bucket, error) {
	model := models.NewBucketModel(r.db)

	return model.UpdateBucket(&models.BucketQuery{
		UserID: userId,
		ID:     id,
	}, bucket)
}

func (r *bucketsRepository) Remove(userId int, id int) error {
	model := models.NewBucketModel(r.db)

	return model.RemoveOneBucket(&models.BucketQuery{
		UserID: userId,
		ID:     id,
	})
}
