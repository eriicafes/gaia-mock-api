package buckets

import (
	"errors"

	"github.com/eriicafes/filedb"
	"github.com/eriicafes/go-api-starter/models"
)

type bucketsService struct {
	bucketsRepository BucketsRepository
}

func NewBucketsService(bucketsRepository BucketsRepository) *bucketsService {
	return &bucketsService{
		bucketsRepository: bucketsRepository,
	}
}

func (s *bucketsService) Get(userId int) (*models.Bucket, error) {
	buckets := s.bucketsRepository.FindAll(userId)

	for _, bucket := range buckets {
		if *bucket.UserID == filedb.ID(userId) {
			return &bucket, nil
		}
	}

	return nil, errors.New("bucket not found")
}

func (s *bucketsService) Put(userId int, bucketData models.Bucket) *models.Bucket {
	buckets := s.bucketsRepository.FindAll(userId)

	for _, bucket := range buckets {
		if *bucket.UserID == filedb.ID(userId) {
			updatedBucket, _ := s.bucketsRepository.Update(userId, int(bucket.ID), bucketData)
			return updatedBucket
		}
	}

	newBucket := s.bucketsRepository.Create(userId, bucketData)

	return newBucket
}
