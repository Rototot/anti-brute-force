package usecases

import (
	"net"

	"github.com/Rototot/anti-brute-force/pkg/domain/repositories"
)

type RemoveIPFromBlackList struct {
	Subnet net.IPNet
}

type RemoveIPFromBlackListHandler struct {
	ipRepository repositories.BlackListIPRepository
}

func NewRemoveIPFromBlackListHandler(ipRepository repositories.BlackListIPRepository) *RemoveIPFromBlackListHandler {
	return &RemoveIPFromBlackListHandler{ipRepository: ipRepository}
}

func (h *RemoveIPFromBlackListHandler) Execute(useCase RemoveIPFromBlackList) error {
	ip, err := h.ipRepository.FindBySubnet(useCase.Subnet)
	if err != nil {
		return err
	}

	if ip == nil {
		return nil
	}

	return h.ipRepository.Remove(ip)
}
