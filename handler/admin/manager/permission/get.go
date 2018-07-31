package permission

import (
	"github.com/gin-gonic/gin"
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"strconv"
)

func Get(c *gin.Context) {
	//获取id并转化成int
	permissionId, _ := strconv.Atoi(c.Param("id"))
	//获取数据
	permission, err := managerModel.GetPermission(uint64(permissionId))

	if err != nil {
		//错误返回
		handler.SendResponse(c, err, nil)
		return
	}
	//返回权限数据
	handler.SendResponse(c, nil, permission)

}
