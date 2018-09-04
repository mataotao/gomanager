package login

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service/admin/manager/loginService"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"

)

func Login(c *gin.Context) {
	var r LoginRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	login, err := loginService.Login(r.Username, r.Password, c.ClientIP())
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	log.Infof("%s登录成功", login.Username)
	handler.SendResponse(c, nil, login)

}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
