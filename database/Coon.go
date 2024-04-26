package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"remember/config"
	"remember/entity"
	"time"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(config.Configure.DbConfig.Database, config.Configure.GetDbUrl())
	if err != nil {
		log.Println(err.Error())
	}
	DB = db
	setPool(DB)
	// todo 初始化表
	DB.AutoMigrate(&entity.User{})          // 初始化用户表
	DB.AutoMigrate(&entity.Bill{})          // 初始化账单表
	DB.AutoMigrate(&entity.SystemMessage{}) // 初始化系统消息表
}
func GetCoon() *gorm.DB {
	return DB
}

func setPool(db *gorm.DB) {
	sqlDB := db.DB()
	sqlDB.SetMaxIdleConns(config.Configure.DbConfig.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(config.Configure.DbConfig.CoonMaxLifetime) * time.Hour)
	sqlDB.SetMaxOpenConns(config.Configure.DbConfig.MaxOpenConns)
}
