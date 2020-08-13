package usecases

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories"
	"net"
)

type AddIPToWhiteList struct {
	Subnet net.IPNet
}

type AddIPToWhiteListHandler struct {
	ipRepository repositories.WhiteListIPRepository
}

func (h *AddIPToWhiteListHandler) Execute(useCase AddIPToWhiteList) error {

	ip, err := h.ipRepository.FindBySubnet(useCase.Subnet)
	if err != nil {
		return err
	}

	if ip == nil {
		ip = &entities.WhiteListIP{
			Subnet: useCase.Subnet,
		}

		return h.ipRepository.Add(ip)
	}

	return nil
}
