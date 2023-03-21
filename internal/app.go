// Package for internal app
package internal

import (
	"sync"

	"test-be-IMP/pkg/config"
)

type Config struct {
	Debug bool   `mapstructure:"debug"`
	Name  string `mapstructure:"name"`
	Port  int    `mapstructure:"port"`
}

var (
	cfg  *Config
	once sync.Once
)

// function init
func init() {
	once.Do(func() {
		var err error
		cfg, err = config.Read[Config]("config.yaml", "app")
		if nil != err {
			panic(err)
		}
	})
}

func Get() *Config {
	if nil != cfg {
		return cfg
	}

	return nil
}
