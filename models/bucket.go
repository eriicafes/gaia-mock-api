package models

import (
	"errors"
	"time"

	"github.com/eriicafes/filedb"
)

const BucketResource = "buckets"

type Bucket struct {
	ID        filedb.ID   `json:"id"`
	Data      interface{} `json:"data"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
	AccountID *string     `json:"accountId"`
}

type BucketQuery struct {
	ID        int
	AccountID string
}

type bucketModel struct {
	db *filedb.Database
}

func NewBucketModel(db *filedb.Database) *bucketModel {
	return &bucketModel{db}
}

func matchBucketQuery(bucket Bucket, query *BucketQuery) bool {
	if query == nil {
		return true
	}

	matchID := bucket.ID == filedb.ID(query.ID)
	if query.ID == 0 {
		matchID = true
	}

	matchUserID := func() bool {
		if bucket.AccountID != nil {
			return *bucket.AccountID == query.AccountID
		}
		return false
	}()
	if query.AccountID == "" {
		matchUserID = true
	}

	return matchID && matchUserID
}

func (m *bucketModel) get() []Bucket {
	var buckets []Bucket

	m.db.Get(BucketResource, &buckets)

	return buckets
}

func (m *bucketModel) set(buckets []Bucket) {
	data := make([]interface{}, 0, len(buckets))

	for _, bucket := range buckets {
		data = append(data, bucket)
	}

	m.db.Set(BucketResource, data)
}

func (m *bucketModel) FindOneBucket(query *BucketQuery) (*Bucket, error) {
	buckets := m.get()

	for _, bucket := range buckets {
		match := matchBucketQuery(bucket, query)

		if match {
			return &bucket, nil
		}
	}

	return nil, errors.New("bucket not found")
}

func (m *bucketModel) FindManyBuckets(query *BucketQuery) []Bucket {
	buckets := m.get()
	var result []Bucket

	for _, bucket := range buckets {
		match := matchBucketQuery(bucket, query)

		if match {
			result = append(result, bucket)
		}
	}

	return result
}

func (m *bucketModel) CreateBucket(bucket Bucket) *Bucket {
	buckets := m.get()

	// override fields
	bucket.ID = filedb.ID(len(buckets) + 1)
	bucket.CreatedAt = time.Now()
	bucket.UpdatedAt = time.Now()

	buckets = append(buckets, bucket)

	m.set(buckets)

	return &bucket
}

func (m *bucketModel) UpdateBucket(query *BucketQuery, updatedBucket Bucket) (*Bucket, error) {
	buckets := m.get()
	var newBuckets []Bucket

	var updated bool

	for _, bucket := range buckets {
		if updated {
			newBuckets = append(newBuckets, bucket)
			continue
		}

		match := matchBucketQuery(bucket, query)

		if match {
			updated = true

			// override fields
			updatedBucket.ID = bucket.ID
			updatedBucket.CreatedAt = bucket.CreatedAt
			updatedBucket.UpdatedAt = time.Now()
			if updatedBucket.AccountID == nil {
				updatedBucket.AccountID = bucket.AccountID
			}

			newBuckets = append(newBuckets, updatedBucket)
			continue
		}

		newBuckets = append(newBuckets, bucket)
	}

	if !updated {
		return nil, errors.New("bucket not found")
	}

	m.set(newBuckets)
	return &updatedBucket, nil
}

func (m *bucketModel) RemoveOneBucket(query *BucketQuery) error {
	buckets := m.get()
	var newBuckets []Bucket

	var removed bool

	for _, bucket := range buckets {
		if removed {
			newBuckets = append(newBuckets, bucket)
			continue
		}

		match := matchBucketQuery(bucket, query)

		if match {
			removed = true
			continue
		}

		newBuckets = append(newBuckets, bucket)
	}

	if !removed {
		return errors.New("bucket not found")
	}

	m.set(newBuckets)
	return nil
}

func (m *bucketModel) RemoveManyBuckets(query *BucketQuery) (int, error) {
	buckets := m.get()
	var newBuckets []Bucket

	var count int

	for _, bucket := range buckets {
		match := matchBucketQuery(bucket, query)

		if match {
			count++
			continue
		}

		newBuckets = append(newBuckets, bucket)
	}

	if count == 0 {
		return count, errors.New("buckets not found")
	}

	m.set(newBuckets)

	return count, nil
}
