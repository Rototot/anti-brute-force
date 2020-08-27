package usecases //nolint:dupl

import (
	"errors"
	"net"
	"testing"

	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

//nolint:funlen
func TestRemoveIPFromBlackListHandler_Execute(t *testing.T) {
	t.Run("when already exists ", func(t *testing.T) {
		expectedIP := &entities.BlackListIP{}
		ipReposioty := mocks.NewMockBlackListIPRepository(gomock.NewController(t))
		ipReposioty.
			EXPECT().
			FindBySubnet(gomock.Any()).
			Return(expectedIP, nil).
			MinTimes(1)

		ipReposioty.
			EXPECT().
			Remove(expectedIP).
			Return(nil).
			MaxTimes(1)

		handler := NewRemoveIPFromBlackListHandler(ipReposioty)

		_, network, _ := net.ParseCIDR("10.0.0.1/8")

		useCase := RemoveIPFromBlackList{
			Subnet: *network,
		}
		r := handler.Execute(useCase)

		require.Nil(t, r)
	})

	t.Run("when delete exists ", func(t *testing.T) {
		expectedIP := &entities.BlackListIP{}
		ipReposioty := mocks.NewMockBlackListIPRepository(gomock.NewController(t))
		ipReposioty.
			EXPECT().
			FindBySubnet(gomock.Any()).
			Return(expectedIP, nil).
			MinTimes(1)

		ipReposioty.
			EXPECT().
			Remove(expectedIP).
			Return(nil).
			MaxTimes(1)

		handler := NewRemoveIPFromBlackListHandler(ipReposioty)

		_, network, _ := net.ParseCIDR("10.0.0.1/8")

		useCase := RemoveIPFromBlackList{
			Subnet: *network,
		}
		r := handler.Execute(useCase)

		require.Nil(t, r)
	})

	t.Run("when delete error ", func(t *testing.T) {
		expectedIP := &entities.BlackListIP{}
		expectedErr := errors.New("test remove error")
		ipReposioty := mocks.NewMockBlackListIPRepository(gomock.NewController(t))
		ipReposioty.
			EXPECT().
			FindBySubnet(gomock.Any()).
			Return(expectedIP, nil).
			MinTimes(1)

		ipReposioty.
			EXPECT().
			Remove(expectedIP).
			Return(expectedErr).
			MaxTimes(1)

		handler := NewRemoveIPFromBlackListHandler(ipReposioty)

		_, network, _ := net.ParseCIDR("10.0.0.1/8")

		useCase := RemoveIPFromBlackList{
			Subnet: *network,
		}
		r := handler.Execute(useCase)

		require.Error(t, r)
		require.Equal(t, r, expectedErr)
	})
}
