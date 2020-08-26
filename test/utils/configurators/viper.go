package configurators

import (
	"sync"

	"github.com/spf13/viper"
)

var (
	testViperOne sync.Once
	testViper    *viper.Viper
)

func NewViper() *viper.Viper {
	testViperOne.Do(func() {
		testViper = viper.New()
		viper.AutomaticEnv()
	})

	return testViper
}
