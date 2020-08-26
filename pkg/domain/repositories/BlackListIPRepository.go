package repositories

import (
	"net"

	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
)

//go:generate mockgen -source=$GOFILE -destination=./mocks/MockBlackListIPRepository.go -package=mocks

type BlackListIPRepository interface {
	Add(ip *entities.BlackListIP) error
	Remove(ip *entities.BlackListIP) error
	FindBySubnet(subnet net.IPNet) (*entities.BlackListIP, error)
	FindAll() ([]*entities.BlackListIP, error)
}
