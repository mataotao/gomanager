package permissionService

import (
	"apiserver/model/admin/managerModel"
	"apiserver/util"
	"sync"
	"fmt"
)

func PermissionList(limit uint64, page uint64) ([]*managerModel.PermissionListInfo, uint64, error) {
	permissionList, total, err := managerModel.ListPermission(limit, page)
	if err != nil {
		return nil, total, err
	}
	infos := make([]*managerModel.PermissionListInfo, 0)

	ids := []uint64{}

	for _, v := range permissionList {
		ids = append(ids, v.Id)
	}

	//解决goroutine同步问题
	wg := sync.WaitGroup{}

	permissions := managerModel.PermissionListLock{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*managerModel.PermissionListInfo, len(permissionList)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	for _, v := range permissionList {
		//每次循环加一个
		wg.Add(1)
		go func(p *managerModel.PermissionModel) {
			//执行结束减一
			defer wg.Done()

			shortId, err := util.GenShortId()
			if err != nil {
				errChan <- err
				return
			}

			permissions.Lock.Lock()
			defer permissions.Lock.Unlock()

			permissions.IdMap[p.Id] = &managerModel.PermissionListInfo{
				Id:            p.Id,
				Label:         p.Label,
				IsContainMenu: p.IsContainMenu,
				Pid:           p.Pid,
				Level:         p.Level,
				Url:           p.Url,
				Sort:          p.Sort,
				CreatedAt:     p.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:     p.UpdatedAt.Format("2006-01-02 15:04:05"),
				SayHello:      fmt.Sprintf("Hello %s", shortId),
			}
		}(v)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, total, err
	}

	for _, id := range ids {
		infos = append(infos, permissions.IdMap[id])
	}

	return infos, total, nil
}

type ListResponse struct {
	Total uint64 `json:"total"`
	PermissionsList []*managerModel.PermissionListInfo `json:"list"`
}