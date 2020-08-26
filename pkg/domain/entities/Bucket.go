package entities

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/valueobjects"
)

type Bucket struct {
	ID valueobjects.BucketID

	Capacity int
}
