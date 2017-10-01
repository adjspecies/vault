// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"go.uber.org/zap/zapcore"
	errgo "gopkg.in/errgo.v1"
	yaml "gopkg.in/yaml.v2"
)

// Config holds the values within a YAML configuration file
type Config struct {
	Host     string        `yaml:"host"`
	Port     int           `yaml:"port"`
	LogLevel zapcore.Level `yaml:"log-level"`
}

// Read loads a YAML config file into a Config object.
func Read(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errgo.Notef(err, "unable to open config file")
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errgo.Notef(err, "unable to read config file")
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, errgo.Notef(err, "unable to parse config file")
	}
	if err := validate(config); err != nil {
		return nil, errgo.Notef(err, "error validating config file")
	}
	return &config, nil
}

// validate checks that the configuration file provided is valid.
func validate(c Config) error {
	var missing []string
	if c.Host == "" {
		missing = append(missing, "host")
	}
	if c.Port == 0 {
		missing = append(missing, "port")
	}
	if len(missing) != 0 {
		return fmt.Errorf("missing fields %s", strings.Join(missing, ", "))
	}
	return nil
}
