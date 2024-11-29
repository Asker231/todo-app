package db

import (
	"fmt"
	"github.com/Asker231/todo-app.git/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Db struct{
	*gorm.DB
}

func ConnectionDb(config *configs.ConfigApp)*Db{
	db, err := gorm.Open(postgres.Open(config.Db.DNS),&gorm.Config{})
	if err != nil{
		fmt.Println(err)
	}
	return &Db{
		DB: db,
	}
}