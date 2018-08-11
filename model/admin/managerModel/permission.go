package managerModel

import (
	"apiserver/model"
	"github.com/spf13/viper"
	"apiserver/requests/admin/manager/permissionRequests"
	validator "gopkg.in/go-playground/validator.v9"
)

type PermissionModel struct {
	model.BaseModel
	Label         string `json:"label" gorm:"column:label;not null" binding:"required" validate:"min=1,max=32"`
	IsContainMenu uint8  `json:"is_contain_menu" gorm:"column:is_contain_menu;not null" binding:"required"`
	Pid           uint64 `json:"pid" gorm:"column:pid;not null" `
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
	permissionModel := PermissionModel{}
	permissionModel.BaseModel.Id = id
	//Unscoped 方法是永久删除,因为数据表有delete_at字段,默认是软删除 用Unscoped永久删除
	return model.DB.Self.Unscoped().Delete(&permissionModel).Error
}

//更新全部字段
func (p *PermissionModel) UpdateAll() error {
	return model.DB.Self.Save(p).Error
}

//修改指定字段
func (p *PermissionModel) Update(data *permissionRequests.UpdateRequest) error {
	return model.DB.Self.Model(p).Updates(data).Error
}

//查询一条数据
func GetPermission(id uint64) (*PermissionModel, error) {
	p := &PermissionModel{}
	d := model.DB.Self.Where("id = ?", id).First(&p)
	return p, d.Error
}

func ListPermission() ([]*PermissionModel, error) {
	permissionList := make([] *PermissionModel, 0)
	//查询数据
	if err := model.DB.Self.Order("pid asc,sort desc,id asc").Find(&permissionList).Error; err != nil {
		return permissionList, err
	}
	return permissionList, nil
}

type PermissionListInfo struct {
	Id            uint64               `json:"id"`
	Label         string               `json:"label" `
	IsContainMenu uint8                `json:"is_contain_menu" `
	Pid           uint64               `json:"pid" `
	Level         uint8                `json:"level" `
	Url           string               `json:"url" `
	Sort          uint64               `json:"sort" `
	CreatedAt     string               `json:"created_at"`
	UpdatedAt     string               `json:"updated_at"`
	Children      []PermissionListInfo `json:"children"`
}
