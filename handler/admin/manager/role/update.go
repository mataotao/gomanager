package role

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"apiserver/pkg/errno"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"strconv"
)

func Update(c *gin.Context) {
	id := c.Param("id")
	rid, _ := strconv.Atoi(id)
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.Error, nil)
		return
	}
	var rm managerModel.RoleModel
	rm.BaseModel.Id = uint64(rid)
	data := &managerModel.RoleModel{Name: r.Name, Description: r.Description}
	if err := rm.Update(data, r.Permission); err != nil {
		log.Error("role update", err)
		handler.SendResponse(c, errno.Error, nil)
		return
	}
	rd, _ := json.Marshal(&r)
	log.Infof("更新角色 id为%s数据为%s", id, rd)
	handler.SendResponse(c, nil, nil)

}

type UpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Permission  []int  `json:"permission"`
}
