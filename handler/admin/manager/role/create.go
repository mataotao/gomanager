package role

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	//"apiserver/util"
	"github.com/gin-gonic/gin"
	"github.com/asaskevich/govalidator"
	//"github.com/lexkong/log"
	//"github.com/lexkong/log/lager"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	if _, err := govalidator.ValidateStruct(r); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	role := managerModel.RoleModel{
		Name:        r.Name,
		Description: r.Description,
	}
	if err := role.Create(r.Permission); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}

type CreateRequest struct {
	Name        string `json:"name" valid:"required"`
	Description string `json:"description"`
	Permission  []int  `json:"Permission" valid:"required"`
}
