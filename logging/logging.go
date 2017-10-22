// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	errgo "gopkg.in/errgo.v1"
)

var logger *zap.SugaredLogger

func LevelFromString(levelString string) (zapcore.Level, error) {
	var level zapcore.Level
	if err := level.UnmarshalText([]byte(levelString)); err != nil {
		return level, err
	}
	return level, nil
}

// Setup builds a sugared logger for use throughout the application.
func Setup(level zapcore.Level) error {
	cfg := zap.NewProductionConfig()
	cfg.Level.SetLevel(level)
	log, err := cfg.Build()
	if err != nil {
		return errgo.Mask(err)
	}
	logger = log.Sugar()
	return nil
}

// Logger retrieves the built logger
func Logger() *zap.SugaredLogger {
	if logger == nil {
		Setup(zapcore.InfoLevel)
	}
	return logger
}
