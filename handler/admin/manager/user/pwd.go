package user

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"apiserver/pkg/errno"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"strconv"
)

func Pwd(c *gin.Context) {
	rId := c.Param("id")
	id, _ := strconv.Atoi(rId)
	var r PwdRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.Error, nil)
		return
	}

	if _, err := govalidator.ValidateStruct(&r); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	var u managerModel.UserModel
	u.Id = uint64(id)
	u.Password = r.Password
	if err := u.Encrypt(); err != nil {
		log.Error("user pwd", err)
		handler.SendResponse(c, errno.Error, nil)
		return
	}

	if err := u.Update(); err != nil {
		log.Error("user pwd", err)
		handler.SendResponse(c, errno.Error, nil)
		return
	}
	log.Infof("用户修改密码id为%s", rId)
	handler.SendResponse(c, nil, nil)
	return
}
