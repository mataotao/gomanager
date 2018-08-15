package user

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"github.com/gin-gonic/gin"
	"github.com/asaskevich/govalidator"
	"time"
	"errors"
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
		LastTime: time.Now().Format("2006-01-02 15:04:05"),
		LastIp:   c.ClientIP(),
	}

	if isUnique := user.Uinque(); isUnique == false {
		handler.SendResponse(c, errors.New("重复"), nil)
		return
	}

	if err := user.Encrypt(); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	if err:=user.Create(r.Roles);err!=nil{
		handler.SendResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, nil)


}
