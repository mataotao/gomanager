package permission

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"apiserver/pkg/errno"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func Create(c *gin.Context) {
	//声明 CreateRequest类型的变量
	var request CreateRequest

	//url获取并赋值
	if err := c.Bind(&request); err != nil {
		//返回错误
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	//赋值
	p := managerModel.PermissionModel{
		Label:         request.Label,
		IsContainMenu: request.IsContainMenu,
		Pid:           request.Pid,
		Level:         request.Level,
		Url:           request.Url,
		Sort:          request.Sort,
		Cond:          request.Cond,
		Icon:          request.Icon,
	}
	//验证
	if err := p.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	//创建
	if err := p.Create(); err != nil {
		log.Error("permission create", err)
		handler.SendResponse(c, errno.Error, nil)
		return
	}
	cd, _ := json.Marshal(&p)
	log.Infof("创建用户成功 数据为%s", cd)
	//返回定义好的消息
	rsp := CreateResponse{request.Label}

	handler.SendResponse(c, nil, rsp)

}
