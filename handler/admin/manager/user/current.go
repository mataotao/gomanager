package user

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"apiserver/pkg/token"
	"github.com/gin-gonic/gin"
)

func Current(c *gin.Context) {
	userInfo, err := token.ParseRequest(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
	}

	var u managerModel.UserModel
	u.Id = userInfo.ID
	info, err := u.Get()
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, info)
}
