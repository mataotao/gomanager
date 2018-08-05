package permissionService

import (
	"apiserver/model/admin/managerModel"
)

func PermissionList() ([]managerModel.PermissionListInfo, error) {
	//查询数据库
	permissionList, err := managerModel.ListPermission()
	if err != nil {
		return nil, err
	}

	infos := tree(0, permissionList)

	return infos, nil
}

//递归实现
func tree(pid uint64, permissionList []*managerModel.PermissionModel) []managerModel.PermissionListInfo {
	var arr []managerModel.PermissionListInfo
	for _, v := range permissionList {
		if pid == v.Pid {
			pTree := managerModel.PermissionListInfo{}
			pTree.Id = v.Id
			pTree.Label = v.Label
			pTree.IsContainMenu = v.IsContainMenu
			pTree.Pid = v.Pid
			pTree.Url = v.Url
			pTree.Level = v.Level
			pTree.Sort = v.Sort
			subTree := tree(v.Id, permissionList)
			pTree.Children = subTree
			arr = append(arr, pTree)
		}
	}
	return arr
}
