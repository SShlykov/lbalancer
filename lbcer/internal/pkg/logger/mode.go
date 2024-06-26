package logger

import (
	"log/slog"
	"os"

	"github.com/SShlykov/lbalancer/lbcer/internal/pkg/logger/loghandler"
)

type Mode int

const (
	ModePretty Mode = iota
)

func modeToHandler(mode Mode, opts loghandler.HandlerOptions) (slog.Handler, error) {
	switch mode {
	case ModePretty:
		return loghandler.NewPrettyHandler(os.Stdout, opts), nil
	default:
		return nil, ErrUnknownMode
	}
}

func ModeFromString(s string) (Mode, error) {
	switch s {
	case "pretty":
		return ModePretty, nil
	default:
		return 0, ErrUnknownMode
	}
}
