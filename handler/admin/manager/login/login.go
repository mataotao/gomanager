package login

import (
	"apiserver/handler"
	"apiserver/service/admin/manager/loginService"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var r LoginRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, err, nil)
	}
	token, err := loginService.Login(r.Username, r.Password, c.ClientIP())
	if err != nil {
		handler.SendResponse(c, err, nil)
	}
	handler.SendResponse(c, nil, LoginResponse{token})

}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
