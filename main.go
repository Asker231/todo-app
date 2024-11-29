package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

	type Todo struct{
		gorm.Model
		Title string `json:"title"`
	}

	func main(){
		err := godotenv.Load(".env")
		if err != nil{
			fmt.Println(err)
		}
		db,err := gorm.Open(postgres.Open(os.Getenv("DSN")),&gorm.Config{})
		if err != nil{
			fmt.Println(err)
		}
		var todo Todo
		db.AutoMigrate(&todo)
	}