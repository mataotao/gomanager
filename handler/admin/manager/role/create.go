package role

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"apiserver/pkg/errno"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
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
	role := managerModel.RoleModel{
		Name:        r.Name,
		Description: r.Description,
	}
	if err := role.Create(r.Permission); err != nil {
		log.Error("role create", err)
		handler.SendResponse(c, errno.Error, nil)
		return
	}
	//写入日志
	rd, _ := json.Marshal(&r)
	log.Infof("创建角色%s", rd)
	handler.SendResponse(c, nil, nil)
}

type CreateRequest struct {
	Name        string `json:"name" valid:"required"`
	Description string `json:"description"`
	Permission  []int  `json:"Permission" valid:"required"`
}
