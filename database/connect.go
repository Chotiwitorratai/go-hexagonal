package database

import (
	"fmt"
	"go-hexagonal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() (*gorm.DB, error) {

	dsn := "root:@tcp(127.0.0.1:3306)/GoHexagonal?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		// return nil,err
	}
	fmt.Println("Connect Database !")
	db.AutoMigrate(&model.Customer{}, &model.Item{}, &model.Order{})
	fmt.Println("Database Migrate!")
	return db, nil
}
