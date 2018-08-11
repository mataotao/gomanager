package role

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"apiserver/service/admin/manager/roleService"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var r listRequest

	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	list, count, err := roleService.List(r.Name, r.Page, r.Limit)

	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	res := listResponse{count, list}

	handler.SendResponse(c, err, res)
}

type listRequest struct {
	Page  uint64 `json:"page"`
	Name  string `json:"name"`
	Limit uint64 `json:"limit"`
}

type listResponse struct {
	Count uint64                       `json:"count"`
	List  []*managerModel.RoleListInfo `json:"list"`
}
