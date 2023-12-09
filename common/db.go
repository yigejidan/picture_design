package common

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=Local",
		DbConfig.UserName,
		DbConfig.Password,
		DbConfig.Host,
		DbConfig.Port,
		DbConfig.Database,
		DbConfig.Charset,
		DbConfig.ParseTime,
	)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database,err:  %v", err.Error())
	}
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
