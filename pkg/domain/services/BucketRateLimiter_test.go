package services

import (
	"errors"
	"testing"

	"github.com/Rototot/anti-brute-force/pkg/domain/constants"
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

//nolint:funlen
func TestBucketRateLimiter_IsLimitExceeded(t *testing.T) {
	t.Run("when is ok", func(t *testing.T) {
		bucket := &entities.Bucket{
			ID:       "test-id",
			Capacity: 10,
			Drips:    5,
		}
		expectedBucket := &entities.Bucket{
			ID:       bucket.ID,
			Capacity: bucket.Capacity,
			Drips:    bucket.Drips + 1,
		}

		bRepostiory := mocks.NewMockBucketRepository(gomock.NewController(t))
		bRepostiory.
			EXPECT().
			Update(expectedBucket).
			Return(nil).
			MaxTimes(1)

		r := &BucketRateLimiter{
			bucketRepository: bRepostiory,
		}

		result, err := r.IsLimitExceeded(bucket)
		require.False(t, result)
		require.Nil(t, err)
	})

	t.Run("when bucket is full", func(t *testing.T) {
		bucket := &entities.Bucket{
			ID:       "test-id",
			Capacity: 10,
			Drips:    10,
		}

		bRepostiory := mocks.NewMockBucketRepository(gomock.NewController(t))
		bRepostiory.
			EXPECT().
			Update(gomock.Any()).
			Return(nil).
			MaxTimes(0)

		r := &BucketRateLimiter{
			bucketRepository: bRepostiory,
		}

		result, err := r.IsLimitExceeded(bucket)
		require.True(t, result)
		require.Error(t, constants.ErrAttemptsIsExceeded, err)
	})

	t.Run("when error on update bucket is full", func(t *testing.T) {
		bucket := &entities.Bucket{
			ID:       "test-id",
			Capacity: 10,
			Drips:    5,
		}

		expectedError := errors.New("test update error")

		bRepostiory := mocks.NewMockBucketRepository(gomock.NewController(t))
		bRepostiory.
			EXPECT().
			Update(gomock.Any()).
			Return(expectedError).
			MaxTimes(1)

		r := &BucketRateLimiter{
			bucketRepository: bRepostiory,
		}

		result, err := r.IsLimitExceeded(bucket)
		require.False(t, result)
		require.Error(t, constants.ErrAttemptsIsExceeded, err)
	})
}

//nolint:funlen
func TestBucketRateLimiter_Clean(t *testing.T) {
	t.Run("when is ok", func(t *testing.T) {
		bucket := &entities.Bucket{
			ID:       "test-id",
			Capacity: 10,
			Drips:    5,
		}

		bRepostiory := mocks.NewMockBucketRepository(gomock.NewController(t))
		bRepostiory.
			EXPECT().
			Remove(bucket).
			Return(nil).
			MaxTimes(1)

		r := &BucketRateLimiter{
			bucketRepository: bRepostiory,
		}

		err := r.Clean(bucket)
		require.Nil(t, err)
	})

	t.Run("when bucket remove return error", func(t *testing.T) {
		bucket := &entities.Bucket{
			ID:       "test-id",
			Capacity: 10,
			Drips:    5,
		}

		bRepostiory := mocks.NewMockBucketRepository(gomock.NewController(t))
		bRepostiory.
			EXPECT().
			Remove(bucket).
			MaxTimes(1)

		r := &BucketRateLimiter{
			bucketRepository: bRepostiory,
		}

		err := r.Clean(bucket)
		require.Nil(t, err)
	})

	t.Run("when bucket is empty", func(t *testing.T) {
		bRepostiory := mocks.NewMockBucketRepository(gomock.NewController(t))
		bRepostiory.
			EXPECT().
			Remove(gomock.Any()).
			Return(nil).
			MaxTimes(0)

		r := &BucketRateLimiter{
			bucketRepository: bRepostiory,
		}

		bucket := &entities.Bucket{
			Capacity: 10,
			Drips:    0,
		}
		err := r.Clean(bucket)
		require.Nil(t, err)
	})

	t.Run("when bucket is nil", func(t *testing.T) {
		bRepostiory := mocks.NewMockBucketRepository(gomock.NewController(t))
		bRepostiory.
			EXPECT().
			Remove(gomock.Any()).
			Return(nil).
			MaxTimes(0)

		r := &BucketRateLimiter{
			bucketRepository: bRepostiory,
		}

		err := r.Clean(nil)
		require.Nil(t, err)
	})
}
