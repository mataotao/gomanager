package permission

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"apiserver/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func Create(c *gin.Context) {
	//写入日志
	log.Info("调用创建用户", lager.Data{"X-Request-Id": util.GetReqID(c)})
	//声明 CreateRequest类型的变量
	var request CreateRequest

	//url获取并赋值
	if err := c.Bind(&request); err != nil {
		//返回错误
		handler.SendResponse(c, err, nil)
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
		handler.SendResponse(c, err, nil)
		return
	}

	//创建
	if err := p.Create(); err != nil {
		handler.SendResponse(c, err, nil)
	}
	log.Info("创建用户成功", lager.Data{"X-Request-Id": util.GetReqID(c)})
	//返回定义好的消息
	rsp := CreateResponse{request.Label}

	handler.SendResponse(c, nil, rsp)

}
