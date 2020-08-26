package services

import (
	"errors"
	"net"
	"testing"
	"time"

	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

//nolint:funlen
func TestIPGuard_HacAccess(t *testing.T) {
	whiteListRepository := mocks.NewMockWhiteListIPRepository(gomock.NewController(t))
	whiteListRepository.
		EXPECT().
		FindAll().
		Return([]*entities.WhiteListIP{
			{
				ID: 5,
				Subnet: net.IPNet{
					IP:   net.ParseIP("192.168.8.0"),
					Mask: net.CIDRMask(21, 8*net.IPv4len),
				},
				Created: time.Now(),
			},
		}, nil).
		AnyTimes()

	blackListRepository := mocks.NewMockBlackListIPRepository(gomock.NewController(t))
	blackListRepository.
		EXPECT().
		FindAll().
		Return([]*entities.BlackListIP{
			{
				ID: 7,
				Subnet: net.IPNet{
					IP:   net.ParseIP("10.10.15.0"),
					Mask: net.CIDRMask(25, 8*net.IPv4len),
				},
				Created: time.Now(),
			},
		}, nil).
		AnyTimes()

	t.Run("when ip in whitelist", func(t *testing.T) {
		g := &IPGuard{
			whitelistRepository: whiteListRepository,
			blacklistRepository: blackListRepository,
		}
		r, err := g.HasAccess(net.ParseIP("192.168.8.1"))

		require.Nil(t, err)
		require.Equal(t, true, r)
	})

	t.Run("when ip in blacklist", func(t *testing.T) {
		g := &IPGuard{
			whitelistRepository: whiteListRepository,
			blacklistRepository: blackListRepository,
		}
		r, err := g.HasAccess(net.ParseIP("10.10.15.2"))

		require.Nil(t, err)
		require.Equal(t, false, r)
	})

	t.Run("when ip not in whitelist and not in blacklist", func(t *testing.T) {
		g := &IPGuard{
			whitelistRepository: whiteListRepository,
			blacklistRepository: blackListRepository,
		}
		r, err := g.HasAccess(net.ParseIP("178.168.8.14"))

		require.Nil(t, err)
		require.Equal(t, true, r)
	})

	t.Run("when error", func(t *testing.T) {
		expectedErr := errors.New("test expected repository error")
		whiteListRepository := mocks.NewMockWhiteListIPRepository(gomock.NewController(t))
		whiteListRepository.
			EXPECT().
			FindAll().
			Return([]*entities.WhiteListIP{}, expectedErr).
			AnyTimes()

		blackListRepository := mocks.NewMockBlackListIPRepository(gomock.NewController(t))
		blackListRepository.
			EXPECT().
			FindAll().
			Return([]*entities.BlackListIP{}, expectedErr).
			AnyTimes()

		g := &IPGuard{
			whitelistRepository: whiteListRepository,
			blacklistRepository: blackListRepository,
		}
		r, err := g.HasAccess(net.ParseIP("178.168.8.14"))

		require.Error(t, err)
		require.Equal(t, false, r)
	})
}
