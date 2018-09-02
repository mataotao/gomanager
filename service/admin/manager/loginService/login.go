package loginService

import (
	"apiserver/model/admin/managerModel"
	"apiserver/pkg/auth"
	"apiserver/pkg/global/redis"
	"apiserver/pkg/token"
	"bytes"
	"errors"
	redisgo "github.com/gomodule/redigo/redis"
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
	if u.Status == managerModel.FREEZE {
		return "", errors.New("账号冻结")
	}

	var roleIds []uint64
	if u.IsRoot == managerModel.ON {
		roleIds, err = managerModel.ListPermissionIds()
	} else {
		var ur managerModel.UserRoleModel
		ur.UserId = u.Id
		roleIds, err = ur.GetRoleIds()
	}
	if err != nil {
		return "", err
	}
	currentTime := time.Now()
	t, err := token.Sign(token.Context{ID: u.Id, Username: u.Username}, "")
	if err != nil {
		return "", err
	}
	//更新数据库
	var userUpdate managerModel.UserModel
	userUpdate.Id = u.Id
	userUpdate.LastTime = currentTime
	userUpdate.LastIp = ip
	if err := userUpdate.Update(); err != nil {
		return "", err
	}
	pool := redis.Pool.Pool.Get()
	defer pool.Close()
	uidStr := strconv.Itoa(int(u.Id))
	//token
	var key bytes.Buffer
	key.WriteString("user:login:")
	key.WriteString(uidStr)
	if _, err := pool.Do("Set", key.String(), t); err != nil {
		return "", err
	}
	//权限
	var permissionKey bytes.Buffer
	permissionKey.WriteString("user:permission:ids:")
	permissionKey.WriteString(uidStr)
	args := redisgo.Args{}.Add(permissionKey.String())
	for _, v := range roleIds {
		args = args.AddFlat(v)
	}
	//删除权限
	if _, err := pool.Do("DEL", permissionKey.String()); err != nil {
		return "", err
	}
	if _, err := pool.Do("SADD", args...); err != nil {
		return "", err
	}
	//删除菜单
	var menuKey bytes.Buffer
	menuKey.WriteString("user:menu:")
	menuKey.WriteString(uidStr)
	if _, err := pool.Do("DEL", menuKey.String()); err != nil {
		return "", err
	}
	return t, nil
}
