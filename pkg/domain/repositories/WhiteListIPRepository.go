package repositories

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"net"
)

type WhiteListIPRepository interface {
	Add(ip *entities.WhiteListIP) error
	Remove(ip *entities.WhiteListIP) error
	FindBySubnet(subnet net.IPNet) (*entities.WhiteListIP, error)
}
