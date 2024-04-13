package main

import (
	"example.com/m/v2/pkg/server/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello World!",
        })
    })
    router.Test(r)
    r.Run(":8080")
}