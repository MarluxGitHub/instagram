package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ConfigReader interface {
	ReadConfig() error
	SetPath(path string)
}

type ConfigReaderImpl struct {
	Path string
}

func (c *ConfigReaderImpl) SetPath(path string) {
	c.Path = path
}

// Read ViperConfig
func (c *ConfigReaderImpl) ReadConfig() (Configuration, error) {
	configuration := Configuration{}

	v := viper.New()
	v.SetConfigType("json")
	v.SetConfigFile(c.Path)

	if err := v.ReadInConfig(); err != nil {
        fmt.Printf("Error reading config file, %s", err)
    }

	err := viper.Unmarshal(&configuration)
    if err != nil {
        fmt.Printf("Unable to decode into Configuration, %v", err)
	}

	return configuration, nil
}