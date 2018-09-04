package user

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"apiserver/pkg/errno"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"strconv"
)

func Update(c *gin.Context) {
	rId := c.Param("id")
	id, _ := strconv.Atoi(rId)

	var r UpdateRequest

	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.Error, nil)
		return
	}

	if _, err := govalidator.ValidateStruct(&r); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	u := managerModel.UserModel{
		Name:    r.Name,
		Mobile:  r.Mobile,
		HeadImg: r.HeadImg,
	}

	u.Id = uint64(id)
	if err := u.Updates(r.Roles); err != nil {
		log.Error("user update", err)
		handler.SendResponse(c, errno.Error, nil)
		return
	}
	ud, _ := json.Marshal(r)
	log.Infof("更新用户id为%s数据为%s", rId, ud)
	handler.SendResponse(c, nil, nil)

}
