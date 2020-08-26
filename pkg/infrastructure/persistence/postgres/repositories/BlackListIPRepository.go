package repositories

import (
	"database/sql"
	"net"

	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
)

type BlackListIPRepository struct {
	conn *sql.DB
}

func NewBlackListIPRepository(conn *sql.DB) *BlackListIPRepository {
	return &BlackListIPRepository{conn: conn}
}

func (r *BlackListIPRepository) Add(ip *entities.BlackListIP) error {
	row := r.conn.QueryRow("INSERT INTO public.blacklists (subnet) VALUES (?) RETURNING id", ip.Subnet.String())

	return row.Scan(&ip.ID)
}

func (r *BlackListIPRepository) Remove(ip *entities.BlackListIP) error {
	_, err := r.conn.Exec("DELETE FROM blacklists WHERE id = ?", ip.ID)

	return err
}

func (r *BlackListIPRepository) FindBySubnet(subnet net.IPNet) (*entities.BlackListIP, error) {
	row := r.conn.QueryRow("SELECT * FROM blacklists WHERE subnet = ?", subnet.String())
	entity := &entities.BlackListIP{}

	err := row.Scan(&entity.ID, &entity.Subnet, &entity.Created)
	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		break
	default:
		return nil, err
	}

	return entity, nil
}

func (r *BlackListIPRepository) FindAll() ([]*entities.BlackListIP, error) {
	panic("implement me")
}
