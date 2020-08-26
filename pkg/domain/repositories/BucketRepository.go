package repositories

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/valueobjects"
)

//go:generate mockgen -source=$GOFILE -destination=./mocks/MockBucketRepository.go -package=mocks

type BucketRepository interface {
	FindOneByID(id valueobjects.BucketID) (*entities.Bucket, error)

	Add(bucket *entities.Bucket) error
	Update(bucket *entities.Bucket) error
	Remove(bucket *entities.Bucket) error
}
