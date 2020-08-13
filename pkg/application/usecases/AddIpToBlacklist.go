package usecases

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories"
	"net"
)

type AddIpToBlacklist struct {
	Subnet net.IPNet
}

type AddIpToBlacklistHandler struct {
	ipRepository repositories.BlackListIPRepository
}

func (h *AddIpToBlacklistHandler) Execute(useCase AddIpToBlacklist) error {

	ip, err := h.ipRepository.FindBySubnet(useCase.Subnet)
	if err != nil {
		return err
	}

	if ip == nil {
		ip = &entities.BlackListIP{
			Subnet: useCase.Subnet,
		}

		return h.ipRepository.Add(ip)
	}

	return nil
}
