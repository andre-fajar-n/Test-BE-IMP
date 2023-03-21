package jwt

import (
	"sync"

	"test-be-IMP/pkg/config"
)

type JwtConfig struct {
	JwtSecret string
	JwtExp    int
}

var Cfg *JwtConfig
var once = new(sync.Once)

func init() {
	once.Do(func() {
		var err error
		Cfg, err = config.Read[JwtConfig]("config.yaml", "jwt")
		if nil != err {
			panic(err)
		}
	})
}
