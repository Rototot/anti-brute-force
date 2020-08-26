package services

import (
	"testing"

	"github.com/Rototot/anti-brute-force/pkg/domain/constants"
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestBucketRateLimiter_IsLimitExceeded(t *testing.T) {
	t.Run("when bucket is empty", func(t *testing.T) {
		bucket := &entities.Bucket{
			ID:       "test-id",
			Capacity: 10,
		}
		expectedBucket := &entities.Bucket{
			ID:       bucket.ID,
			Capacity: bucket.Capacity,
		}

		bRepostiory := mocks.NewMockBucketRepository(gomock.NewController(t))
		bRepostiory.
			EXPECT().
			CountDrips(expectedBucket).
			Return(5).
			MaxTimes(1)

		bRepostiory.
			EXPECT().
			AddDrip(expectedBucket).
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
		}

		bRepostiory := mocks.NewMockBucketRepository(gomock.NewController(t))
		bRepostiory.
			EXPECT().
			CountDrips(gomock.Any()).
			Return(12).
			MaxTimes(1)

		r := &BucketRateLimiter{
			bucketRepository: bRepostiory,
		}

		result, err := r.IsLimitExceeded(bucket)
		require.True(t, result)
		require.Error(t, constants.ErrAttemptsIsExceeded, err)
	})
}
