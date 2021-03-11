package server

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config provide property of configuration file
type Config struct {
	DBUrl string `mapstructure:"DATABASE_URL"`
	Port  string `mapstructure:"PORT"`
}

// GetConfig from the file name and located in the path
// if name is empty then configuration file name become "config"
func GetConfig(name string, path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("GetConfig : %v", err)
	}

	config := new(Config)
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("GetConfig : %v", err)
	}

	return config, nil
}
