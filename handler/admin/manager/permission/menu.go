package permission

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/pkg/token"
	"apiserver/service/admin/manager/permissionService"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func Menu(c *gin.Context) {
	userInfo, err := token.ParseRequest(c)
	if err != nil {
		log.Error("permission list", err)
		handler.SendResponse(c, errno.ErrAuthInvalid, nil)
		return
	}
	infos, err := permissionService.Menu(userInfo.ID)
	if err != nil {
		log.Error("permission list", err)
		handler.SendResponse(c, errno.Error, nil)
	}
	handler.SendResponse(c, nil, infos)
}
