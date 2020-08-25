package factories

import (
	"errors"
	"github.com/Rototot/anti-brute-force/pkg/domain/constants"
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	errors2 "github.com/pkg/errors"
)

var ErrInvalidBucketType = errors.New("invalid bucket type")

type bucketFactory struct {
	conf BucketConfigurator
}

func NewBucketFactory(conf BucketConfigurator) *bucketFactory {
	return &bucketFactory{conf: conf}
}

func (f *bucketFactory) Create(bType constants.BucketType) (*entities.Bucket, error) {
	switch bType {
	case constants.BucketTypeIp:
		bucketCap := f.conf.IpBucketCapacity()
		if bucketCap == 0 {
			bucketCap = constants.DefaultIPBucketCapacity
		}

		return &entities.Bucket{
			Capacity: bucketCap,
		}, nil
	case constants.BucketTypeLogin:
		bucketCap := f.conf.LoginBucketCapacity()
		if bucketCap == 0 {
			bucketCap = constants.DefaultLoginBucketCapacity
		}

		return &entities.Bucket{
			Capacity: f.conf.LoginBucketCapacity(),
		}, nil
	case constants.BucketTypePassword:
		bucketCap := f.conf.PasswordBucketCapacity()
		if bucketCap == 0 {
			bucketCap = constants.DefaultPasswordBucketCapacity
		}

		return &entities.Bucket{
			Capacity: f.conf.IpBucketCapacity(),
		}, nil
	default:
		return nil, errors2.Wrapf(ErrInvalidBucketType, "value = %s", bType)
	}
}
