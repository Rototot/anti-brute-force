package configurators

import (
	"os"
	"strconv"
)

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

func NewBucketConfigurator() BucketConfigurator {
	var err error
	var conf BucketConfigurator

	if os.Getenv("APP_BUCKET_IP_CAP") != "" {
		conf.ipCapacity, err = strconv.Atoi(os.Getenv("APP_BUCKET_IP_CAP"))
		if err != nil {
			panic(err)
		}
	}

	if os.Getenv("APP_BUCKET_LOGIN_CAP") != "" {
		conf.loginCapacity, err = strconv.Atoi(os.Getenv("APP_BUCKET_LOGIN_CAP"))
		if err != nil {
			panic(err)
		}
	}

	if os.Getenv("APP_BUCKET_PASSWORD_CAP") != "" {
		conf.passwordCapacity, err = strconv.Atoi(os.Getenv("APP_BUCKET_PASSWORD_CAP"))
		if err != nil {
			panic(err)
		}
	}

	if os.Getenv("APP_BUCKET_IP_CAP") != "" {
		conf.ipCapacity, err = strconv.Atoi(os.Getenv("APP_BUCKET_IP_CAP"))
		if err != nil {
			panic(err)
		}
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

func (f BucketConfigurator) IPBucketCapacity() int {
	return f.ipCapacity
}

func (f BucketConfigurator) LoginBucketCapacity() int {
	return f.loginCapacity
}

func (f BucketConfigurator) PasswordBucketCapacity() int {
	return f.passwordCapacity
}
