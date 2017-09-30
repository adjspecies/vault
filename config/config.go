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
