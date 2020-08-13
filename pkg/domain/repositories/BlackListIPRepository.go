package repositories

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"net"
)

type BlackListIPRepository interface {
	Add(ip *entities.BlackListIP) error
	Remove(ip *entities.BlackListIP) error
	FindBySubnet(subnet net.IPNet) (*entities.BlackListIP, error)
}
