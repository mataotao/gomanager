package userService

import (
	"apiserver/model/admin/managerModel"
	//"sync"
	"sync"
	"errors"
	"bytes"
)

func List(cond map[string]interface{}) (*ListResponse, error) {
	var u managerModel.UserModel
	u.Name = cond["name"].(string)
	u.Username = cond["username"].(string)
	u.Status = cond["status"].(uint8)
	//管理员信息
	infos, roleIds, count, err := u.List(cond["page"].(uint64), cond["limit"].(uint64), cond["roleId"].(uint64))
	if err != nil {
		return nil, err
	}
	userLen := len(infos)
	ids := make([]uint64, userLen)
	for i, v := range infos {
		ids[i] = v.Id
	}
	var r managerModel.RoleModel
	//角色信息
	roleInfos, err := r.All()
	if err != nil {
		return nil, err
	}

	roleMaps := make(map[uint64]*managerModel.RoleModel)

	for _, v := range roleInfos {
		roleMaps[v.Id] = v
	}
	userList := managerModel.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*managerModel.UserListInfo, userLen),
	}
	wg := sync.WaitGroup{}
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)
	for _, v := range infos {
		wg.Add(1)
		go func(u *managerModel.UserModel) {
			defer wg.Done()
			rId, ok := roleIds[u.Id]
			if ok == false {
				errChan <- errors.New("not key")
				return
			}
			var roleNames bytes.Buffer
			for _, id := range rId {
				rn, ok := roleMaps[id]
				if ok == false {
					errChan <- errors.New("not key")
					return
				}
				roleNames.WriteString(rn.Name)
				roleNames.WriteString(" ")
			}
			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IdMap[u.Id] = &managerModel.UserListInfo{
				Id:       u.Id,
				Username: u.Username,
				Name:     u.Name,
				Mobile:   u.Mobile,
				HeadImg:  u.HeadImg,
				LastTime: u.LastTime.Format("2006-01-02 15:04:05"),
				LastIp:   u.LastIp,
				IsRoot:   u.IsRoot,
				Status:   u.Status,
				RoleName: roleNames.String(),
			}
		}(v)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()
	/**
	select 语句使一个 Go 程可以等待多个通信操作。
    select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行
	select 有比较多的限制，其中最大的一条限制就是每个case语句里必须是一个IO操作，确切的说，应该是一个面向channel的IO操作。
	 */
	select {
	case <-finished:
	case err := <-errChan:
		return nil, err

	}
	currentUserList := make([]*managerModel.UserListInfo, userLen)

	for i, id := range ids {
		currentUserList[i] = userList.IdMap[id]
	}

	res := &ListResponse{
		Count: count,
		List:  currentUserList,
	}

	return res, nil
}

type ListResponse struct {
	Count uint64                       `json:"count"`
	List  []*managerModel.UserListInfo `json:"list"`
}
