// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package config_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap/zapcore"
	yaml "gopkg.in/yaml.v2"

	"github.com/adjspecies/vault/config"
)

var readTests = []struct {
	about          string
	content        []byte
	expectedConfig *config.Config
	expectedError  string
}{
	{
		about: "when the config is valid",
		content: mustMarshalYAML(map[string]interface{}{
			"host":        "localhost",
			"port":        5542,
			"log-level":   "info",
			"environment": "production",
		}),
		expectedConfig: &config.Config{
			Host:        "localhost",
			Port:        5542,
			LogLevel:    zapcore.InfoLevel,
			Environment: "production",
		},
	},
	{
		about: "and default environment to development",
		content: mustMarshalYAML(map[string]interface{}{
			"host": "localhost",
			"port": 5542,
		}),
		expectedConfig: &config.Config{
			Host:        "localhost",
			Port:        5542,
			LogLevel:    zapcore.Level(-10),
			Environment: "development",
		},
	},
	{
		about:         "config is unreadable",
		content:       []byte("bad wolf"),
		expectedError: "unable to parse config file",
	},
	{
		about: "config is invalid",
		content: mustMarshalYAML(map[string]interface{}{
			"bad": "wolf",
			"not": 42,
		}),
		expectedError: "missing fields host, port",
	},
	{
		about: "environment is invalid",
		content: mustMarshalYAML(map[string]interface{}{
			"host":        "localhost",
			"port":        5542,
			"log-level":   "info",
			"environment": "bad-wolf",
		}),
		expectedError: "environment must be `development` or `production`",
	},
}

func TestRead(t *testing.T) {
	Convey("It should be able to parse a config file", t, func() {
		for _, test := range readTests {
			Convey(fmt.Sprintf("...%s", test.about), func() {
				// Set up a temp file
				f, err := ioutil.TempFile("", "config")
				So(err, ShouldBeNil)
				defer f.Close()
				defer os.Remove(f.Name())
				_, err = f.Write(test.content)
				So(err, ShouldBeNil)

				conf, err := config.Read(f.Name())
				if test.expectedError != "" {
					So(err, ShouldBeError)
					return
				}
				So(err, ShouldBeNil)
				So(conf, ShouldResemble, test.expectedConfig)
			})
		}
	})
}

func mustMarshalYAML(v interface{}) []byte {
	out, err := yaml.Marshal(v)
	if err != nil {
		panic(err)
	}
	return out
}
