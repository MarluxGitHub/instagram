package main

import (
	heartbeat "github.com/MarluxGitHub/instagram/pkg/domain/heartbeat/interfaces"
	"github.com/MarluxGitHub/instagram/pkg/lib/api/rest"
	"github.com/MarluxGitHub/instagram/pkg/lib/config"
	"github.com/MarluxGitHub/instagram/pkg/lib/db/mysql"
	"gorm.io/gorm"
)

var conf config.Configuration = config.Configuration{}
var database *gorm.DB

func main() {
	readConfig()
	connectToDatabase()
	initInterfaces()
}

func readConfig() {
	configReader := config.ConfigReaderImpl{}
	configReader.SetPath("config/development.config.json")

	c, err := configReader.ReadConfig()

	if err != nil {
		panic(err)
	}

	conf = c
}

func connectToDatabase() {
	db, err := mysql.ConnectToMysql(conf.DB.Mysql)

	if err != nil {
		panic(err)
	}

	database = db
}

func initInterfaces() {
	restApi := rest.NewRestApi()
	restApi.Init()

	heartbeat := heartbeat.NewHeartbeatRestEndpoint()
	restApi.AddEndpoint(&heartbeat)

	restApi.Start()
}
