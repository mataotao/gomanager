package middleware

import (
	"apiserver/pkg/errno"
	"apiserver/pkg/global/auth"
	"apiserver/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		info, err := token.ParseRequest(c);
		if err != nil {
			//handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.String(401, "key 无效")
			c.Error(errno.ErrTokenInvalid)
			c.Abort()
			return
		}
		authRoute := auth.Route(c.HandlerName(), info.ID)
		if authRoute == false {
			c.String(401, "权限不足")
			c.Error(errno.ErrTokenInvalid)
			c.Abort()
			return
		}
		return
		c.Next()
	}
}
