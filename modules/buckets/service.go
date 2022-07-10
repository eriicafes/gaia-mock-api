package buckets

import (
	"errors"

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

func (s *bucketsService) GetAll() []models.Bucket {
	buckets := s.bucketsRepository.FindAll()

	return buckets
}

func (s *bucketsService) Get(accountId string) (*models.Bucket, error) {
	buckets := s.bucketsRepository.FindAll()

	for _, bucket := range buckets {
		if *bucket.AccountID == accountId {
			return &bucket, nil
		}
	}

	return nil, errors.New("bucket not found")
}

func (s *bucketsService) Put(accountId string, bucketData models.Bucket) *models.Bucket {
	buckets := s.bucketsRepository.FindAll()

	for _, bucket := range buckets {
		if *bucket.AccountID == accountId {
			updatedBucket, _ := s.bucketsRepository.Update(accountId, int(bucket.ID), bucketData)
			return updatedBucket
		}
	}

	newBucket := s.bucketsRepository.Create(accountId, bucketData)

	return newBucket
}
