package factories

import (
	"errors"

	"github.com/Rototot/anti-brute-force/pkg/domain/constants"
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	errors2 "github.com/pkg/errors"
)

var ErrInvalidBucketType = errors.New("invalid bucket type")

type BucketFactory struct {
	conf BucketConfigurator
}

func NewBucketFactory(conf BucketConfigurator) *BucketFactory {
	return &BucketFactory{conf: conf}
}

func (f *BucketFactory) Create(bType constants.BucketType) (*entities.Bucket, error) {
	var bucketCap int

	switch bType {
	case constants.BucketTypeIp:
		bucketCap = f.conf.IPBucketCapacity()
	case constants.BucketTypeLogin:
		bucketCap = f.conf.LoginBucketCapacity()
	case constants.BucketTypePassword:
		bucketCap = f.conf.PasswordBucketCapacity()
	default:
		return nil, errors2.Wrapf(ErrInvalidBucketType, "value = %s", bType)
	}

	return &entities.Bucket{
		Capacity: bucketCap,
	}, nil
}
