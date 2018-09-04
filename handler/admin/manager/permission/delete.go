package permission

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"strconv"
)

func Delete(c *gin.Context) {

	//接受参数 字符串转成数字
	id := c.Param("id")
	permissionId, _ := strconv.Atoi(id)
	//调用删除 uint64转化对应的类型
	if err := managerModel.DeletePermission(uint64(permissionId)); err != nil {
		log.Error("permission delete", err)
		//返回错误
		handler.SendResponse(c, errno.Error, nil)
		return
	}
	//日志
	log.Infof("删除权限成功 id为%s", id)
	//返回成功
	handler.SendResponse(c, nil, nil)
}
