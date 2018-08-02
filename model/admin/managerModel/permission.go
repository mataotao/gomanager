package managerModel

import (
	"apiserver/model"
	"github.com/spf13/viper"
	"apiserver/requests/admin/manager/permissionRequests"
	validator "gopkg.in/go-playground/validator.v9"
	"apiserver/pkg/constvar"
	"sync"
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

func ListPermission(limit uint64, page uint64) ([]*PermissionModel, uint64, error) {
	//默认值
	if limit == 0 {
		limit = constvar.DefaultLimit
	}
	//默认值
	if page == 0 {
		page = constvar.DefaultPage
	}
	//计算开始偏移量
	start := (page - 1) * limit
	//用map装数据
	permissionList := make([] *PermissionModel, 0)
	//总条数
	var total uint64
	//查询总条数
	if err := model.DB.Self.Model(&PermissionModel{}).Count(&total).Error; err != nil {
		return permissionList, total, err
	}

	//查询数据
	if err := model.DB.Self.Offset(start).Limit(limit).Order("id desc").Find(&permissionList).Error; err != nil {
		return permissionList, total, err
	}
	return permissionList, total, nil
}
//Lock锁 IdMap id为建  sync.Mutex 是因为在并发处理中，更新同一个变量为了保证数据一致性
type PermissionListLock struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*PermissionListInfo
}

type PermissionListInfo struct {
	Id            uint64 `json:"id"`
	Label         string `json:"label" `
	IsContainMenu uint8  `json:"is_contain_menu" `
	Pid           uint8  `json:"pid" `
	Level         uint8  `json:"level" `
	Url           string `json:"url" `
	Sort          uint64 `json:"sort" `
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
	SayHello      string `json:"sayHello"`
}
