package auth

import (
	"apiserver/pkg/global/redis"
	"bytes"
	redisgo "github.com/gomodule/redigo/redis"
	"strconv"
)

func Route(route string, id uint64) bool {
	pool := redis.Pool.Pool.Get()
	defer pool.Close()
	var key bytes.Buffer
	key.WriteString("user:permission:route:")
	key.WriteString(route)
	routeIds, err := redisgo.String(pool.Do("GET", key.String()))
	if err != nil {
		return false
	}
	var userKey bytes.Buffer
	userKey.WriteString("user:permission:ids:")
	userKey.WriteString(strconv.Itoa(int(id)))
	userAllPermission, err := redisgo.Bool(pool.Do("SISMEMBER", redisgo.Args{}.Add(userKey.String()).AddFlat(routeIds)...))
	if err != nil {
		return false
	}
	return userAllPermission
}
