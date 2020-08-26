//go:generate mockgen -source=$GOFILE -destination=./mocks/MockResetLoginAttemptsHandler.go -package=mocks

package usecases

import (
	"net"
	"sync"

	"github.com/Rototot/anti-brute-force/pkg/domain/valueobjects"

	"github.com/Rototot/anti-brute-force/pkg/domain/repositories"
)

type ResetLoginAttempts struct {
	Login string
	IP    net.IP
}

type ResetLoginAttemptsHandler struct {
	bucketRepository repositories.BucketRepository
}

func NewResetLoginAttemptsHandler(bucketRepository repositories.BucketRepository) *ResetLoginAttemptsHandler {
	return &ResetLoginAttemptsHandler{bucketRepository: bucketRepository}
}

func (h *ResetLoginAttemptsHandler) Execute(useCase ResetLoginAttempts) error {
	var wg sync.WaitGroup
	errs := make(chan error, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		errs <- h.bucketRepository.Remove(valueobjects.BucketID(useCase.IP))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		errs <- h.bucketRepository.Remove(valueobjects.BucketID(useCase.Login))
	}()

	wg.Wait()

	close(errs)
	for err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}
