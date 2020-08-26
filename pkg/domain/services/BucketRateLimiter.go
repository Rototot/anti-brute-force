package services

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories"
)

type BucketRateLimiter struct {
	bucketRepository repositories.BucketRepository
}

func NewBucketRateLimiter(bucketRepository repositories.BucketRepository) *BucketRateLimiter {
	return &BucketRateLimiter{bucketRepository: bucketRepository}
}

func (r *BucketRateLimiter) IsLimitExceeded(bucket *entities.Bucket) (bool, error) {
	qtyDrips := r.bucketRepository.CountDrips(bucket)

	if qtyDrips >= bucket.Capacity {
		return true, nil
	}

	if err := r.bucketRepository.AddDrip(bucket); err != nil {
		return false, err
	}

	return false, nil
}
