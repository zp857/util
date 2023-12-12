package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(config Config) (err error) {
	if config.DBName == "" {
		err = fmt.Errorf("db-name 不能为空")
		return
	}
	switch config.DBType {
	case "mysql":
		db, err = GormMysql(config)
	case "pgsql":
		db, err = GormPgsql(config)
	default:
		err = fmt.Errorf("只支持 mysql 和 pgsql")
	}
	if err == nil {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(config.MaxIdleConns)
		sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	}
	return
}

func GormMysql(config Config) (db *gorm.DB, err error) {
	mysqlConfig := mysql.Config{
		DSN:                       config.DSN(),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}
	db, err = gorm.Open(mysql.New(mysqlConfig), config.GormConfig(config.Prefix, config.Singular))
	return
}

func GormPgsql(config Config) (db *gorm.DB, err error) {
	pgsqlConfig := postgres.Config{
		DSN:                  config.DSN(),
		PreferSimpleProtocol: false,
	}
	db, err = gorm.Open(postgres.New(pgsqlConfig), config.GormConfig(config.Prefix, config.Singular))
	return
}

func GetDB() *gorm.DB {
	return db
}
