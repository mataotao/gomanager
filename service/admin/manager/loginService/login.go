package loginService

import (
	"apiserver/model/admin/managerModel"
	"apiserver/pkg/auth"
	"apiserver/pkg/token"
	"fmt"
	"time"
)

func Login(username string, pwd string, ip string) (string, error) {
	u, err := managerModel.GetUser(username)
	if err != nil {
		return "", err
	}

	if err := auth.Compare(u.Password, pwd); err != nil {
		return "", err
	}
	t, err := token.Sign(token.Context{ID: u.Id, Username: u.Username}, "")
	if err != nil {
		return "", err
	}
	var userUpdate managerModel.UserModel
	userUpdate.Id = u.Id
	userUpdate.LastTime = time.Now()
	userUpdate.LastIp = ip
	if err := userUpdate.Update(); err != nil {
		return "", err
	}
	fmt.Println(err)
	return t, nil
}
