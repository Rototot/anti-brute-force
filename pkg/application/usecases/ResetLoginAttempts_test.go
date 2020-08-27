package usecases

import (
	"errors"
	"net"
	"testing"

	"github.com/Rototot/anti-brute-force/pkg/domain/repositories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestResetLoginAttemptsHandler_Execute(t *testing.T) {
	t.Run("when ok", func(t *testing.T) {
		useCase := ResetLoginAttempts{
			IP:    net.ParseIP("192.168.1.1"),
			Login: "test_login",
		}
		bRepostiory := mocks.NewMockBucketRepository(gomock.NewController(t))
		bRepostiory.
			EXPECT().
			Remove(gomock.Any()).
			Return(nil).
			MaxTimes(2)

		handler := NewResetLoginAttemptsHandler(bRepostiory)

		r := handler.Execute(useCase)

		require.Nil(t, r)
	})

	t.Run("when remove with error", func(t *testing.T) {
		useCase := ResetLoginAttempts{
			IP:    net.ParseIP("192.168.1.1"),
			Login: "test_login",
		}
		expectedErr := errors.New("test_remove_error")
		bRepostiory := mocks.NewMockBucketRepository(gomock.NewController(t))
		bRepostiory.
			EXPECT().
			Remove(gomock.Any()).
			Return(expectedErr).
			MaxTimes(2)

		handler := NewResetLoginAttemptsHandler(bRepostiory)

		r := handler.Execute(useCase)

		require.Error(t, r)
		require.Equal(t, r, expectedErr)
	})
}
