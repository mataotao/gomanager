package user

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Pwd(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var r PwdRequest
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
	u.Password = r.Password
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	if err := u.Update(); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
	return
}
