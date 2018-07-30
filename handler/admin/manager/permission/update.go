package permission

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"github.com/asaskevich/govalidator"
	"apiserver/model/admin/managerModel"
	"apiserver/requests/admin/manager/permissionRequests"
	"apiserver/util"
	"apiserver/handler"
	"strconv"
)

func Update(c *gin.Context) {
	log.Info("调用更新权限方法", lager.Data{"X-Request-Id": util.GetReqID(c)})
	permissionId, _ := strconv.Atoi(c.Param("id"))
	var updateRequest permissionRequests.UpdateRequest

	if err := c.Bind(&updateRequest); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	if _, err := govalidator.ValidateStruct(updateRequest); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	var permissionModel managerModel.PermissionModel

	permissionModel.Id = uint64(permissionId)

	if err := permissionModel.Update(&updateRequest); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	log.Info("更新权限成功", lager.Data{"X-Request-Id": util.GetReqID(c)})
	handler.SendResponse(c, nil, nil)
}
