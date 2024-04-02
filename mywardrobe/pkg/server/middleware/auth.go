package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
  "os"
)


const (
	identity = "identity"
	jwtSecret = "key"
)

type login struct {
	Password string `form:"password" json:"password" binding:"required"`
  }

type User struct {
	UserName  string
  }


func AuthMiddleware() *jwt.GinJWTMiddleware {
  usernameEnv, passwordEnv := os.Getenv("USERNAME"), os.Getenv("PASSWORD")
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "gin-jwt",
    Key:         []byte(jwtSecret),
    Timeout:     time.Hour * 24,
    MaxRefresh:  time.Hour * 24,
    IdentityKey: identity,
    PayloadFunc: func(data interface{}) jwt.MapClaims {
      if v, ok := data.(*User); ok {
        return jwt.MapClaims{
			identity: v.UserName,
        }
      }
      return jwt.MapClaims{}
    },
    IdentityHandler: func(c *gin.Context) interface{} {
      claims := jwt.ExtractClaims(c)
      return &User{
        UserName: claims[identity].(string),
      }
    },
    Authenticator: func(c *gin.Context) (interface{}, error) {
      var loginVals login
      if err := c.ShouldBind(&loginVals); err != nil {
        return "", jwt.ErrMissingLoginValues
      }
      password := loginVals.Password

      if password == passwordEnv {
        return &User{UserName: usernameEnv}, nil
      }

      return nil, jwt.ErrFailedAuthentication
    },
    Authorizator: func(data interface{}, c *gin.Context) bool {
      if v, ok := data.(*User); ok && v.UserName == usernameEnv {
        return true
      }

      return false
    },
    Unauthorized: func(c *gin.Context, code int, message string) {
      c.JSON(code, gin.H{
        "code":    code,
        "message": message,
      })
    },
    // TokenLookup is a string in the form of "<source>:<name>" that is used
    // to extract token from the request.
    // Optional. Default value "header:Authorization".
    // Possible values:
    // - "header:<name>"
    // - "query:<name>"
    // - "cookie:<name>"
    // - "param:<name>"
    TokenLookup: "header: Authorization, query: token, cookie: jwt",
    // TokenLookup: "query:token",
    // TokenLookup: "cookie:token",
    TokenHeadName: "Bearer",
    TimeFunc: time.Now,
	})

	if err!= nil {
        panic(err)
    }
	return authMiddleware
}