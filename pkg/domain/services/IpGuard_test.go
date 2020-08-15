package services

import (
	"errors"
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories"
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net"
	"testing"
	"time"
)

func TestIPGuard_HacAccess(t *testing.T) {
	type fields struct {
		whitelistRepository repositories.WhiteListIPRepository
		blacklistRepository repositories.BlackListIPRepository
	}
	type args struct {
		ip net.IP
	}

	t.Run("no error", func(t *testing.T) {
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

		tests := []struct {
			name    string
			fields  fields
			args    args
			want    bool
			wantErr bool
		}{
			{
				name: "when ip in whitelist",
				fields: fields{
					whitelistRepository: whiteListRepository,
					blacklistRepository: blackListRepository,
				},
				args: args{
					ip: net.ParseIP("192.168.8.1"),
				},
				want: true,
			},
			{
				name: "when ip in blacklist",
				fields: fields{
					whitelistRepository: whiteListRepository,
					blacklistRepository: blackListRepository,
				},
				args: args{
					ip: net.ParseIP("10.10.15.2"),
				},
				want: false,
			},
			{
				name: "when ip not in whitelist and not in blacklist",
				fields: fields{
					whitelistRepository: whiteListRepository,
					blacklistRepository: blackListRepository,
				},
				args: args{
					ip: net.ParseIP("178.168.8.14"),
				},
				want: true,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				g := &IPGuard{
					whitelistRepository: tt.fields.whitelistRepository,
					blacklistRepository: tt.fields.blacklistRepository,
				}
				got, err := g.HasAccess(tt.args.ip)
				if (err != nil) != tt.wantErr {
					t.Errorf("isAllowed() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("isAllowed() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("with error", func(t *testing.T) {
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

		tests := []struct {
			name    string
			fields  fields
			args    args
			want    bool
			wantErr bool
		}{
			{
				name: "when network contain ip",
				fields: fields{
					whitelistRepository: whiteListRepository,
					blacklistRepository: blackListRepository,
				},
				args: args{
					ip: net.ParseIP("192.168.8.1"),
				},
				want:    false,
				wantErr: true,
			},
			{
				name: "when network do not contain ip",
				fields: fields{
					whitelistRepository: whiteListRepository,
					blacklistRepository: blackListRepository,
				},
				args: args{
					ip: net.ParseIP("178.168.8.14"),
				},
				want:    false,
				wantErr: true,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				g := &IPGuard{
					whitelistRepository: tt.fields.whitelistRepository,
					blacklistRepository: tt.fields.blacklistRepository,
				}
				_, err := g.HasAccess(tt.args.ip)
				require.Error(t, err, expectedErr)
			})
		}
	})
}
