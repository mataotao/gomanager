package permission

import (
	"github.com/gin-gonic/gin"
	"apiserver/handler"
	"apiserver/service/admin/manager/permissionService"
)

func List(c *gin.Context) {
	infos, err := permissionService.PermissionList()
	if err != nil {
		handler.SendResponse(c, err, nil)
	}
	handler.SendResponse(c, nil, infos)
}
