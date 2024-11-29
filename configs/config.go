package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


	type ConfigApp struct{
		TOKEN string
	}


	func CompileConfig()*ConfigApp{
		err := godotenv.Load()
		if err != nil{
			fmt.Println(err)
		}
		return &ConfigApp{
			TOKEN: os.Getenv("TOKEN"),
		}
	}