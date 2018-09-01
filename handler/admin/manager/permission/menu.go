package permission

import (
	"apiserver/handler"
	"apiserver/pkg/token"
	"apiserver/service/admin/manager/permissionService"
	"github.com/gin-gonic/gin"
)

func Menu(c *gin.Context) {
	userInfo, err := token.ParseRequest(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
	}
	infos, err := permissionService.Menu(userInfo.ID)
	if err != nil {
		handler.SendResponse(c, err, nil)
	}
	handler.SendResponse(c, nil, infos)
}
