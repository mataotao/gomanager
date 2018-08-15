package managerModel

import (
	"apiserver/model"
	"github.com/spf13/viper"
)

type UserRoleModel struct {
	model.BaseModel
	MemberId uint64 `json:"member_id"`
	RoleId   uint64 `json:"name"`
}

func (u *UserRoleModel) TableName() string {
	return viper.GetString("db.prefix") + "user_role"
}
