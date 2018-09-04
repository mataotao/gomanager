package role

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
<<<<<<< HEAD
	"github.com/gin-gonic/gin"
=======
	"apiserver/pkg/errno"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
>>>>>>> 937b3a9ca74cb2958e2ed35828a9b73ebf6808bf
	"strconv"
)

func Update(c *gin.Context) {
	id := c.Param("id")
	rid, _ := strconv.Atoi(id)
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
<<<<<<< HEAD
		handler.SendResponse(c, err, nil)
=======
		handler.SendResponse(c, errno.Error, nil)
>>>>>>> 937b3a9ca74cb2958e2ed35828a9b73ebf6808bf
		return
	}
	var rm managerModel.RoleModel
	rm.BaseModel.Id = uint64(rid)
	data := &managerModel.RoleModel{Name: r.Name, Description: r.Description}
	if err := rm.Update(data, r.Permission); err != nil {
<<<<<<< HEAD
		handler.SendResponse(c, err, nil)
		return
	}
=======
		log.Error("role update", err)
		handler.SendResponse(c, errno.Error, nil)
		return
	}
	rd, _ := json.Marshal(&r)
	log.Infof("更新角色 id为%s数据为%s", id, rd)
>>>>>>> 937b3a9ca74cb2958e2ed35828a9b73ebf6808bf
	handler.SendResponse(c, nil, nil)

}

type UpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Permission  []int  `json:"permission"`
}
