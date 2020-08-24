package repositories

import (
	"database/sql"
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"net"
)

type WhiteListIPRepository struct {
	conn *sql.Conn
}

func (r *WhiteListIPRepository) Add(ip *entities.WhiteListIP) error {
	panic("implement me")
}

func (r *WhiteListIPRepository) Remove(ip *entities.WhiteListIP) error {
	panic("implement me")
}

func (r *WhiteListIPRepository) FindOneBySubnet(subnet net.IPNet) (*entities.WhiteListIP, error) {
	panic("implement me")
}

func (r *WhiteListIPRepository) FindAll() ([]*entities.WhiteListIP, error) {
	panic("implement me")
}
