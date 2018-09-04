package user

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Freeze(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var r FreezeRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	if _, err := govalidator.ValidateStruct(&r); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	var u managerModel.UserModel
	u.Id = uint64(id)
	u.Status = r.Status
	if err := u.Update(); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
