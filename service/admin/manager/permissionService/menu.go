package permissionService

import (
	"apiserver/model/admin/managerModel"
	"apiserver/pkg/global/auth"
	"strconv"
)

func Menu(uid uint64) ([]managerModel.MenuInfo, error) {
	//查询数据库
	permissionList, err := managerModel.ListPermission()
	if err != nil {
		return nil, err
	}

	infos := menuTree(0, permissionList, uid)

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
