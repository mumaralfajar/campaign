package main

import (
	"bwastartup/user"
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

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "John"
	userInput.Occupation = "Programmer"
	userInput.Email = "john@mail.com"
	userInput.Password = "password"

	registerUser, err := userService.RegisterUser(userInput)
	if err != nil {
		return
	}

	log.Println(registerUser)
}
