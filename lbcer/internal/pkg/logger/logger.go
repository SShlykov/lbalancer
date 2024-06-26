package logger

import (
	"log/slog"

	"github.com/SShlykov/lbalancer/lbcer/internal/pkg/logger/loghandler"
)

type Logger interface {
	Warn(msg string, attrs ...any)
	Info(msg string, attrs ...any)
	Debug(msg string, attrs ...any)
	Error(msg string, attrs ...any)
}

type loggerImp struct {
	logger *slog.Logger
}

// Setup создает новый логгер с заданными параметрами
//
// level - нижний уровень логирования (debug, info, warn, error);
// mode  - режим логирования (pretty);
func Setup(level Level, mode Mode) (Logger, error) {
	opts := loghandler.HandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: level,
		},
	}

	handler, err := modeToHandler(mode, opts)
	if err != nil {
		return nil, err
	}

	return &loggerImp{logger: slog.New(handler)}, nil
}

func (l *loggerImp) Warn(msg string, attrs ...any) {
	l.logger.Warn(msg, attrs...)
}

func (l *loggerImp) Info(msg string, attrs ...any) {
	l.logger.Info(msg, attrs...)
}

func (l *loggerImp) Debug(msg string, attrs ...any) {
	l.logger.Debug(msg, attrs...)
}

func (l *loggerImp) Error(msg string, attrs ...any) {
	l.logger.Error(msg, attrs...)
}
