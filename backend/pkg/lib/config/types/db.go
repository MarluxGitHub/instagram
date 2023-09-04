package types

type Db struct {
	Mysql Mysql
}

type Mysql struct {
	Host     string	`mapstructure:"host" json:"host" yaml:"host"`
	Port     int 	`mapstructure:"port" json:"port" yaml:"port"`
	Name	 string `mapstructure:"name" json:"name" yaml:"name"`
	User     string `mapstructure:"user" json:"user" yaml:"user"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}