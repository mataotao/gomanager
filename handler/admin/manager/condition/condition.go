package condition

import (
	"apiserver/handler"
	"apiserver/model/admin/managerModel"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"reflect"
	"strconv"
	"sync"
)

func Condition(c *gin.Context) {
	var r Conditions
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.Error, nil)
		return
	}

	conds := GetFieldName(r)
	res := make(map[string]interface{}, 0)
	wg := sync.WaitGroup{}
	finished := make(chan bool, 1)
	errChan := make(chan error, 1)
	for k, cond := range conds {
		wg.Add(1)
		go func(key string, v interface{}) {
			defer wg.Done()
			if len(v.([]string)) > 0 {
				switch key {
				case "Role":
					role, err := Role(v.([]string))
					if err != nil {
						errChan <- err
						return
					}
					res["role"] = role
				case "User":
					user, err := User(v.([]string))
					if err != nil {
						errChan <- err
						return
					}
					res["user"] = user
				}
			}
		}(k, cond)

	}
	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case <-finished:
	case err := <-errChan:
		log.Error("condition", err)
		handler.SendResponse(c, err, res)
		return

	}
	handler.SendResponse(c, nil, res)
}

type Conditions struct {
	Role []string `json:"role"`
	User []string `json:"user"`
}

type Response struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//获取名称跟方法
func GetFieldName(structName Conditions) map[string]interface{} {
	t := reflect.TypeOf(structName)
	v := reflect.ValueOf(structName)
	result := make(map[string]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() { //判断是否为可导出字段
			result[t.Field(i).Name] = v.Field(i).Interface()
		}
	}

	return result
}

//角色
func Role(c []string) (map[string]interface{}, error) {
	res := make(map[string]interface{}, len(c))
	for _, v := range c {
		switch v {
		//角色列表
		case "list":
			var list managerModel.RoleModel
			l, err := list.All()
			if err != nil {
				return res, err
			}
			r := make([]Response, len(l))
			for k, role := range l {
				t := Response{strconv.Itoa(int(role.Id)), role.Name}
				r[k] = t
			}
			res["list"] = r

		}
	}
	return res, nil
}

func User(c []string) (map[string]interface{}, error) {
	res := make(map[string]interface{}, len(c))
	for _, v := range c {
		switch v {
		//用户状态
		case "status":
			status := make([]map[string]string, 0)
			initOne := map[string]string{"key": "1", "value": "正常"}
			initTwo := map[string]string{"key": "2", "value": "冻结"}
			status = append(status, initOne, initTwo)
			res["status"] = status

		}
	}
	return res, nil
}
