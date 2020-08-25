package usecases

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories"
	"net"
)

type AddIPToWhiteList struct {
	Subnet net.IPNet
}

type addIPToWhiteListHandler struct {
	ipRepository repositories.WhiteListIPRepository
}

func NewAddIPToWhiteListHandler(ipRepository repositories.WhiteListIPRepository) *addIPToWhiteListHandler {
	return &addIPToWhiteListHandler{ipRepository: ipRepository}
}

func (h *addIPToWhiteListHandler) Execute(useCase AddIPToWhiteList) error {

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
