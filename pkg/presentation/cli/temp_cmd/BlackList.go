package temp_cmd

import (
	"github.com/Rototot/anti-brute-force/pkg/application/usecases"
	"net"
)

type blacklistCreator interface {
	Execute(useCase usecases.AddIpToBlacklist) error
}

type blacklistRemover interface {
	Execute(useCase usecases.RemoveIpFromBlackList) error
}

type BlackListCommand struct {
	creator blacklistCreator
	remover blacklistRemover
}

func (c *BlackListCommand) Add(rawNetwork string) error {

	_, network, err := net.ParseCIDR(rawNetwork)
	if err != nil {
		return err
	}

	return c.creator.Execute(usecases.AddIpToBlacklist{
		Subnet: *network,
	})
}

func (c *BlackListCommand) Remove(rawNetwork string) error {
	_, network, err := net.ParseCIDR(rawNetwork)
	if err != nil {
		return err
	}

	return c.remover.Execute(usecases.RemoveIpFromBlackList{
		Subnet: *network,
	})
}
