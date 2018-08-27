package user_bak

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service/admin"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// List list the users in the database.
func List(c *gin.Context) {
	log.Info("List function called.")
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := user.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}