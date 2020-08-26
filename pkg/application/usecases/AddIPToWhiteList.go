package usecases

import (
	"net"

	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories"
)

type AddIPToWhiteList struct {
	Subnet net.IPNet
}

type AddIPToWhiteListHandler struct {
	ipRepository repositories.WhiteListIPRepository
}

func NewAddIPToWhiteListHandler(ipRepository repositories.WhiteListIPRepository) *AddIPToWhiteListHandler {
	return &AddIPToWhiteListHandler{ipRepository: ipRepository}
}

func (h *AddIPToWhiteListHandler) Execute(useCase AddIPToWhiteList) error {
	ip, err := h.ipRepository.FindOneBySubnet(useCase.Subnet)
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
