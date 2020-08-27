package usecases

import (
	"net"
	"testing"

	"github.com/Rototot/anti-brute-force/pkg/application/usecases/mocks"
	"github.com/Rototot/anti-brute-force/pkg/domain/constants"
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/valueobjects"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

//nolint:funlen
func TestCheckLoginAttemptHandler_Execute(t *testing.T) {
	expectedIPBucket := entities.Bucket{
		ID:       valueobjects.BucketID("192.168.1.1"),
		Capacity: 1000,
	}
	expectedLoginBucket := entities.Bucket{
		ID:       valueobjects.BucketID("test_login"),
		Capacity: 10,
	}

	expectedPasswordBucket := entities.Bucket{
		ID:       valueobjects.BucketID("test_pass"),
		Capacity: 100,
	}

	t.Run("when limited not exceeded", func(t *testing.T) {
		useCase := CheckLoginAttempt{
			IP:       net.ParseIP(string(expectedIPBucket.ID)),
			Login:    string(expectedLoginBucket.ID),
			Password: string(expectedPasswordBucket.ID),
		}

		bFactory := mocks.NewMockbucketFactory(gomock.NewController(t))
		bFactory.EXPECT().Create(gomock.Eq(constants.BucketTypeIp)).Return(&expectedIPBucket, nil).MinTimes(1)
		bFactory.EXPECT().Create(gomock.Eq(constants.BucketTypePassword)).Return(&expectedPasswordBucket, nil).MinTimes(1)
		bFactory.EXPECT().Create(gomock.Eq(constants.BucketTypeLogin)).Return(&expectedLoginBucket, nil).MinTimes(1)

		ipGuard := mocks.NewMockipGuard(gomock.NewController(t))
		ipGuard.EXPECT().HasAccess(useCase.IP).Return(true, nil).MinTimes(1)

		rateLimiter := mocks.NewMockrateLimiter(gomock.NewController(t))
		rateLimiter.EXPECT().IsLimitExceeded(gomock.Eq(&expectedIPBucket)).Return(false, nil).MinTimes(1)
		rateLimiter.EXPECT().IsLimitExceeded(gomock.Eq(&expectedLoginBucket)).Return(false, nil).MinTimes(1)
		rateLimiter.EXPECT().IsLimitExceeded(gomock.Eq(&expectedPasswordBucket)).Return(false, nil).MinTimes(1)

		handler := NewCheckLoginAttemptHandler(bFactory, ipGuard, rateLimiter)

		r := handler.Execute(useCase)

		require.Nil(t, r)
	})

	t.Run("when has not access", func(t *testing.T) { //nolint:dupl
		useCase := CheckLoginAttempt{
			IP:       net.ParseIP(string(expectedIPBucket.ID)),
			Login:    string(expectedLoginBucket.ID),
			Password: string(expectedPasswordBucket.ID),
		}

		bFactory := mocks.NewMockbucketFactory(gomock.NewController(t))
		bFactory.EXPECT().Create(gomock.Eq(constants.BucketTypeIp)).Return(&expectedIPBucket, nil).MinTimes(0)
		bFactory.EXPECT().Create(gomock.Eq(constants.BucketTypePassword)).Return(&expectedPasswordBucket, nil).MinTimes(0)
		bFactory.EXPECT().Create(gomock.Eq(constants.BucketTypeLogin)).Return(&expectedLoginBucket, nil).MinTimes(0)

		ipGuard := mocks.NewMockipGuard(gomock.NewController(t))
		ipGuard.EXPECT().HasAccess(useCase.IP).Return(false, nil).MinTimes(1)

		rateLimiter := mocks.NewMockrateLimiter(gomock.NewController(t))
		rateLimiter.EXPECT().IsLimitExceeded(gomock.Eq(&expectedIPBucket)).Return(false, nil).MinTimes(0)
		rateLimiter.EXPECT().IsLimitExceeded(gomock.Eq(&expectedLoginBucket)).Return(false, nil).MinTimes(0)
		rateLimiter.EXPECT().IsLimitExceeded(gomock.Eq(&expectedPasswordBucket)).Return(false, nil).MinTimes(0)

		handler := NewCheckLoginAttemptHandler(bFactory, ipGuard, rateLimiter)

		r := handler.Execute(useCase)

		require.Error(t, r)
		require.Equal(t, r, constants.ErrAccessDenied)
	})

	t.Run("when limited exceeded by ip", func(t *testing.T) { //nolint:dupl
		useCase := CheckLoginAttempt{
			IP:       net.ParseIP(string(expectedIPBucket.ID)),
			Login:    string(expectedLoginBucket.ID),
			Password: string(expectedPasswordBucket.ID),
		}

		bFactory := mocks.NewMockbucketFactory(gomock.NewController(t))
		bFactory.EXPECT().Create(gomock.Eq(constants.BucketTypeIp)).Return(&expectedIPBucket, nil).MinTimes(1)
		bFactory.EXPECT().Create(gomock.Eq(constants.BucketTypePassword)).Return(&expectedPasswordBucket, nil).MinTimes(1)
		bFactory.EXPECT().Create(gomock.Eq(constants.BucketTypeLogin)).Return(&expectedLoginBucket, nil).MinTimes(1)

		ipGuard := mocks.NewMockipGuard(gomock.NewController(t))
		ipGuard.EXPECT().HasAccess(useCase.IP).Return(true, nil).MinTimes(1)

		rateLimiter := mocks.NewMockrateLimiter(gomock.NewController(t))
		rateLimiter.EXPECT().IsLimitExceeded(gomock.Eq(&expectedIPBucket)).Return(true, nil).MinTimes(1)
		rateLimiter.EXPECT().IsLimitExceeded(gomock.Eq(&expectedLoginBucket)).Return(false, nil).MinTimes(1)
		rateLimiter.EXPECT().IsLimitExceeded(gomock.Eq(&expectedPasswordBucket)).Return(false, nil).MinTimes(1)

		handler := NewCheckLoginAttemptHandler(bFactory, ipGuard, rateLimiter)

		r := handler.Execute(useCase)

		require.Error(t, r)
		require.Equal(t, r, constants.ErrAttemptsIsExceeded)
	})

	t.Run("when limited exceeded by login", func(t *testing.T) { //nolint:dupl
		useCase := CheckLoginAttempt{
			IP:       net.ParseIP(string(expectedIPBucket.ID)),
			Login:    string(expectedLoginBucket.ID),
			Password: string(expectedPasswordBucket.ID),
		}

		bFactory := mocks.NewMockbucketFactory(gomock.NewController(t))
		bFactory.EXPECT().Create(gomock.Eq(constants.BucketTypeIp)).Return(&expectedIPBucket, nil).MinTimes(1)
		bFactory.EXPECT().Create(gomock.Eq(constants.BucketTypePassword)).Return(&expectedPasswordBucket, nil).MinTimes(1)
		bFactory.EXPECT().Create(gomock.Eq(constants.BucketTypeLogin)).Return(&expectedLoginBucket, nil).MinTimes(1)

		ipGuard := mocks.NewMockipGuard(gomock.NewController(t))
		ipGuard.EXPECT().HasAccess(useCase.IP).Return(true, nil).MinTimes(1)

		rateLimiter := mocks.NewMockrateLimiter(gomock.NewController(t))
		rateLimiter.EXPECT().IsLimitExceeded(gomock.Eq(&expectedIPBucket)).Return(false, nil).MinTimes(1)
		rateLimiter.EXPECT().IsLimitExceeded(gomock.Eq(&expectedLoginBucket)).Return(true, nil).MinTimes(1)
		rateLimiter.EXPECT().IsLimitExceeded(gomock.Eq(&expectedPasswordBucket)).Return(false, nil).MinTimes(1)

		handler := NewCheckLoginAttemptHandler(bFactory, ipGuard, rateLimiter)

		r := handler.Execute(useCase)

		require.Error(t, r)
		require.Equal(t, r, constants.ErrAttemptsIsExceeded)
	})

	t.Run("when limited exceeded by password", func(t *testing.T) { //nolint:dupl
		useCase := CheckLoginAttempt{
			IP:       net.ParseIP(string(expectedIPBucket.ID)),
			Login:    string(expectedLoginBucket.ID),
			Password: string(expectedPasswordBucket.ID),
		}

		bFactory := mocks.NewMockbucketFactory(gomock.NewController(t))
		bFactory.EXPECT().Create(gomock.Eq(constants.BucketTypeIp)).Return(&expectedIPBucket, nil).MinTimes(1)
		bFactory.EXPECT().Create(gomock.Eq(constants.BucketTypePassword)).Return(&expectedPasswordBucket, nil).MinTimes(1)
		bFactory.EXPECT().Create(gomock.Eq(constants.BucketTypeLogin)).Return(&expectedLoginBucket, nil).MinTimes(1)

		ipGuard := mocks.NewMockipGuard(gomock.NewController(t))
		ipGuard.EXPECT().HasAccess(useCase.IP).Return(true, nil).MinTimes(1)

		rateLimiter := mocks.NewMockrateLimiter(gomock.NewController(t))
		rateLimiter.EXPECT().IsLimitExceeded(gomock.Eq(&expectedIPBucket)).Return(false, nil).MinTimes(1)
		rateLimiter.EXPECT().IsLimitExceeded(gomock.Eq(&expectedLoginBucket)).Return(false, nil).MinTimes(1)
		rateLimiter.EXPECT().IsLimitExceeded(gomock.Eq(&expectedPasswordBucket)).Return(true, nil).MinTimes(1)

		handler := NewCheckLoginAttemptHandler(bFactory, ipGuard, rateLimiter)

		r := handler.Execute(useCase)

		require.Error(t, r)
		require.Equal(t, r, constants.ErrAttemptsIsExceeded)
	})

	t.Run("when limited exceeded by all", func(t *testing.T) { //nolint:dupl
		useCase := CheckLoginAttempt{
			IP:       net.ParseIP(string(expectedIPBucket.ID)),
			Login:    string(expectedLoginBucket.ID),
			Password: string(expectedPasswordBucket.ID),
		}

		bFactory := mocks.NewMockbucketFactory(gomock.NewController(t))
		bFactory.EXPECT().Create(gomock.Eq(constants.BucketTypeIp)).Return(&expectedIPBucket, nil).MinTimes(1)
		bFactory.EXPECT().Create(gomock.Eq(constants.BucketTypePassword)).Return(&expectedPasswordBucket, nil).MinTimes(1)
		bFactory.EXPECT().Create(gomock.Eq(constants.BucketTypeLogin)).Return(&expectedLoginBucket, nil).MinTimes(1)

		ipGuard := mocks.NewMockipGuard(gomock.NewController(t))
		ipGuard.EXPECT().HasAccess(useCase.IP).Return(true, nil).MinTimes(1)

		rateLimiter := mocks.NewMockrateLimiter(gomock.NewController(t))
		rateLimiter.EXPECT().IsLimitExceeded(gomock.Eq(&expectedIPBucket)).Return(true, nil).MinTimes(1)
		rateLimiter.EXPECT().IsLimitExceeded(gomock.Eq(&expectedLoginBucket)).Return(true, nil).MinTimes(1)
		rateLimiter.EXPECT().IsLimitExceeded(gomock.Eq(&expectedPasswordBucket)).Return(true, nil).MinTimes(1)

		handler := NewCheckLoginAttemptHandler(bFactory, ipGuard, rateLimiter)

		r := handler.Execute(useCase)

		require.Error(t, r)
		require.Equal(t, r, constants.ErrAttemptsIsExceeded)
	})
}
