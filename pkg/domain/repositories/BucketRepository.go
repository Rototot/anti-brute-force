package repositories

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/valueobjects"
)

//go:generate mockgen -source=$GOFILE -destination=./mocks/MockBucketRepository.go -package=mocks

type BucketRepository interface {
	AddDrip(bucket *entities.Bucket) error
	CountDrips(bucket *entities.Bucket) int
	Remove(id valueobjects.BucketID) error
}
