package services

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"net"
	"testing"
)

func TestIpGuard_IsAllowed(t *testing.T) {
	type args struct {
		ip        net.IP
		whiteList *entities.WhiteListIP
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "when network contain ip",
			args: args{
				ip: net.ParseIP("192.168.8.1"),
				whiteList: &entities.WhiteListIP{
					//192.168.8.0/21
					Subnet: net.IPNet{
						IP:   net.ParseIP("192.168.8.0"),
						Mask: net.CIDRMask(21, 8*net.IPv4len),
					},
				},
			},
			want: true,
		},
		{
			name: "when network do not contain ip",
			args: args{
				ip: net.ParseIP("178.168.8.14"),
				whiteList: &entities.WhiteListIP{
					//192.168.8.0/21
					Subnet: net.IPNet{
						IP:   net.ParseIP("192.168.8.0"),
						Mask: net.CIDRMask(21, 8*net.IPv4len),
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &IPGuard{}
			if got := g.IsAllowed(tt.args.ip, tt.args.whiteList); got != tt.want {
				t.Errorf("IsAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIpGuard_IsForbidden1(t *testing.T) {
	type args struct {
		ip        net.IP
		blackList *entities.BlackListIP
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "when network contain ip",
			args: args{
				ip: net.ParseIP("192.168.8.14"),
				blackList: &entities.BlackListIP{
					//192.168.8.0/21
					Subnet: net.IPNet{
						IP:   net.ParseIP("192.168.11.10"),
						Mask: net.CIDRMask(21, 8*net.IPv4len),
					},
				},
			},
			want: true,
		},
		{
			name: "when network do not contain ip",
			args: args{
				ip: net.ParseIP("178.168.8.14"),
				blackList: &entities.BlackListIP{
					//192.168.8.0/21
					Subnet: net.IPNet{
						IP:   net.ParseIP("192.168.11.10"),
						Mask: net.CIDRMask(21, 8*net.IPv4len),
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &IPGuard{}
			if got := g.IsForbidden(tt.args.ip, tt.args.blackList); got != tt.want {
				t.Errorf("IsForbidden() = %v, want %v", got, tt.want)
			}
		})
	}
}
