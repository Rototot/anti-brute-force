package usecases

import (
	"net"

	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories"
)

type AddIPToBlacklist struct {
	// todo перейти на ссылки
	Subnet net.IPNet
}

type AddIPToBlacklistHandler struct {
	ipRepository repositories.BlackListIPRepository
}

func NewAddIPToBlacklistHandler(ipRepository repositories.BlackListIPRepository) *AddIPToBlacklistHandler {
	return &AddIPToBlacklistHandler{ipRepository: ipRepository}
}

func (h *AddIPToBlacklistHandler) Execute(useCase AddIPToBlacklist) error {
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
