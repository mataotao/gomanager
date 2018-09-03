package loginService

import (
	"apiserver/model/admin/managerModel"
	"apiserver/pkg/auth"
	"apiserver/pkg/errno"
	"apiserver/pkg/global/redis"
	"apiserver/pkg/token"
	"bytes"
	redisgo "github.com/gomodule/redigo/redis"
	"github.com/lexkong/log"
	"strconv"
	"time"
)

func Login(username string, pwd string, ip string) (*managerModel.LoginInfo, error) {

	u, err := managerModel.GetUser(username)
	if err != nil {
		log.Error("login", err)
		return nil, errno.ErrUserNameOrPassword
	}

	if err := auth.Compare(u.Password, pwd); err != nil {
		log.Error("login", err)
		return nil, errno.ErrUserNameOrPassword
	}
	if u.Status == managerModel.FREEZE {
		return nil, errno.UserFreeze
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
		log.Error("login", err)
		return nil, errno.ErrUserNameOrPassword
	}
	currentTime := time.Now()
	t, err := token.Sign(token.Context{ID: u.Id, Username: u.Username}, "")
	if err != nil {
		log.Error("login", err)
		return nil, errno.ErrUserNameOrPassword
	}
	loginInfo := &managerModel.LoginInfo{
		Token:    t,
		Username: u.Username,
		IsRoot:   u.IsRoot,
		HeadImg:  u.HeadImg,
	}
	//更新数据库
	var userUpdate managerModel.UserModel
	userUpdate.Id = u.Id
	userUpdate.LastTime = currentTime
	userUpdate.LastIp = ip
	if err := userUpdate.Update(); err != nil {
		log.Error("login", err)
		return nil, errno.ErrUserNameOrPassword
	}
	pool := redis.Pool.Pool.Get()
	defer pool.Close()
	uidStr := strconv.Itoa(int(u.Id))
	//token
	var key bytes.Buffer
	key.WriteString("user:login:")
	key.WriteString(uidStr)
	loginStr := key.String()
	if _, err := pool.Do("Set", loginStr, t); err != nil {
		log.Error("login", err)
		return nil, errno.ErrUserNameOrPassword
	}
	//权限
	var permissionKey bytes.Buffer
	permissionKey.WriteString("user:permission:ids:")
	permissionKey.WriteString(uidStr)
	permissionKeyStr := permissionKey.String()
	args := redisgo.Args{}.Add(permissionKeyStr)
	for _, v := range roleIds {
		args = args.AddFlat(v)
	}
	//删除权限
	if _, err := pool.Do("DEL", permissionKeyStr); err != nil {
		log.Error("login", err)
		return nil, errno.ErrUserNameOrPassword
	}
	if _, err := pool.Do("SADD", args...); err != nil {
		log.Error("login", err)
		return nil, errno.ErrUserNameOrPassword
	}
	//删除菜单
	var menuKey bytes.Buffer
	menuKey.WriteString("user:menu:")
	menuKey.WriteString(uidStr)
	if _, err := pool.Do("DEL", menuKey.String()); err != nil {
		log.Error("login", err)
		return nil, errno.ErrUserNameOrPassword
	}
	//设置有效时间
	if _, err = pool.Do("EXPIRE", loginStr, 20*60); err != nil {
		log.Error("login", err)
		return nil, errno.ErrUserNameOrPassword
	}
	if _, err = pool.Do("EXPIRE", permissionKeyStr, 20*60); err != nil {
		log.Error("login", err)
		return nil, errno.ErrUserNameOrPassword
	}
	return loginInfo, nil
}
