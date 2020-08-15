package factories

//go:generate mockgen --source=$GOFILE --destination=./mocks/MockBucketConfigurator.go -package=mocks

type BucketConfigurator interface {
	IpBucketCapacity() int
	LoginBucketCapacity() int
	PasswordBucketCapacity() int
}
