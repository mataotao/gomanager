package user

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"apiserver/pkg/errno"
	"apiserver/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func Current(c *gin.Context) {
	userInfo, err := token.ParseRequest(c)
	if err != nil {
		handler.SendResponse(c, errno.Error, nil)
		return
	}

	var u managerModel.UserModel
	u.Id = userInfo.ID
	info, err := u.Get()
	if err != nil {
		log.Error("user current", err)
		handler.SendResponse(c, errno.Error, nil)
		return
	}
	handler.SendResponse(c, nil, info)
}
