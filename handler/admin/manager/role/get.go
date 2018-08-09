package role

import (
	"github.com/gin-gonic/gin"
	"apiserver/handler"
	"strconv"
	"apiserver/model/admin/managerModel"
)

func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var rm managerModel.RoleModel
	rm.BaseModel.Id = uint64(id)
	info, err := rm.Get()
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, info)

}
