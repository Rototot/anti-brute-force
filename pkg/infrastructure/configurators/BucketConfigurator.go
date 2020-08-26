package configurators

import "github.com/spf13/viper"

const (
	defaultBucketIPCap       = 1000
	defaultBucketPasswordCap = 100
	defaultBucketLoginCap    = 10
)

type BucketConfigurator struct {
	ipCapacity       int
	loginCapacity    int
	passwordCapacity int
}

func NewBucketConfigurator(v *viper.Viper) *BucketConfigurator {
	conf := &BucketConfigurator{
		ipCapacity:       v.GetInt("APP_BUCKET_IP_CAP"),
		loginCapacity:    v.GetInt("APP_BUCKET_LOGIN_CAP"),
		passwordCapacity: v.GetInt("APP_BUCKET_PASSWORD_CAP"),
	}

	if conf.ipCapacity == 0 {
		conf.ipCapacity = defaultBucketIPCap
	}

	if conf.loginCapacity == 0 {
		conf.loginCapacity = defaultBucketLoginCap
	}

	if conf.passwordCapacity == 0 {
		conf.passwordCapacity = defaultBucketPasswordCap
	}

	return conf
}

func (f *BucketConfigurator) IPBucketCapacity() int {
	return f.ipCapacity
}

func (f *BucketConfigurator) LoginBucketCapacity() int {
	return f.loginCapacity
}

func (f *BucketConfigurator) PasswordBucketCapacity() int {
	return f.passwordCapacity
}
