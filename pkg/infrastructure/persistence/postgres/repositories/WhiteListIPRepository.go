package repositories

import (
	"database/sql"
	"net"

	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
)

type WhiteListIPRepository struct {
	conn *sql.DB
}

func NewWhiteListIPRepository(conn *sql.DB) *WhiteListIPRepository {
	return &WhiteListIPRepository{conn: conn}
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
