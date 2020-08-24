package configurators

type BucketConfigurator struct {
	ipCapacity       int
	loginCapacity    int
	passwordCapacity int
}

func (f *BucketConfigurator) IpBucketCapacity() int {
	panic("implement me")
}

func (f *BucketConfigurator) LoginBucketCapacity() int {
	panic("implement me")
}

func (f *BucketConfigurator) PasswordBucketCapacity() int {
	panic("implement me")
}
