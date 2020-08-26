package repositories

import (
	"net"

	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
)

//go:generate mockgen -source=$GOFILE -destination=./mocks/MockWhiteListIPRepository.go -package=mocks

type WhiteListIPRepository interface {
	Add(ip *entities.WhiteListIP) error
	Remove(ip *entities.WhiteListIP) error
	FindOneBySubnet(subnet net.IPNet) (*entities.WhiteListIP, error)
	FindAll() ([]*entities.WhiteListIP, error)
}
