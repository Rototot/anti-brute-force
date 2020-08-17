package temp_cmd

import (
	"github.com/Rototot/anti-brute-force/pkg/application/usecases"
	"net"
)

type whitelistCreator interface {
	Execute(useCase usecases.AddIPToWhiteList) error
}

type whitelistRemover interface {
	Execute(useCase usecases.RemoveIpFromWhiteList) error
}

type WhiteListCommand struct {
	creator whitelistCreator
	remover whitelistRemover
}

func (c *WhiteListCommand) Add(rawNetwork string) error {

	_, network, err := net.ParseCIDR(rawNetwork)
	if err != nil {
		return err
	}

	return c.creator.Execute(usecases.AddIPToWhiteList{
		Subnet: *network,
	})
}

func (c *WhiteListCommand) Remove(rawNetwork string) error {
	_, network, err := net.ParseCIDR(rawNetwork)
	if err != nil {
		return err
	}

	return c.remover.Execute(usecases.RemoveIpFromWhiteList{
		Subnet: *network,
	})
}
