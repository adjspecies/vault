// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

// Package logging provides logging infrastructure to Vault.
//
// It is designed so that `Setup` is called once, and from then on, any package
// may call `logging.Logger()` to retrieve the constructed logger. On the off
// chance that `Setup` was not called, `Logger` will run it with sensible
// defaults.
//
// For testing, an observer is set up which provides access to logged messages.
// This is useful for commands, which often do work such as printing; insight
// into logs may stand in place of `fmt.Printf` statements.
package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	. "go.uber.org/zap/zaptest/observer"
	errgo "gopkg.in/errgo.v1"
)

var logger *zap.SugaredLogger
var observedLogger zapcore.Core
var logs *ObservedLogs

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

// Logger returns a sugared logger. If the logger hasn't been set up, it calls
// Setup.
func Logger() *zap.SugaredLogger {
	if logger == nil {
		Setup("development", zapcore.Level(-10))
	}
	return logger
}

// ObserverLogs provides the list of logs generated during the observation
// process.
func ObserverLogs() *ObservedLogs {
	return logs
}

// ObserveLogging constructs a logger through the zap/zaptest/observer framework
// so that logs may be accessible in tests.
func ObserveLogging(level zapcore.Level) {
	observedLogger, logs = New(level)
	logger = zap.New(observedLogger).With().Sugar()
}
