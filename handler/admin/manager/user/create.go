package user

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
<<<<<<< HEAD
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
=======
	"apiserver/pkg/errno"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
>>>>>>> 937b3a9ca74cb2958e2ed35828a9b73ebf6808bf
	"time"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.Error, nil)
		return
	}

	if _, err := govalidator.ValidateStruct(&r); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
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
		handler.SendResponse(c, errno.UserNameNotUnique, nil)
		return
	}

	if err := user.Encrypt(); err != nil {
		log.Error("user create", err)
		handler.SendResponse(c, errno.Error, nil)
		return
	}

	if err := user.Create(r.Roles); err != nil {
		log.Error("user create", err)
		handler.SendResponse(c, errno.Error, nil)
		return
	}
	ud, _ := json.Marshal(&r)
	log.Infof("创建用户成功 数据为%s", ud)
	handler.SendResponse(c, nil, nil)

}
