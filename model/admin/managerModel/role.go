package managerModel

import (
	"github.com/spf13/viper"
	"apiserver/model"
	"bytes"
	"strconv"
	globalModel "apiserver/pkg/global/model"
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
