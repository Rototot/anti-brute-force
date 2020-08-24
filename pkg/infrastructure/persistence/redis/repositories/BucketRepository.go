package repositories

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/valueObjects"
	"github.com/gomodule/redigo/redis"
)

type BucketRepository struct {
	pool *redis.Pool
}

func (r *BucketRepository) FindOneByID(id valueObjects.BucketID) (*entities.Bucket, error) {
	panic("implement me")
}

func (r *BucketRepository) Add(bucket *entities.Bucket) error {
	panic("implement me")
}

func (r *BucketRepository) Update(bucket *entities.Bucket) error {
	panic("implement me")
}

func (r *BucketRepository) Remove(bucket *entities.Bucket) error {
	panic("implement me")
}
