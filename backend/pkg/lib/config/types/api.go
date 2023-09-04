package types

type Api struct {
	Rest Rest `mapstructure:"rest" json:"rest" yaml:"rest"`
}

type Rest struct {
	Port int `mapstructure:"port" json:"port" yaml:"port"`
}