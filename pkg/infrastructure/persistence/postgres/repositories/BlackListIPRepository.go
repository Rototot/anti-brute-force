package repositories

import (
	"database/sql"
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"net"
)

type BlackListIPRepository struct {
	conn *sql.DB
}

func (r *BlackListIPRepository) Add(ip *entities.BlackListIP) error {
	rows, err := r.conn.Query("INSERT INTO public.blacklists (subnet) VALUES (?) RETURNING id", ip.Subnet.String())
	if err != nil {
		return err
	}
	defer rows.Close()

	return rows.Scan(&ip.ID)
}

func (r *BlackListIPRepository) Remove(ip *entities.BlackListIP) error {
	_, err := r.conn.Query("DELETE FROM blacklists WHERE id = ?", ip.ID)

	return err
}

func (r *BlackListIPRepository) FindBySubnet(subnet net.IPNet) (*entities.BlackListIP, error) {
	row, err := r.conn.Query("SELECT * FROM blacklists WHERE subnet = ?", subnet.String())
	if err != nil {
		return nil, err
	}

	entity := &entities.BlackListIP{}

	err = row.Scan(&entity.ID, &entity.Subnet, &entity.Created)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *BlackListIPRepository) FindAll() ([]*entities.BlackListIP, error) {
	panic("implement me")
}
