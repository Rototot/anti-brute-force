package usecases

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories"
	"net"
)

type AddIpToBlacklist struct {
	// todo перейти на ссылки
	Subnet net.IPNet
}

type addIpToBlacklistHandler struct {
	ipRepository repositories.BlackListIPRepository
}

func NewAddIpToBlacklistHandler(ipRepository repositories.BlackListIPRepository) *addIpToBlacklistHandler {
	return &addIpToBlacklistHandler{ipRepository: ipRepository}
}

func (h *addIpToBlacklistHandler) Execute(useCase AddIpToBlacklist) error {

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
