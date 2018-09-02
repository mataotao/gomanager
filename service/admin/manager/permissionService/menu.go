package permissionService

import (
	"apiserver/model/admin/managerModel"
	"apiserver/pkg/global/auth"
	"apiserver/pkg/global/redis"
	"bytes"
	"encoding/json"
	redisgo "github.com/gomodule/redigo/redis"
	"strconv"
)

func Menu(uid uint64) ([]managerModel.MenuInfo, error) {
	//查询redis
	pool := redis.Pool.Pool.Get()
	defer pool.Close()
	var key bytes.Buffer
	key.WriteString("user:menu:")
	key.WriteString(strconv.Itoa(int(uid)))
	menuKey := key.String()
	menus, err := redisgo.String(pool.Do("GET", menuKey))
	if menus != "" {
		var menuSlice []managerModel.MenuInfo
		json.Unmarshal([]byte(menus), &menuSlice)
		return menuSlice, nil
	}
	//查询数据库
	permissionList, err := managerModel.ListPermission()
	if err != nil {
		return nil, err
	}

	infos := menuTree(0, permissionList, uid)
	jsonData, err := json.Marshal(infos)

	if _, err := pool.Do("SET", menuKey, jsonData); err != nil {
		return nil, err
	}

	return infos, nil
}

//递归实现
func menuTree(pid uint64, permissionList []*managerModel.PermissionModel, uid uint64) []managerModel.MenuInfo {
	var arr []managerModel.MenuInfo
	for _, v := range permissionList {
		if auth.Permission(v.Id, uid) == false {
			continue
		}
		if pid == v.Pid {
			pTree := managerModel.MenuInfo{}
			pTree.Icon = v.Icon
			pTree.Title = v.Label
			if v.IsContainMenu == managerModel.ON {
				pTree.Index = strconv.Itoa(int(v.Id))
			} else {
				pTree.Index = v.Url
			}
			subTree := menuTree(v.Id, permissionList, uid)
			pTree.Subs = subTree
			arr = append(arr, pTree)
		}
	}
	return arr
}
