package auth

import (
	"apiserver/pkg/global/redis"
	"bytes"
	redisgo "github.com/gomodule/redigo/redis"
	"strconv"
	"strings"
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
	rids := make([]string, 0)
	if strings.Contains(routeIds, ",") {
		rids = strings.Split(routeIds, ",")
	} else {
		rids = append(rids, routeIds)
	}
	res := false
	for _, v := range rids {
		userAllPermission, err := redisgo.Bool(pool.Do("SISMEMBER", redisgo.Args{}.Add(userKey.String()).AddFlat(v)...))
		if err == nil && userAllPermission == true {
			res = true
		}
	}
	return res
}
