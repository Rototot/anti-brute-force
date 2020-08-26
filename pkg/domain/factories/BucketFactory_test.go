package factories

import (
	"testing"

	"github.com/Rototot/anti-brute-force/pkg/domain/constants"
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/factories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestBucketFactory_Create(t *testing.T) {
	const expectedIPCapacity = 1000
	const expectedLoginCapacity = 10
	const expectedPasswordCapacity = 100

	configurator := mocks.NewMockBucketConfigurator(gomock.NewController(t))
	configurator.EXPECT().IpBucketCapacity().Return(expectedIPCapacity).AnyTimes()
	configurator.EXPECT().LoginBucketCapacity().Return(expectedLoginCapacity).AnyTimes()
	configurator.EXPECT().PasswordBucketCapacity().Return(expectedPasswordCapacity).AnyTimes()

	t.Run("unknown bucket type", func(t *testing.T) {
		f := &BucketFactory{conf: configurator}

		r, err := f.Create(constants.BucketType(100000000))

		require.Nil(t, r)
		require.Error(t, err)
	})

	t.Run("ip bucket", func(t *testing.T) {
		f := &BucketFactory{conf: configurator}
		expected := &entities.Bucket{
			Capacity: expectedIPCapacity,
		}

		r, err := f.Create(constants.BucketTypeIp)

		require.Nil(t, err)
		require.Equal(t, expected, r)
	})

	t.Run("login bucket", func(t *testing.T) {
		f := &BucketFactory{conf: configurator}
		expected := &entities.Bucket{
			Capacity: expectedLoginCapacity,
		}

		r, err := f.Create(constants.BucketTypeLogin)

		require.Nil(t, err)
		require.Equal(t, expected, r)
	})

	t.Run("password bucket", func(t *testing.T) {
		f := &BucketFactory{conf: configurator}
		expected := &entities.Bucket{
			Capacity: expectedPasswordCapacity,
		}

		r, err := f.Create(constants.BucketTypePassword)

		require.Nil(t, err)
		require.Equal(t, expected, r)
	})
}
