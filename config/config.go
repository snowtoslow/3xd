package config

import (
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

//Config struct to store the future configs
type Config struct {
	AppName string `mapstructure:"app_name" json:"app_name" yaml:"app_name"`
	MadeBy  string `mapstructure:"made_by" json:"made_by" yaml:"made_by"`
}

func ParseConfigFile(path string) (Config, error) {
	viper.AddConfigPath(filepath.Dir(path))
	if fileExtension := filepath.Ext(path); fileExtension != "" {
		viper.SetConfigType(strings.Trim(fileExtension, "."))
	}
	_, nameWithExtension := filepath.Split(path)
	if name := strings.TrimSuffix(nameWithExtension, filepath.Ext(path)); name != "" {
		viper.SetConfigName(name)
	}

	var cfg Config
	if err := viper.ReadInConfig(); err != nil {
		return cfg, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
