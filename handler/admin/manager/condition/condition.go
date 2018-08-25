package condition

import (
	"apiserver/handler"
	"github.com/gin-gonic/gin"
	"apiserver/model/admin/managerModel"
	"reflect"
)

func Condition(c *gin.Context) {
	var r Conditions
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, err, nil)
	}

	conds := GetFieldName(r)
	res := make(map[string]interface{}, 0)
	for k, v := range conds {
		if len(v.([]string)) > 0 {
			switch k {
			case "Role":
				res["Role"] = Role(v.([]string))
			case "User":
				res["User"] = User(v.([]string))
			}
		}
	}
	handler.SendResponse(c, nil, res)
}

type Conditions struct {
	Role []string `json:"role"`
	User []string `json:"user"`
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
func Role(c []string) map[string]interface{} {
	res := make(map[string]interface{}, len(c))
	for _, v := range c {
		switch v {
		//角色列表
		case "list":
			var list managerModel.RoleModel
			l, _ := list.All()
			res["list"] = l

		}
	}
	return res
}

func User(c []string) map[string]interface{} {
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
	return res
}
