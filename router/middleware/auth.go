package middleware

import (
	"apiserver/pkg/errno"
	"apiserver/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			//handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.String(401, "key 无效")
			c.Error(errno.ErrTokenInvalid)
			c.Abort()
			return
		}
		return
		c.Next()
	}
}
