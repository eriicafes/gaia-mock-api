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

func (s *bucketsService) GetAll(accountId string) []models.Bucket {
	buckets := s.bucketsRepository.FindAll(accountId)

	return buckets
}

func (s *bucketsService) Get(accountId string, name string) (*models.Bucket, error) {
	buckets := s.bucketsRepository.FindAll(accountId)

	for _, bucket := range buckets {
		if *bucket.AccountID == accountId && bucket.Name == name {
			return &bucket, nil
		}
	}

	return nil, errors.New("bucket not found")
}

func (s *bucketsService) Put(accountId string, name string, bucketData models.Bucket) *models.Bucket {
	buckets := s.bucketsRepository.FindAll(accountId)

	for _, bucket := range buckets {
		if *bucket.AccountID == accountId && bucket.Name == name {
			bucketData.Name = name
			updatedBucket, _ := s.bucketsRepository.Update(accountId, int(bucket.ID), bucketData)
			return updatedBucket
		}
	}

	bucketData.Name = name

	newBucket := s.bucketsRepository.Create(accountId, bucketData)

	return newBucket
}
