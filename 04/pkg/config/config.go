package config

import (
	"github.com/BurntSushi/toml"
)

// Config represents mix of settings for the app.
type Config struct {
	Store Store `toml:"store"`
	Cache Cache `toml:"cache"`
}

// New creates reads application configuration from the file.
func New(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
