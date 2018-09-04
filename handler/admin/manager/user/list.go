package user

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service/admin/manager/userService"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.Error, nil)
		return
	}
	cond := map[string]interface{}{"username": r.Username, "name": r.Name, "status": r.Status, "page": r.Page, "limit": r.Limit, "roleId": r.RoleId}
	res, err := userService.List(cond)
	if err != nil {
		log.Error("user list", err)
		handler.SendResponse(c, errno.Error, nil)
		return
	}
	handler.SendResponse(c, nil, res)
}
