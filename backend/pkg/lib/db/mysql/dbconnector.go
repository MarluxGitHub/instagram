package mysql

import (
	"strconv"

	"github.com/MarluxGitHub/instagram/pkg/lib/config/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToMysql(conf types.Mysql) (*gorm.DB, error) {

	strPort := strconv.Itoa(conf.Port)
	dsn := conf.User + ":" + conf.Password + "@tcp(" + conf.Host + ":" + strPort + ")/" + conf.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}