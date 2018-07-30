package managerModel

import (
	"apiserver/model"
	"github.com/spf13/viper"
	validator "gopkg.in/go-playground/validator.v9"
)

type PermissionModel struct {
	model.BaseModel
	Label         string `json:"label" gorm:"column:label;not null" binding:"required" validate:"min=1,max=32"`
	IsContainMenu uint8  `json:"is_contain_menu" gorm:"column:is_contain_menu;not null" binding:"required" validate:"min=1,max=32"`
	Pid           uint8  `json:"pid" gorm:"column:pid;not null" `
	Level         uint8  `json:"level" gorm:"column:level;not null" binding:"required" validate:"required"`
	Url           string `json:"url" gorm:"column:url"`
	Sort          uint64 `json:"sort" gorm:"column:sort;default:'500'"`
}

//表名
func (p *PermissionModel) TableName() string {
	return viper.GetString("db.prefix") + "permission"
}

//验证
func (p *PermissionModel) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

///创建
func (p *PermissionModel) Create() error {
	return model.DB.Self.Create(&p).Error
}

//删除
func DeletePermission(id uint64) error {
	permission := PermissionModel{}
	permission.BaseModel.Id = id
	//Unscoped 方法是永久删除,因为数据表有delete_at字段,默认是软删除 用Unscoped永久删除
	return model.DB.Self.Unscoped().Delete(&permission).Error
}
//更新全部字段
func (p *PermissionModel) Update() error {
	return model.DB.Self.Save(p).Error
}
