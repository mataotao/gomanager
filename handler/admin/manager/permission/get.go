package permission

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"strconv"
)

func Get(c *gin.Context) {
	//获取id并转化成int
	permissionId, _ := strconv.Atoi(c.Param("id"))
	//获取数据
	permission, err := managerModel.GetPermission(uint64(permissionId))

	if err != nil {
		log.Error("permission get", err)
		//错误返回
		handler.SendResponse(c, errno.Error, nil)
		return
	}
	//返回权限数据
	handler.SendResponse(c, nil, permission)

}
