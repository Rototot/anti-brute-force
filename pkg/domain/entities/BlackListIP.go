package entities

import (
	"net"
	"time"
)

type BlackListIP struct {
	ID      int
	Subnet  net.IPNet
	Created time.Time
}
