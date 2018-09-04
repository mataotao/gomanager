package user

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
<<<<<<< HEAD
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
=======
	"apiserver/pkg/errno"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
>>>>>>> 937b3a9ca74cb2958e2ed35828a9b73ebf6808bf
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
