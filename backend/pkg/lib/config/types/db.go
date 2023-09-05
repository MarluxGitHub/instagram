package types

type Db struct {
	Mysql Mysql
}

type Mysql struct {
	Port     int
	Name	 string
	Host     string
	User     string
	Password string
}