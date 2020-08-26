package entities

import "github.com/Rototot/anti-brute-force/pkg/domain/valueobjects"

type Bucket struct {
	ID valueobjects.BucketID

	Capacity int

	//  The drips in the bucket
	Drips int
}

func (b *Bucket) IsFull() bool {
	return b.Drips >= b.Capacity
}

func (b *Bucket) IsEmpty() bool {
	return b.Drips <= 0
}

func (b *Bucket) AddDrips(qty int) {
	b.Drips += qty
}