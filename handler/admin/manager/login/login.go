package login

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service/admin/manager/loginService"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// @Summary Login generates the authentication token
// @Produce  json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ"}}"
// @Router /login [post]
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
