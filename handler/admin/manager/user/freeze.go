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

func Freeze(c *gin.Context) {
	rId := c.Param("id")
	id, _ := strconv.Atoi(rId)
	var r FreezeRequest
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
	u.Status = r.Status
	if err := u.Update(); err != nil {
		log.Error("user freeze", err)
		handler.SendResponse(c, errno.Error, nil)
		return
	}
	log.Infof("用户冻结/解冻id为%s状态为%s", rId, u.Status)
	handler.SendResponse(c, nil, nil)
}
