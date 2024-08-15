package main

import (
	"group3-blogApi/config"
	"group3-blogApi/config/db"
	"group3-blogApi/delivery/routers"
)

func main() {
    config.InitiEnvConfigs() 
    db.ConnectDB(config.EnvConfigs.MongoURI)
      
    router := routers.SetupRouter()

    router.Run(config.EnvConfigs.LocalServerPort)
}
