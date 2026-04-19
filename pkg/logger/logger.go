package logger

import (
	"os"

	"github.com/rs/zerolog"
)

const fromKey = "_from"

type Logger interface {
	Panic() *zerolog.Event
	Fatal() *zerolog.Event
	Error() *zerolog.Event
	Warn() *zerolog.Event
	Info() *zerolog.Event
	Debug() *zerolog.Event
}

func NewLogger(caller string, options ...func(*zerolog.ConsoleWriter)) Logger {
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    true,
		TimeFormat: "2006-01-02 15:04:05",
	}

	for _, option := range options {
		option(&consoleWriter)
	}

	subLog := zerolog.New(consoleWriter).With().
		Timestamp().
		Str(fromKey, caller).
		Logger()

	return &logger{
		log: &subLog,
	}
}

func WithColorEnabled(enabled bool) func(*zerolog.ConsoleWriter) {
	return func(w *zerolog.ConsoleWriter) {
		w.NoColor = !enabled
	}
}

type logger struct {
	log *zerolog.Logger
}

func (l *logger) Panic() *zerolog.Event {
	return l.log.Panic()
}

func (l *logger) Fatal() *zerolog.Event {
	return l.log.Fatal()
}

func (l *logger) Error() *zerolog.Event {
	return l.log.Error()
}

func (l *logger) Warn() *zerolog.Event {
	return l.log.Warn()
}

func (l *logger) Info() *zerolog.Event {
	return l.log.Info()
}

func (l *logger) Debug() *zerolog.Event {
	return l.log.Debug()
}
