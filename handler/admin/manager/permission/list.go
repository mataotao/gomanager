package permission

import (
	"apiserver/handler"
	"apiserver/service/admin/manager/permissionService"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	infos, err := permissionService.PermissionList()
	if err != nil {
		handler.SendResponse(c, err, nil)
	}
	handler.SendResponse(c, nil, infos)
}
