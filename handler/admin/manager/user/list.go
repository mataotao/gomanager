package user

import (
	"apiserver/handler"
	"apiserver/service/admin/manager/userService"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	cond := map[string]interface{}{"username": r.Username, "name": r.Name, "status": r.Status, "page": r.Page, "limit": r.Limit, "roleId": r.RoleId}
	res, err := userService.List(cond)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, res)
}
