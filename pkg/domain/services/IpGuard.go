package services

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories"
	"net"
)

type IPGuard struct {
	whitelistRepository repositories.WhiteListIPRepository
	blacklistRepository repositories.BlackListIPRepository
}

func (g *IPGuard) HasAccess(ip net.IP) (bool, error) {
	isForbidden, err := g.isForbidden(ip)
	if err != nil {
		return false, err
	}

	// access forbidden
	if isForbidden {
		return false, nil
	}

	isAllowed, err := g.isAllowed(ip)
	if err != nil {
		return false, err
	}

	// access granted
	if isAllowed {
		return true, nil
	}

	// if not exists in blacklist and not exists in white lists, than access granted
	return true, nil
}

func (g *IPGuard) isAllowed(ip net.IP) (bool, error) {
	lists, err := g.whitelistRepository.FindAll()
	if err != nil {
		return false, err
	}

	var isAllowed bool
	for _, list := range lists {
		if list.Subnet.Contains(ip) {
			isAllowed = true
			break
		}
	}

	return isAllowed, nil
}

func (g *IPGuard) isForbidden(ip net.IP) (bool, error) {
	lists, err := g.blacklistRepository.FindAll()
	if err != nil {
		return false, err
	}

	var IsForbidden bool
	for _, list := range lists {
		if list.Subnet.Contains(ip) {
			IsForbidden = true
			break
		}
	}

	return IsForbidden, nil
}
