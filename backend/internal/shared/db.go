package shared

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "tommy:12522363h!A@tcp(127.0.0.1:3306)/pasunes_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Faild to connect to the Database!")
	}
	fmt.Println("Connected to the Database")
	
	
	DB.AutoMigrate(&User{})

}

type User struct {
	Id       uint   `gorm:"primeryKey"`
	Email    string `gorm:"unique"`
	Password string
}
