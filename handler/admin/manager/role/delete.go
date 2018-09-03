package role

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Delete(c *gin.Context) {
	rid, _ := strconv.Atoi(c.Param("id"))
	var r managerModel.RoleModel
	r.BaseModel.Id = uint64(rid)
	if err := r.Delete(); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
