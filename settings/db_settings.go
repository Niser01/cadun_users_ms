package settings

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed db_settings.yaml
var settingsFile []byte

//DatabaseConfig is the configuration for the database

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"pasword"`
	Name     string `yaml:"name"`
}

// Settings is the configuration for the application
type Settings struct {
	Port string         `yaml:"port"`
	DB   DatabaseConfig `yaml:"database"`
}

// New returns a new Settings object with the configuration loaded
func New() (*Settings, error) {
	var s Settings
	err := yaml.Unmarshal(settingsFile, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
