package middleware

import (
	"apiserver/pkg/errno"
	"apiserver/pkg/global/auth"
	"apiserver/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		info, err := token.ParseRequest(c)
		if err != nil {
			//handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.String(401, "")
			c.Error(errno.ErrTokenInvalid)
			c.Abort()
			return
		}
		handlerName := c.HandlerName()
		authRoute := auth.Route(handlerName, info.ID)
		if authRoute == false {
			c.String(401, "")
			c.Error(errno.ErrAuthInvalid)
			c.Abort()
			return
		}
		log.Infof("用户%s调用%s", info.Username, handlerName)
		c.Next()
	}
}
