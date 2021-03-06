package role

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"strconv"
)

func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var rm managerModel.RoleModel
	rm.BaseModel.Id = uint64(id)
	info, err := rm.Get()
	if err != nil {
		log.Error("role get", err)
		handler.SendResponse(c, errno.Error, nil)
		return
	}
	handler.SendResponse(c, nil, info)

}
