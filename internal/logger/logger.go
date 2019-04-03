package logger

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Default(verbosity int) *zap.SugaredLogger {
	tmpfile, err := ioutil.TempFile("", "pgutil")
	if err != nil {
		fmt.Fprintln(os.Stderr, "unable to open trace log:", err)
		tmpfile = nil
	} else {
		fmt.Fprintln(os.Stderr, "logging to:", tmpfile.Name())
	}

	return New(verbosity, os.Stderr, tmpfile)
}

func Discard() *zap.SugaredLogger {
	return New(0, ioutil.Discard, nil)
}

func New(verbosity int, defaultLog io.Writer, traceLog io.Writer) *zap.SugaredLogger {
	var level zapcore.Level
	switch {
	case verbosity >= 3:
		level = zapcore.DebugLevel
	case verbosity == 2:
		level = zapcore.InfoLevel
	case verbosity == 1:
		level = zapcore.WarnLevel
	default:
		level = zapcore.ErrorLevel
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			MessageKey:     "msg",
			NameKey:        "logger",
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
		}),
		zapcore.AddSync(defaultLog),
		level,
	)
	if traceLog != nil {
		trace := zapcore.NewCore(
			zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
				MessageKey:     "msg",
				LevelKey:       "level",
				NameKey:        "logger",
				TimeKey:        "time",
				EncodeLevel:    zapcore.CapitalLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
			}),
			zapcore.AddSync(traceLog),
			zapcore.DebugLevel,
		)
		core = zapcore.NewTee(core, trace)
	}
	return zap.New(core).Sugar()
}
