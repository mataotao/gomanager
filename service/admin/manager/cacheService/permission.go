package cacheService

import (
	"apiserver/model/admin/managerModel"
	"apiserver/pkg/global/redis"
	"bytes"
	redisgo "github.com/gomodule/redigo/redis"
	"strconv"
	"strings"
	"sync"
)

func Permission() error {
	list, err := managerModel.ListPermission()
	if err != nil {
		return err
	}
	/**
	     Mutex（互斥锁）
    Mutex 为互斥锁，Lock() 加锁，Unlock() 解锁
    在一个 goroutine 获得 Mutex 后，其他 goroutine 只能等到这个 goroutine 释放该 Mutex
    使用 Lock() 加锁后，不能再继续对其加锁，直到利用 Unlock() 解锁后才能再加锁
    在 Lock() 之前使用 Unlock() 会导致 panic 异常
    已经锁定的 Mutex 并不与特定的 goroutine 相关联，这样可以利用一个 goroutine 对其加锁，再利用其他 goroutine 对其解锁
    在同一个 goroutine 中的 Mutex 解锁之前再次进行加锁，会导致死锁
    适用于读写不确定，并且只有一个读或者写的场景
	 */
	mutex := new(sync.Mutex)
	wg := sync.WaitGroup{}
	pool := redis.Pool.Pool.Get()
	defer pool.Close()
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)
	searchData, err := redisgo.Strings(pool.Do("KEYS", "user:permission:*"))
	if err != nil {
		return err
	}
	if len(searchData) > 0 {
		if err != nil {
			return err
		}
		delKeys := redisgo.Args{}
		for _, dk := range searchData {
			delKeys = delKeys.Add(dk)
		}
		//批量删除key
		_, err = pool.Do("DEL", delKeys...)
		if err != nil {
			return err
		}

	}
	for _, v := range list {
		wg.Add(1)
		go func(p *managerModel.PermissionModel) {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			if p.Cond != "" {
				condArr := strings.Split(p.Cond, ",")
				for _, per := range condArr {
					var key bytes.Buffer
					key.WriteString("user:permission:")
					key.WriteString(per)
					data, err := pool.Do("GET", key.String())
					if err != nil {
						errChan <- err
						return
					}
					value := strconv.Itoa(int(p.Id))
					if data != nil {
						dataStr, err := redisgo.String(pool.Do("GET", key.String()))
						if err != nil {
							errChan <- err
							return
						}
						var datas bytes.Buffer
						datas.WriteString(dataStr)
						datas.WriteString(",")
						datas.WriteString(value)
						value = datas.String()
					}
					_, err = pool.Do("SET", key.String(), value)
					if err != nil {
						errChan <- err
						return
					}
				}
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
		return err
	}
	return nil

}
