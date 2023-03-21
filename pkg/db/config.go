// Package to config database
package db

import (
	"fmt"
	"sync"

	"test-be-IMP/pkg/config"
)

type Config struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	Db       string
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Db,
	)
}

var cfg *Config
var once = new(sync.Once)

func init() {
	once.Do(func() {
		var err error
		cfg, err = config.Read[Config]("config.yaml", "db")
		if nil != err {
			panic(err)
		}
	})
}

func Get() *Config {
	return cfg
}
