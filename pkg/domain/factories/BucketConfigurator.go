package factories

//go:generate mockgen --source=$GOFILE --destination=./mocks/MockBucketConfigurator.go -package=mocks

type BucketConfigurator interface {
	IPBucketCapacity() int
	LoginBucketCapacity() int
	PasswordBucketCapacity() int
}
