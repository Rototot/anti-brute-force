package usecases

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net"
	"testing"
)

func TestAddIPToWhitelistHandler_Execute(t *testing.T) {
	t.Run("when already exists ", func(t *testing.T) {
		ipReposioty := mocks.NewMockWhiteListIPRepository(gomock.NewController(t))
		ipReposioty.
			EXPECT().
			FindOneBySubnet(gomock.Any()).
			Return(&entities.WhiteListIP{}, nil).
			MinTimes(1)

		ipReposioty.
			EXPECT().
			Add(gomock.Any()).
			Return(nil).
			MaxTimes(0)

		handler := NewAddIPToWhiteListHandler(ipReposioty)

		_, network, _ := net.ParseCIDR("10.0.0.1/8")

		useCase := AddIPToWhiteList{
			Subnet: *network,
		}
		r := handler.Execute(useCase)

		require.Nil(t, r)
	})

	t.Run("when create exists ", func(t *testing.T) {
		ipReposioty := mocks.NewMockWhiteListIPRepository(gomock.NewController(t))
		ipReposioty.
			EXPECT().
			FindOneBySubnet(gomock.Any()).
			Return(nil, nil).
			MinTimes(1)

		ipReposioty.
			EXPECT().
			Add(gomock.Any()).
			Return(nil).
			MaxTimes(1)

		handler := NewAddIPToWhiteListHandler(ipReposioty)

		_, network, _ := net.ParseCIDR("10.0.0.1/8")

		useCase := AddIPToWhiteList{
			Subnet: *network,
		}
		r := handler.Execute(useCase)

		require.Nil(t, r)
	})
}
