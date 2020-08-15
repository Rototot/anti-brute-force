package services

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/constants"
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories"
)

type BucketRateLimiter struct {
	bucketRepository repositories.BucketRepository
}

func (r *BucketRateLimiter) IsLimitExceeded(bucket *entities.Bucket) (bool, error) {
	if bucket.IsFull() {
		return true, constants.ErrAttemptsIsExceeded
	}

	bucket.AddDrips(1)
	err := r.bucketRepository.Update(bucket)
	if err != nil {
		return false, err
	}

	return false, nil
}

func (r *BucketRateLimiter) Clean(bucket *entities.Bucket) error {
	if bucket == nil || bucket.IsEmpty() {
		return nil
	}

	err := r.bucketRepository.Remove(bucket)
	if err != nil {
		return err
	}

	return nil
}
