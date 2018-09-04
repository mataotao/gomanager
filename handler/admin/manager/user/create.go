package user

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"time"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	if _, err := govalidator.ValidateStruct(&r); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	user := managerModel.UserModel{
		Username: r.Username,
		Name:     r.Name,
		Mobile:   r.Mobile,
		Password: r.Password,
		HeadImg:  r.HeadImg,
		LastTime: time.Now(),
		LastIp:   c.ClientIP(),
		Status:   managerModel.ON,
	}

	if isUnique := user.Uinque(); isUnique == false {
		handler.SendResponse(c, errors.New("重复"), nil)
		return
	}

	if err := user.Encrypt(); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	if err := user.Create(r.Roles); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, nil)

}
