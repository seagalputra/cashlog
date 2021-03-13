package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// GetConfig from the file name and located in the path
// if name is empty then configuration file name become "config"
func ReadConfig(name string, path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("GetConfig : %v", err)
	}

	return nil
}

// Get configuration value with given string of key
func Get(key string) string {
	return viper.GetString(key)
}
