package permissionService

import (
	"apiserver/model/admin/managerModel"
	"apiserver/util"
	"sync"
	"fmt"
)

func PermissionList(limit uint64, page uint64) ([]*managerModel.PermissionListInfo, uint64, error) {
	//查询数据库
	permissionList, total, err := managerModel.ListPermission(limit, page)
	if err != nil {
		return nil, total, err
	}
	//最后返回infos
	infos := make([]*managerModel.PermissionListInfo, 0)

	//ids 如果用goroutine排序会乱
	ids := []uint64{}

	//把正确顺序的id依次放入ids
	for _, v := range permissionList {
		ids = append(ids, v.Id)
	}

	//解决goroutine同步问题
	wg := sync.WaitGroup{}

	//sync.Mutex 是因为在并发处理中，更新同一个变量为了保证数据一致性
	permissions := managerModel.PermissionListLock{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*managerModel.PermissionListInfo, len(permissionList)),
	}

	//错误信息频道
	errChan := make(chan error, 1)
	//结束
	finished := make(chan bool, 1)

	for _, v := range permissionList {
		//每次循环加一个
		wg.Add(1)
		go func(p *managerModel.PermissionModel) {
			//执行结束减一
			defer wg.Done()

			shortId, err := util.GenShortId()
			if err != nil {
				//如果有错误错误添加到errChan频道
				errChan <- err
				return
			}
			//加锁
			permissions.Lock.Lock()
			//整个执行完成解锁
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
		//等待
		wg.Wait()
		//关闭频道
		close(finished)
	}()
	//判断哪个频道成立  // select 就是监听 IO 操作，当 IO 操作发生时，触发相应的动作,select有比较多的限制，其中最大的一条限制就是每个case语句里必须是一个IO操作，确切的说，应该是一个面向channel的IO操作。
	select {
	case <-finished:
	case err := <-errChan:
		return nil, total, err
	}
	//映射数据
	for _, id := range ids {
		infos = append(infos, permissions.IdMap[id])
	}

	return infos, total, nil
}

type ListResponse struct {
	Total           uint64                             `json:"total"`
	PermissionsList []*managerModel.PermissionListInfo `json:"list"`
}
