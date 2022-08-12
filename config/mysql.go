package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB(mysqlCfg *Mysql) *gorm.DB {
	dsn := fmt.Sprintf(
		"%v:%v@%v(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", mysqlCfg.Username, mysqlCfg.Password,
		mysqlCfg.Protocol, mysqlCfg.Host, mysqlCfg.Port, mysqlCfg.DbName,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 255,
	}), &gorm.Config{})

	if err != nil {
		panic("Failied to connect to db")
	} else {
		dbConfig, _ := db.DB()
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		dbConfig.SetMaxIdleConns(mysqlCfg.MaxIdleConn)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		dbConfig.SetMaxOpenConns(mysqlCfg.MaxOPenConn)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		dbConfig.SetConnMaxLifetime(mysqlCfg.ConnMaxLifeTime)

		DB = db
	}

	return DB
}

func GetDB() *gorm.DB {
	return DB
}
