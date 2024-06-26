package loghandler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"strings"

	"github.com/SShlykov/lbalancer/lbcer/internal/pkg/logger/colors"
)

type PrettyHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *PrettyHandler) Handle(_ context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = colors.Colorize(level, colors.Magenta)
	case slog.LevelInfo:
		level = colors.Colorize(level, colors.Blue)
	case slog.LevelWarn:
		level = colors.Colorize(level, colors.Yellow)
	case slog.LevelError:
		level = colors.Colorize(level, colors.Red)
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	b, err := json.MarshalIndent(fields, "", "")
	if err != nil {
		return err
	}
	opts := strings.Replace(string(b), "\n", " ", -1)

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := colors.Colorize(r.Message, colors.Cyan)

	h.l.Println(timeStr, level, msg, colors.Colorize(opts, colors.White))

	return nil
}

func NewPrettyHandler(out io.Writer, opts HandlerOptions) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
	}

	return h
}
