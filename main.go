package main

import (
	"bwastartup/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	//dsn := "root:@tcp(127.0.0.1:3306)/campaign?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//
	//var users []user.User
	//
	//db.Find(&users)
	//
	//println("Total user:", len(users))
	//
	//for _, user := range users {
	//	println(user.Name, user.Email)
	//}

	router := gin.Default()
	router.GET("/handler", handler)
	//handle error for router.Run()
	err := router.Run()
	if err != nil {
		return
	}
}

func handler(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/campaign_startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK, users)
}
