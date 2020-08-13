package usecases

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories"
	"net"
)

type RemoveIpFromWhiteList struct {
	Subnet net.IPNet
}

type RemoveIpFromWhiteListHandler struct {
	ipRepository repositories.WhiteListIPRepository
}

func (h *RemoveIpFromWhiteListHandler) Execute(useCase RemoveIpFromWhiteList) error {
	ip, err := h.ipRepository.FindBySubnet(useCase.Subnet)
	if err != nil {
		return err
	}

	if ip == nil {
		return nil
	}

	return h.ipRepository.Remove(ip)
}
