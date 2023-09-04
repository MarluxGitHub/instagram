package config

import "github.com/MarluxGitHub/instagram/pkg/lib/config/types"

type Configuration struct {
	Api types.Api `mapstructure:"api" json:"api" yaml:"api"`
	DB  types.Db `mapstructure:"db" json:"db" yaml:"db"`
}

