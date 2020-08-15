//go:generate mockgen -source=$GOFILE -destination=./mocks/MockResetLoginAttemptsHandler.go -package=mocks

package usecases

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories"
	"github.com/Rototot/anti-brute-force/pkg/domain/valueObjects"
	"net"
)

type ResetLoginAttempts struct {
	Login string
	IP    net.IP
}

type bucketCleaner interface {
	Clean(bucket *entities.Bucket) error
}

type ResetLoginAttemptsHandler struct {
	bucketRepository repositories.BucketRepository
	bucketCleaner    bucketCleaner
}

func (h *ResetLoginAttemptsHandler) Execute(useCase *ResetLoginAttempts) error {

	bucketSearchers := []func() (*entities.Bucket, error){
		func() (*entities.Bucket, error) {
			return h.bucketRepository.FindOneByID(valueObjects.BucketID(useCase.IP.String()))
		},
		func() (*entities.Bucket, error) {
			return h.bucketRepository.FindOneByID(valueObjects.BucketID(useCase.Login))
		},
	}

	for _, searcher := range bucketSearchers {
		bucket, err := searcher()
		if err != nil {
			return err
		}

		err = h.bucketCleaner.Clean(bucket)
		if err != nil {
			return err
		}
	}

	return nil
}
