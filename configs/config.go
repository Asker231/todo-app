package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)
type ConfigApp struct{
	ConfAuth AuthConfig
}


type AuthConfig struct{
	TOKEN string
}

	func CompileConfig()*ConfigApp{
		err := godotenv.Load("../.env")
		if err != nil{
			fmt.Println(err)
		}
		return &ConfigApp{
			ConfAuth: AuthConfig{
				TOKEN: os.Getenv("TOKEN"),
			},
		}
	}