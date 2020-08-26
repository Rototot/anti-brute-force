package repositories

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/Rototot/anti-brute-force/pkg/domain/valueobjects"

	"github.com/go-redis/redis/v8"

	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
)

type BucketRepository struct {
	client *redis.Client
}

func NewBucketRepository(client *redis.Client) *BucketRepository {
	return &BucketRepository{client: client}
}

func (r *BucketRepository) CountDrips(bucket *entities.Bucket) int {
	keyDrips := r.generateKey(string(bucket.ID), "drips", "*")

	keys, err := r.client.Keys(context.Background(), keyDrips).Result()
	if err != nil {
		return 0
	}

	return len(keys)
}

func (r *BucketRepository) AddDrip(bucket *entities.Bucket) error {
	keyDrip := r.generateKey(string(bucket.ID), "drips", strconv.FormatInt(time.Now().UnixNano(), 10))

	return r.client.Set(context.Background(), keyDrip, time.Now().Unix(), time.Minute).Err()
}

func (r *BucketRepository) Remove(id valueobjects.BucketID) error {
	key := r.generateKey(string(id), "drips", "*")

	return r.client.Del(context.Background(), key).Err()
}

func (r *BucketRepository) generateKey(params ...string) string {
	keyParts := append([]string{"rate:bucket"}, params...)

	return strings.Join(keyParts, ":")
}
