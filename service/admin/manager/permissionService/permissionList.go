package permissionService

import (
	"apiserver/model/admin/managerModel"
	"apiserver/pkg/global/auth"
)

func PermissionList(uid uint64) ([]managerModel.PermissionListInfo, error) {
	//查询数据库
	permissionList, err := managerModel.ListPermission()
	if err != nil {
		return nil, err
	}

	infos := tree(0, permissionList, uid)

	return infos, nil
}

//递归实现
func tree(pid uint64, permissionList []*managerModel.PermissionModel, uid uint64) []managerModel.PermissionListInfo {
	var arr []managerModel.PermissionListInfo
	for _, v := range permissionList {
		if auth.Permission(v.Id, uid) == false {
			continue
		}
		if pid == v.Pid {
			pTree := managerModel.PermissionListInfo{}
			pTree.Id = v.Id
			pTree.Label = v.Label
			pTree.IsContainMenu = v.IsContainMenu
			pTree.Pid = v.Pid
			pTree.Url = v.Url
			pTree.Level = v.Level
			pTree.Sort = v.Sort
			subTree := tree(v.Id, permissionList, uid)
			pTree.Children = subTree
			arr = append(arr, pTree)
		}
	}
	return arr
}
