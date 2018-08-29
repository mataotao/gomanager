package loginService

import (
	"apiserver/model/admin/managerModel"
	"apiserver/pkg/auth"
	"apiserver/pkg/global/redis"
	"apiserver/pkg/token"
	"bytes"
	"strconv"
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
	currentTime := time.Now()
	t, err := token.Sign(token.Context{ID: u.Id, Username: u.Username}, "")
	if err != nil {
		return "", err
	}
	var userUpdate managerModel.UserModel
	userUpdate.Id = u.Id
	userUpdate.LastTime = currentTime
	userUpdate.LastIp = ip
	if err := userUpdate.Update(); err != nil {
		return "", err
	}
	pool := redis.Pool.Pool.Get()
	defer pool.Close()
	var key bytes.Buffer
	key.WriteString("user:login:")
	key.WriteString(strconv.Itoa(int(u.Id)))
	if _, err := pool.Do("Set", key.String(), t); err != nil {
		return "", err
	}

	return t, nil
}
