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
        return configuration, err
    }

    if err := v.Unmarshal(&configuration); err != nil {
        fmt.Printf("Error unmarshalling config file, %s", err)
        return configuration, err
    }

    return configuration, nil
}