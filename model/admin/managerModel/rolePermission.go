package managerModel

import (
	"apiserver/model"
	"github.com/spf13/viper"
)

type RolePermissionModel struct {
	model.BaseModel
	RoleId       int `json:"role_id"`
	PermissionId int `json:"permission_id"`
}

func (r *RolePermissionModel) TableName() string {
	return viper.GetString("db.prefix") + "role_permission"
}
