package role

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"strconv"
)

func Delete(c *gin.Context) {
	id := c.Param("id")
	rid, _ := strconv.Atoi(id)
	var r managerModel.RoleModel
	r.BaseModel.Id = uint64(rid)
	if err := r.Delete(); err != nil {
		log.Error("role delete", err)
		handler.SendResponse(c, errno.Error, nil)
		return
	}
	log.Infof("删除角色id为%s", id)
	handler.SendResponse(c, nil, nil)
}
