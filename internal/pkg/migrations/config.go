// Package for config and manipulation migrations
package migrations

import (
	"sync"

	"test-be-IMP/pkg/config"
)

type Config struct {
	Migration Migration `mapstructure:"migration"`
}

type Migration struct {
	Source string `mapstructure:"source"`
}

var cfg *Config
var once = new(sync.Once)

func init() {
	once.Do(func() {
		var err error
		cfg, err = config.Read[Config]("config.yaml", "migrations")
		if nil != err {
			panic(err)
		}
	})
}
