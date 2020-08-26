package repositories //nolint:dupl

import (
	"database/sql"
	"net"
	"time"

	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
)

type WhiteListIPRepository struct {
	conn *sql.DB
}

func NewWhiteListIPRepository(conn *sql.DB) *WhiteListIPRepository {
	return &WhiteListIPRepository{conn: conn}
}

func (r *WhiteListIPRepository) Add(ip *entities.WhiteListIP) error {
	row := r.conn.QueryRow("INSERT INTO whitelists (subnet) VALUES ($1) RETURNING id", ip.Subnet.String())

	return row.Scan(&ip.ID)
}

func (r *WhiteListIPRepository) Remove(ip *entities.WhiteListIP) error {
	_, err := r.conn.Exec("DELETE FROM whitelists WHERE subnet = ?", ip.ID)

	return err
}

func (r *WhiteListIPRepository) FindOneBySubnet(subnet net.IPNet) (*entities.WhiteListIP, error) {
	entity := &entities.WhiteListIP{}
	err := r.conn.
		QueryRow("SELECT id, subnet, created_at FROM whitelists WHERE subnet = $1", subnet.String()).
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

func (r *WhiteListIPRepository) FindAll() ([]*entities.WhiteListIP, error) {
	var listEntities []*entities.WhiteListIP

	rows, err := r.conn.Query("SELECT id, subnet, created_at FROM whitelists")

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

		listEntities = append(listEntities, &entities.WhiteListIP{
			ID:      id,
			Subnet:  *network,
			Created: created,
		})
	}

	return listEntities, nil
}
