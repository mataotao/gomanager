package user

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"github.com/gin-gonic/gin"
	"github.com/asaskevich/govalidator"
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
	if err := u.Freeze(&u); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
