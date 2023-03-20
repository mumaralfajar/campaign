package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/campaign_startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	_ = db

	//userRepository := user.NewRepository(db)
	//user := user.User{
	//	Name: "John",
	//}
	//
	//_, err = userRepository.Save(user)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//
	//log.Println("User is created")
}
