package temp_cmd

import (
	"fmt"
	"github.com/Rototot/anti-brute-force/pkg/application/usecases"
	"net"
)

type bucketCleaner interface {
	Execute(useCase *usecases.ResetLoginAttempts) error
}

type BucketCommand struct {
	cleaner bucketCleaner
}

func (c *BucketCommand) Reset(login, rawIP string) error {
	ip := net.ParseIP(rawIP)
	if ip == nil {
		return fmt.Errorf("incorrect ip %s", rawIP)
	}

	return c.cleaner.Execute(&usecases.ResetLoginAttempts{
		Login: login,
		IP:    ip,
	})
}
