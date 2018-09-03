package role

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Update(c *gin.Context) {
	rid, _ := strconv.Atoi(c.Param("id"))
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	var rm managerModel.RoleModel
	rm.BaseModel.Id = uint64(rid)
	data := &managerModel.RoleModel{Name: r.Name, Description: r.Description}
	if err := rm.Update(data, r.Permission); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, nil)

}

type UpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Permission  []int  `json:"permission"`
}
