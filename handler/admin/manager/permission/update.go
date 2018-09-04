package permission

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"apiserver/pkg/errno"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"strconv"
)

func Update(c *gin.Context) {

	//获取id 并转化类型
	permissionId, _ := strconv.Atoi(c.Param("id"))
	//定义要获取字段的类型
	var updateRequest UpdateRequest

	//绑定参数
	if err := c.Bind(&updateRequest); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	//验证参数是否合法
	if _, err := govalidator.ValidateStruct(&updateRequest); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}
	//定义类型
	var permissionModel managerModel.PermissionModel
	//类型转换并赋值
	permissionModel.Id = uint64(permissionId)
	permissionModel.Label = updateRequest.Label
	permissionModel.Sort = updateRequest.Sort
	permissionModel.IsContainMenu = updateRequest.IsContainMenu
	permissionModel.Url = updateRequest.Url
	permissionModel.Cond = updateRequest.Cond
	permissionModel.Icon = updateRequest.Icon
	//更新字段
	if err := permissionModel.Update(); err != nil {
		log.Error("permission update", err)
		handler.SendResponse(c, err, nil)
		return
	}
	ud, _ := json.Marshal(&updateRequest)
	log.Infof("更新权限成功,更新的数据为%s", ud)
	handler.SendResponse(c, nil, nil)
}
