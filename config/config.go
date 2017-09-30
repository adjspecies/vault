// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package config

type Config struct {
    //
}

func Read(path string) (*Config, error) {
    var config = Config{}
    return &config, nil
}

func validate(c Config) error {
    return nil
}
