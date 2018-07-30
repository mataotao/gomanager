package permission

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"apiserver/model/admin/managerModel"
	"apiserver/util"
	"apiserver/handler"
	"strconv"
)

func Update(c *gin.Context) {
	log.Info("调用更新权限方法", lager.Data{"X-Request-Id": util.GetReqID(c)})
	permissionId, _ := strconv.Atoi(c.Param("id"))
	var permissionModel managerModel.PermissionModel

	if err := c.Bind(&permissionModel); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	permissionModel.Id = uint64(permissionId)

	if err := permissionModel.Validate(); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	if err := permissionModel.Update(); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	log.Info("更新权限成功", lager.Data{"X-Request-Id": util.GetReqID(c)})
	handler.SendResponse(c, nil, nil)
}
