package usecases

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories"
	"net"
)

type RemoveIpFromBlackList struct {
	Subnet net.IPNet
}

type removeIpFromBlackListHandler struct {
	ipRepository repositories.BlackListIPRepository
}

func NewRemoveIpFromBlackListHandler(ipRepository repositories.BlackListIPRepository) *removeIpFromBlackListHandler {
	return &removeIpFromBlackListHandler{ipRepository: ipRepository}
}

func (h *removeIpFromBlackListHandler) Execute(useCase RemoveIpFromBlackList) error {
	ip, err := h.ipRepository.FindBySubnet(useCase.Subnet)
	if err != nil {
		return err
	}

	if ip == nil {
		return nil
	}

	return h.ipRepository.Remove(ip)
}
