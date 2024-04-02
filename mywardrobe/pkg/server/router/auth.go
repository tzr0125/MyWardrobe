package router

import (
	"example.com/m/v2/pkg/server/middleware"
	"github.com/gin-gonic/gin"
)

func Auth(r *gin.Engine) {

	authMiddleware := middleware.AuthMiddleware()

	r.POST("/login", authMiddleware.LoginHandler)
	r.GET("/logout", authMiddleware.LogoutHandler)

}	