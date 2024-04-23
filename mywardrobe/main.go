package main

import (
	"example.com/m/v2/pkg/server/router"
	"github.com/gin-gonic/gin"
    "example.com/m/v2/pkg/server/storage/db"
)

func main() {

    // 启动数据库
    db := db.InitMongoDB()
    defer db.Close()

    // 启动路由
	r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello World!",
        })
    })
    router.Test(r)
    r.Run(":8080")
}