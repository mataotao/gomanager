package roleService

import (
	"apiserver/model/admin/managerModel"
	"sync"
)

func List(name string, page uint64, limit uint64) ([]*managerModel.RoleListInfo, uint64, error) {
	roleModel := managerModel.RoleModel{
		Name: name,
	}
	roles, count, err := roleModel.List(page, limit)
	if err != nil {
		return nil, count, err
	}

	resRoleList := make([]*managerModel.RoleListInfo, len(roles))
	//goroutine会乱序
	ids := []uint64{}
	for _, role := range roles {
		ids = append(ids, role.Id)
	}
	wg := sync.WaitGroup{}
	roleList := managerModel.RoleList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*managerModel.RoleListInfo, len(roles)),
	}
	finished := make(chan bool, 1)
	for _, v := range roles {
		wg.Add(1)
		go func(r *managerModel.RoleModel) {
			defer wg.Done()
			//加锁，防止读写同时
			roleList.Lock.Lock()
			defer roleList.Lock.Unlock()
			roleList.IdMap[r.Id] = &managerModel.RoleListInfo{
				Id:          r.Id,
				Name:        r.Name,
				Description: r.Description,
				CreatedAt:   r.CreatedAt.Format("2006-01-02 15:04:05"),
			}

		}(v)
	}
	go func() {
		//等待
		wg.Wait()
		close(finished)
	}()
	//判断哪个频道成立
	// select有比较多的限制，其中最大的一条限制就是每个case语句里必须是一个IO操作，确切的说，应该是一个面向channel的IO操作。
	select {
	case <-finished:

	}
	for i, v := range ids {
		resRoleList[i] = roleList.IdMap[v]
	}
	return resRoleList, count, nil
}
