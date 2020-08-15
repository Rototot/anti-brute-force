package services

type BucketConfigurator interface {
	IpBucketCapacity() int
	LoginBucketCapacity() int
	PasswordBucketCapacity() int
}
