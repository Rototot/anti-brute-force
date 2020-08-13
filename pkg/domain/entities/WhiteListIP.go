package entities

import (
	"net"
	"time"
)

type WhiteListIP struct {
	ID      int
	Subnet  net.IPNet
	Created time.Time
}
