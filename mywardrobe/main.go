package main

import (
	"example.com/m/v2/pkg/server/router"
	db "example.com/m/v2/pkg/server/storage/db/postgresql"
	"github.com/gin-gonic/gin"
)

func main() {
    // 处理数据库
    db, err := db.InitPostgresql()
    if err != nil{
        return 
    }
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