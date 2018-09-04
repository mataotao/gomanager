package permission

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"apiserver/requests/admin/manager/permissionRequests"
	"apiserver/util"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"strconv"
)

func Update(c *gin.Context) {
	//日志
	log.Info("调用更新权限方法", lager.Data{"X-Request-Id": util.GetReqID(c)})
	//获取id 并转化类型
	permissionId, _ := strconv.Atoi(c.Param("id"))
	//定义要获取字段的类型
	var updateRequest permissionRequests.UpdateRequest

	//绑定参数
	if err := c.Bind(&updateRequest); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	//验证参数是否合法
	if _, err := govalidator.ValidateStruct(&updateRequest); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	//定义类型
	var permissionModel managerModel.PermissionModel
	//类型转换并赋值
	permissionModel.Id = uint64(permissionId)

	//更新字段
	if err := permissionModel.Update(&updateRequest); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	log.Info("更新权限成功", lager.Data{"X-Request-Id": util.GetReqID(c)})
	handler.SendResponse(c, nil, nil)
}
