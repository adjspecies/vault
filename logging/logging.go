// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	errgo "gopkg.in/errgo.v1"
)

var logger *zap.SugaredLogger

// Setup builds a sugared logger for use throughout the application.
func Setup(environment string, level zapcore.Level) error {
	var cfg zap.Config
	if environment == "production" {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}
	if level != zapcore.Level(-10) {
		cfg.Level.SetLevel(level)
	}
	log, err := cfg.Build()
	if err != nil {
		return errgo.Mask(err)
	}
	logger = log.Sugar()
	return nil
}

func Logger() *zap.SugaredLogger {
	if logger == nil {
		Setup("development", zapcore.Level(-10))
	}
	return logger
}
