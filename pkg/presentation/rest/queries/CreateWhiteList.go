package queries

import (
	"errors"
	"net"
)

var ErrInvalidSubnet = errors.New("invalid subnet")

type CreateWhiteListQuery struct {
	Subnet string `json:"ip" validate:"cidrv4"`
}

func (q *CreateWhiteListQuery) NormalizeSubnet() (*net.IPNet, error) {
	_, network, err := net.ParseCIDR(q.Subnet)
	if err != nil || network == nil {
		return nil, ErrInvalidSubnet
	}

	return network, nil
}
