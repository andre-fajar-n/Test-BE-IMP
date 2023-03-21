// Package to contains log
package log

import (
	"os"

	"test-be-IMP/internal"

	"github.com/rs/zerolog"
)

var log zerolog.Logger
var errorLevels = []zerolog.Level{
	zerolog.PanicLevel,
	zerolog.ErrorLevel,
	zerolog.FatalLevel,
	zerolog.WarnLevel,
}

func init() {
	log = zerolog.New(os.Stdout).
		With().
		Caller().
		Stack().
		Timestamp().
		Logger()

	if internal.Get().Debug {
		log.Hook(&hook{errorLevels})
	}
}

func Get() *zerolog.Logger {
	return &log
}

type hook struct {
	levels []zerolog.Level
}

func (h *hook) Run(e *zerolog.Event, level zerolog.Level, message string) {
	for _, l := range h.levels {
		if l == level {
			e.Msg(message)
		}
	}
}
