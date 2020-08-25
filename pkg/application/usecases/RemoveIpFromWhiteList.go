package usecases

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories"
	"net"
)

type RemoveIpFromWhiteList struct {
	Subnet net.IPNet
}

type removeIpFromWhiteListHandler struct {
	ipRepository repositories.WhiteListIPRepository
}

func NewRemoveIpFromWhiteListHandler(ipRepository repositories.WhiteListIPRepository) *removeIpFromWhiteListHandler {
	return &removeIpFromWhiteListHandler{ipRepository: ipRepository}
}

func (h *removeIpFromWhiteListHandler) Execute(useCase RemoveIpFromWhiteList) error {
	ip, err := h.ipRepository.FindOneBySubnet(useCase.Subnet)
	if err != nil {
		return err
	}

	if ip == nil {
		return nil
	}

	return h.ipRepository.Remove(ip)
}
