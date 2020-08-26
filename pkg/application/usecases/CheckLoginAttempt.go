//go:generate mockgen -source=$GOFILE -destination=./mocks/MockCheckLoginAttemptHandler.go -package=mocks

package usecases

import (
	"net"

	"github.com/Rototot/anti-brute-force/pkg/domain/constants"
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories"
	"github.com/Rototot/anti-brute-force/pkg/domain/valueobjects"
)

type CheckLoginAttempt struct {
	Login    string
	Password string
	IP       net.IP
}

type bucketFactory interface {
	Create(bType constants.BucketType) (*entities.Bucket, error)
}

type ipGuard interface {
	HasAccess(ip net.IP) (bool, error)
}

type rateLimiter interface {
	IsLimitExceeded(bucket *entities.Bucket) (bool, error)
}

type CheckLoginAttemptHandler struct {
	bucketRepository repositories.BucketRepository
	bucketFactory    bucketFactory
	ipGuard          ipGuard
	rateLimiter      rateLimiter
}

func NewCheckLoginAttemptHandler(bucketRepository repositories.BucketRepository, bucketFactory bucketFactory, ipGuard ipGuard, rateLimiter rateLimiter) *CheckLoginAttemptHandler {
	return &CheckLoginAttemptHandler{bucketRepository: bucketRepository, bucketFactory: bucketFactory, ipGuard: ipGuard, rateLimiter: rateLimiter}
}

func (h *CheckLoginAttemptHandler) Execute(useCase CheckLoginAttempt) error {
	hasAccess, err := h.ipGuard.HasAccess(useCase.IP)
	if err != nil {
		return err
	}

	if !hasAccess {
		return constants.ErrAccessDenied
	}

	bucketInitializes := []func() (*entities.Bucket, error){
		func() (*entities.Bucket, error) {
			return h.findOrCreateBucket(valueobjects.BucketID(useCase.IP.String()), constants.BucketTypeIp)
		},
		func() (*entities.Bucket, error) {
			return h.findOrCreateBucket(valueobjects.BucketID(useCase.Login), constants.BucketTypeLogin)
		},
		func() (*entities.Bucket, error) {
			return h.findOrCreateBucket(valueobjects.BucketID(useCase.Password), constants.BucketTypePassword)
		},
	}

	for _, bucketInitializer := range bucketInitializes {
		bucket, err := bucketInitializer()
		if err != nil {
			return err
		}
		isExceeded, err := h.rateLimiter.IsLimitExceeded(bucket)
		if err != nil {
			return err
		}

		if isExceeded {
			return constants.ErrAttemptsIsExceeded
		}
	}

	return nil
}

func (h *CheckLoginAttemptHandler) findOrCreateBucket(id valueobjects.BucketID, bType constants.BucketType) (*entities.Bucket, error) {
	bucket, err := h.bucketRepository.FindOneByID(id)
	if err != nil {
		return nil, err
	}

	if bucket == nil {
		bucket, err = h.bucketFactory.Create(bType)
		if err != nil {
			return nil, err
		}

		bucket.ID = id
		err = h.bucketRepository.Add(bucket)
		if err != nil {
			return nil, err
		}
	}

	return bucket, err
}
