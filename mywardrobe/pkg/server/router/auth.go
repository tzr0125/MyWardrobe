package router

import (
	"example.com/m/v2/pkg/server/middleware"
	"github.com/gin-gonic/gin"
)

func Auth(r *gin.Engine) gin.HandlerFunc {

	authMiddleware := middleware.AuthMiddleware()
	r.Handlers()
}	