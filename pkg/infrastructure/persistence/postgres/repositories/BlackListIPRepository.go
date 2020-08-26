package repositories //nolint:dupl

import (
	"database/sql"
	"net"
	"time"

	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
)

type BlackListIPRepository struct {
	conn *sql.DB
}

func NewBlackListIPRepository(conn *sql.DB) *BlackListIPRepository {
	return &BlackListIPRepository{conn: conn}
}

func (r *BlackListIPRepository) Add(ip *entities.BlackListIP) error {
	row := r.conn.QueryRow("INSERT INTO blacklists (subnet) VALUES ($1) RETURNING id", ip.Subnet.String())

	return row.Scan(&ip.ID)
}

func (r *BlackListIPRepository) Remove(ip *entities.BlackListIP) error {
	_, err := r.conn.Exec("DELETE FROM blacklists WHERE subnet = ?", ip.ID)

	return err
}

func (r *BlackListIPRepository) FindBySubnet(subnet net.IPNet) (*entities.BlackListIP, error) {
	entity := &entities.BlackListIP{}
	err := r.conn.
		QueryRow("SELECT id, subnet, created_at FROM blacklists WHERE subnet = $1", subnet.String()).
		Scan(&entity.ID, &entity.Subnet, &entity.Created)

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
	var listEntities []*entities.BlackListIP

	rows, err := r.conn.Query("SELECT id, subnet, created_at FROM blacklists")

	if rows.Err() == sql.ErrNoRows || err == sql.ErrNoRows {
		return listEntities, nil
	}

	if rows.Err() != nil || err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		entity := &entities.BlackListIP{}

		var id int
		var rawSubnet string
		var created time.Time

		err := rows.Scan(&entity.ID, &rawSubnet, &entity.Created)
		if err != nil {
			return nil, err
		}

		_, network, err := net.ParseCIDR(rawSubnet)
		if err != nil {
			return listEntities, err
		}

		listEntities = append(listEntities, &entities.BlackListIP{
			ID:      id,
			Subnet:  *network,
			Created: created,
		})
	}

	return listEntities, nil
}
