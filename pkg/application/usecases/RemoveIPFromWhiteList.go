package usecases

import (
	"net"

	"github.com/Rototot/anti-brute-force/pkg/domain/repositories"
)

type RemoveIPFromWhiteList struct {
	Subnet net.IPNet
}

type RemoveIPFromWhiteListHandler struct {
	ipRepository repositories.WhiteListIPRepository
}

func NewRemoveIPFromWhiteListHandler(ipRepository repositories.WhiteListIPRepository) *RemoveIPFromWhiteListHandler {
	return &RemoveIPFromWhiteListHandler{ipRepository: ipRepository}
}

func (h *RemoveIPFromWhiteListHandler) Execute(useCase RemoveIPFromWhiteList) error {
	ip, err := h.ipRepository.FindOneBySubnet(useCase.Subnet)
	if err != nil {
		return err
	}

	if ip == nil {
		return nil
	}

	return h.ipRepository.Remove(ip)
}
