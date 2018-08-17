package managerModel

import (
	"apiserver/model"
	"github.com/spf13/viper"
)

type UserRoleModel struct {
	model.BaseModel
	UserId uint64 `json:"user_id"`
	RoleId   uint64 `json:"role_id"`
}

func (u *UserRoleModel) TableName() string {
	return viper.GetString("db.prefix") + "user_role"
}
