package managerModel

import (
	"apiserver/pkg/auth"
	"apiserver/model"
	"github.com/spf13/viper"
	globalModel "apiserver/pkg/global/model"
)

const (
	ON uint8 = 1
)

type UserModel struct {
	model.BaseModel
	Username string `json:"username"`
	Name     string `json:"name"`
	Mobile   uint64 `json:"mobile"`
	Password string `json:"password"`
	HeadImg  string `json:"head_img"`
	LastTime string `json:"last_time"`
	LastIp   string `json:"last_ip"`
	IsRoot   uint8  `json:"is_root"`
	Status   uint8  `json:"status"`
}

func (u *UserModel) TableName() string {
	return viper.GetString("db.prefix") + "user"
}

func (u *UserModel) Uinque() bool {
	user := &UserModel{}
	err := model.DB.Self.Where("username = ?", u.Username).First(&user).Error
	if err != nil {
		return true
	}
	return false
}

func (u *UserModel) Create(roleIds []uint64) error {
	tx := model.DB.Self.Begin()
	if err := tx.Create(&u).Error; err != nil {
		tx.Rollback()
		return err
	}
	roleData := make([][]int, len(roleIds))

	for i, v := range roleIds {
		t := []int{int(u.Id), int(v)}
		roleData[i] = t
	}
	field := []string{"member_id", "role_id"}
	var urm UserRoleModel
	sql := globalModel.MultiInsertIntSql(urm.TableName(), field, roleData)
	if err := tx.Exec(sql).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil

}

func (u *UserModel) Update() error {
	return model.DB.Self.Model(&u).Updates(u).Error
}
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}
