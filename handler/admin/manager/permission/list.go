package permission

import (
	"github.com/gin-gonic/gin"
	"apiserver/requests/admin/manager/permissionRequests"
	"apiserver/handler"
	"apiserver/service/admin/manager/permissionService"
)

func List(c *gin.Context) {
	var p permissionRequests.ListRequest
	//绑定数据
	if err := c.Bind(&p); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	infos, total, err := permissionService.PermissionList(p.Limit, p.Page)
	if err != nil {
		handler.SendResponse(c, err, nil)
	}
	handler.SendResponse(c, nil, permissionService.ListResponse{
		Total:           total,
		PermissionsList: infos,
	})
}
