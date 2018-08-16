package user

import (
	"apiserver/handler"
	"github.com/gin-gonic/gin"
	"github.com/asaskevich/govalidator"
	"apiserver/model/admin/managerModel"
	"strconv"
)

func Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var r UpdateRequest

	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	if _,err := govalidator.ValidateStruct(&r); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	u := managerModel.UserModel{
		Name:    r.Name,
		Mobile:  r.Mobile,
		HeadImg: r.HeadImg,
	}

	u.Id = uint64(id)
	if err := u.Updates(r.Roles); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, nil)

}
