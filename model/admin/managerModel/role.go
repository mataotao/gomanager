package managerModel

import (
	"github.com/spf13/viper"
	"apiserver/model"
	"bytes"
	"strconv"
	globalModel "apiserver/pkg/global/model"
	"apiserver/pkg/constvar"
	"fmt"
	"sync"
)

type RoleModel struct {
	model.BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *RoleModel) TableName() string {
	return viper.GetString("db.prefix") + "role"
}

func (r *RoleModel) Create(p []int) error {
	//开启事务
	tx := model.DB.Self.Begin()
	if err := tx.Create(&r).Error; err != nil {
		tx.Rollback()
		return err
	}
	l := len(p) - 1
	var rp RolePermissionModel
	//字符串拼接用bytes.Buffer 性能非常好
	var buffer bytes.Buffer
	buffer.WriteString("INSERT INTO ")
	buffer.WriteString(rp.TableName())
	buffer.WriteString("(`role_id`,`permission_id`) VALUES ")
	for i, v := range p {
		rId := strconv.Itoa(int(r.Id))
		pId := strconv.Itoa(v)
		buffer.WriteString("(")
		buffer.WriteString(rId)
		buffer.WriteString(",")
		buffer.WriteString(pId)
		buffer.WriteString(")")
		if i < l {
			buffer.WriteString(",")
		}
	}
	//转化成字符串
	sql := buffer.String()
	if err := tx.Exec(sql).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *RoleModel) Delete() error {
	tx := model.DB.Self.Begin()
	if err := tx.Delete(&r).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("role_id = ?", r.Id).Delete(RolePermissionModel{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *RoleModel) Update(data *RoleModel, p []int) error {
	tx := model.DB.Self.Begin()
	if err := tx.Model(&r).Updates(data).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("role_id = ?", r.Id).Delete(RolePermissionModel{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	var rpm RolePermissionModel
	field := []string{"`role_id`", "`permission_id`"}
	roleData := make([][]int, len(p))
	id := int(r.Id)
	for i, v := range p {
		t := []int{id, v}
		roleData[i] = t
	}
	sql := globalModel.MultiInsertIntSql(rpm.TableName(), field, roleData)
	if err := tx.Exec(sql).Error; err != nil {
		return err
	}
	tx.Commit()
	return nil

}

func (r *RoleModel) Get() (InfoResponse, error) {
	info := make([]info, 0)
	err := model.DB.Self.Table("manager_role as r").Select("r.name, r.description,rp.permission_id").Joins("LEFT JOIN manager_role_permission rp ON rp.role_id=r.id").Where("r.id = ?", r.Id).Scan(&info)
	var ir InfoResponse
	ir.Name = info[0].Name
	ir.Description = info[0].Description
	ir.Id = r.Id
	for _, v := range info {
		ir.Permission = append(ir.Permission, v.PermissionId)
	}
	return ir, err.Error
}

func (r *RoleModel) List(page uint64, limit uint64) ([]*RoleModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}
	if page == 0 {
		page = constvar.DefaultPage
	}
	start := (page - 1) * limit
	var count uint64
	list := make([]*RoleModel, 0)
	where := fmt.Sprintf("name like '%%%s%%'", r.Name)

	if err := model.DB.Self.Model(&RoleModel{}).Where(where).Count(&count).Error; err != nil {
		return list, count, err
	}

	if err := model.DB.Self.Model(&RoleModel{}).Where(where).Offset(start).Limit(limit).Order("id desc").Find(&list).Error; err != nil {
		return list, count, err
	}
	return list, count, nil

}

type info struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	PermissionId int    `json:"permission_id"`
}

type InfoResponse struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Permission  []int  `json:"permission"`
}

type RoleList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*RoleListInfo
}

type RoleListInfo struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
