package auth

import (
	"apiserver/pkg/global/redis"
	"bytes"
	redisgo "github.com/gomodule/redigo/redis"
	"strconv"
)

func Permission(pid uint64, uid uint64) bool {
	var key bytes.Buffer
	key.WriteString("user:permission:ids:")
	key.WriteString(strconv.Itoa(int(uid)))
	pool := redis.Pool.Pool.Get()
	defer pool.Close()
	res, err := redisgo.Bool(pool.Do("SISMEMBER", redisgo.Args{}.Add(key.String()).AddFlat(pid)...))
	if err != nil {
		return false
	}
	return res
}
