package global

import (
	"apiserver/handler"
	"apiserver/service/admin/manager/cacheService"
	"github.com/gin-gonic/gin"
)

func Cache(c *gin.Context) {

	if err := cacheService.Cache(); err != nil {
		handler.SendResponse(c, err, nil)
	}
	handler.SendResponse(c, nil, nil)
}
