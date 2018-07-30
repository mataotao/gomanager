package permission

import (
	"strconv"
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"apiserver/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func Delete(c *gin.Context) {

	//接受参数 字符串转成数字
	permissionId, _ := strconv.Atoi(c.Param("id"))
	//调用删除 uint64转化对应的类型
	if err := managerModel.DeletePermission(uint64(permissionId)); err != nil {
		//返回错误
		handler.SendResponse(c, err, nil)
		//日志
		log.Info("删除权限失败", lager.Data{"X-Request-Id": util.GetReqID(c)})
		return
	}
	//日志
	log.Info("删除权限成功", lager.Data{"X-Request-Id": util.GetReqID(c)})
	//返回成功
	handler.SendResponse(c, nil, nil)
}
