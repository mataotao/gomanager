package managerModel

import (
	"apiserver/model"
	"github.com/spf13/viper"
)

type UserRoleModel struct {
	model.BaseModel
	UserId uint64 `json:"user_id"`
	RoleId uint64 `json:"role_id"`
}

func (u *UserRoleModel) TableName() string {
	return viper.GetString("db.prefix") + "user_role"
}

func (u *UserRoleModel) GetRoleIds() ([]uint64, error) {
	var roleIds []uint64
	err := model.DB.Self.Table("manager_user_role as ur").Joins("LEFT JOIN manager_role_permission rp ON rp.role_id=ur.role_id").Where("ur.user_id = ?", u.UserId).Pluck("rp.permission_id", &roleIds)
	if err.Error != nil {
		return roleIds, err.Error
	}
	return roleIds, nil
}
