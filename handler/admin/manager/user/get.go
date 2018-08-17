package user

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var u managerModel.UserModel
	u.Id = uint64(id)
	info, err := u.Get()
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, info)
}
