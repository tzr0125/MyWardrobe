package router

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func Test(r *gin.Engine) {
	r.POST("/api/test/getauth", func(c *gin.Context){
		c.Header("set-cookie", "csrftoken=123456")
		c.JSON(200, gin.H{"key": "12345"})
	})
	r.POST("/api/test/testauth", func(c *gin.Context){
		s := c.Request.Header.Get("Cookie")
		if len(s) == 0 {
			c.JSON(400, gin.H{
                "message": "no cookie",
            })
            return
		}
		fmt.Println(s)
		s = c.Request.Header.Get("Authorization")
		if len(s) == 0 {
			c.JSON(400, gin.H{
                "message": "no authorization",
            })
            return
		}
		fmt.Println(s)
		c.JSON(200, gin.H{"message": "ok"})
	})
}	