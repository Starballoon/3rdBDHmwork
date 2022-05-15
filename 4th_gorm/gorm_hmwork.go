package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func main() {
	if err := DbInit(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println("Succeed")
}

var db *gorm.DB

func DbInit() error {
	var err error
	dsn := "root:LLQtT$3v@tcp(127.0.0.1:3306)/homework?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&User{})
	return err
}

type User struct {
	UUID    string `gorm:"uuid"`
	Name    string `gorm:"name"`
	Age     int    `gorm:"age"`
	Version int    `gorm:"version"`
}

func (User) TableName() string {
	return "user"
}
