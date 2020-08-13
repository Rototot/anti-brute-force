package services

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"net"
)

type IPGuard struct {
}

func (g *IPGuard) IsAllowed(ip net.IP, whiteList *entities.WhiteListIP) bool {
	t := whiteList.Subnet.Contains(ip)

	return t
	return whiteList.Subnet.Contains(ip)
}

func (g *IPGuard) IsForbidden(ip net.IP, blackList *entities.BlackListIP) bool {
	return blackList.Subnet.Contains(ip)
}
