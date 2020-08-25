package configurators

import (
	"github.com/spf13/viper"
	"sync"
)

var testViperOne sync.Once
var testViper *viper.Viper

func NewViper() *viper.Viper {

	testViperOne.Do(func() {
		testViper = viper.New()
		viper.AutomaticEnv()
	})

	return testViper
}
