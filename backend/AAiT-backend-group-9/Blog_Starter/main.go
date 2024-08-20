package main

import (
    "Blog_Starter/api/router"
    "Blog_Starter/config"
    "time"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    env, _ := config.Load()
    db, _ := config.GetClient(env.DatabaseURL, env.DatabaseName)
    timeout, _ := time.ParseDuration(env.TimeOut)
    router.Setup(env, timeout, db, r)
    r.Run("localhost:" + env.Port)
}
