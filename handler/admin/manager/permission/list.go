package permission

import (
	"apiserver/handler"
	"apiserver/pkg/token"
	"apiserver/service/admin/manager/permissionService"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func List(c *gin.Context) {
	userInfo, err := token.ParseRequest(c)
	if err != nil {
		log.Error("permission list",err)
		handler.SendResponse(c, errno.ErrAuthInvalid, nil)
		return
	}
	infos, err := permissionService.PermissionList(userInfo.ID)
	if err != nil {
		log.Error("permission list",err)
		handler.SendResponse(c, errno.Error, nil)
		return
	}
	handler.SendResponse(c, nil, infos)
}
