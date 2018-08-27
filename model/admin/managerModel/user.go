package managerModel

import (
	"apiserver/model"
	"apiserver/pkg/auth"
	"apiserver/pkg/constvar"
	globalModel "apiserver/pkg/global/model"
	"github.com/spf13/viper"
	"runtime"
	"sync"
	"time"
)

const (
	ON uint8 = 1
)

type UserModel struct {
	model.BaseModel
	Username string    `json:"username"`
	Name     string    `json:"name"`
	Mobile   uint64    `json:"mobile"`
	Password string    `json:"password"`
	HeadImg  string    `json:"head_img"`
	LastTime time.Time `json:"last_time"`
	LastIp   string    `json:"last_ip"`
	IsRoot   uint8     `json:"is_root"`
	Status   uint8     `json:"status"`
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
	field := []string{"user_id", "role_id"}
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

func (u *UserModel) Updates(roles []uint64) error {
	tx := model.DB.Self.Begin()
	if err := tx.Model(&u).Updates(u).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("user_id = ?", u.Id).Delete(UserRoleModel{}).Error; err != nil {
		return err
	}
	roleData := make([][]int, len(roles))
	for i, v := range roles {
		t := []int{int(u.Id), int(v)}
		roleData[i] = t
	}
	var urm UserRoleModel
	field := []string{"user_id", "role_id"}
	sql := globalModel.MultiInsertIntSql(urm.TableName(), field, roleData)

	if err := tx.Exec(sql).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil

}

func (u *UserModel) Get() (*userInfo, error) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(2)
	finished := make(chan bool, 1)
	errChan := make(chan error, 1)
	go func(id uint64) {
		defer wg.Done()
		if err := model.DB.Self.First(&u, id).Error; err != nil {
			errChan <- err
		}
	}(u.Id)
	var roleIds []uint64
	var ur UserRoleModel
	roleTableName := ur.TableName()
	go func(id uint64) {
		defer wg.Done()
		if err := model.DB.Self.Table(roleTableName).Where("user_id = ? ", id).Pluck("role_id", &roleIds).Error; err != nil {
			errChan <- err
		}
	}(u.Id)

	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case <-finished:
	case err := <-errChan:
		return nil, err
	}
	info := &userInfo{
		Id:       u.Id,
		Username: u.Username,
		Name:     u.Name,
		Mobile:   u.Mobile,
		HeadImg:  u.HeadImg,
		Roles:    roleIds,
	}
	return info, nil
}

func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

func (u *UserModel) List(page, limit, roleId uint64) ([]*UserModel, map[uint64][]uint64, uint64, error) {
	var count uint64

	DB := model.DB.Self.Table(u.TableName())
	if u.Name != "" {
		DB = DB.Where("name = ?", u.Name)
	}
	if u.Username != "" {
		DB = DB.Where("username = ?", u.Username)
	}
	if u.Status != 0 {
		DB = DB.Where("status = ?", u.Status)
	}
	if roleId != 0 {
		var userIds []UserRoleModel
		if err := model.DB.Self.Where("role_id = ? ", roleId).Find(&userIds).Error; err != nil {
			return nil, nil, count, err
		}
		if len(userIds) > 0 {
			uIds := make([]uint64, len(userIds))
			for i, v := range userIds {
				uIds[i] = v.UserId
			}
			DB = DB.Where("id in (?)", uIds)
		}
	}

	if limit == 0 {
		limit = constvar.DefaultLimit
	}
	if page == 0 {
		page = constvar.DefaultPage
	}
	start := (page - 1) * limit

	var userInfoList []*UserModel
	if err := DB.Count(&count).Error; err != nil {
		return nil, nil, count, err
	}

	if err := DB.Offset(start).Limit(limit).Order("id desc").Find(&userInfoList).Error; err != nil {
		return nil, nil, count, err
	}
	currentUIds := make([]uint64, len(userInfoList))

	for i, v := range userInfoList {
		currentUIds[i] = v.Id
	}
	var userRole []UserRoleModel

	if err := model.DB.Self.Model(&UserRoleModel{}).Where("user_id in (?)", currentUIds).Find(&userRole).Error; err != nil {
		return nil, nil, count, err
	}
	userRoleIds := make(map[uint64][]uint64, 0)
	for _, v := range userRole {
		userRoleIds[v.UserId] = append(userRoleIds[v.UserId], v.RoleId)
	}
	return userInfoList, userRoleIds, count, nil
}

type userInfo struct {
	Id       uint64   `json:"id"`
	Username string   `json:"username"`
	Name     string   `json:"name"`
	Mobile   uint64   `json:"mobile"`
	HeadImg  string   `json:"head_img"`
	Roles    []uint64 `json:"roles"`
}

type UserListInfo struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Mobile   uint64 `json:"mobile"`
	HeadImg  string `json:"head_img"`
	LastTime string `json:"last_time"`
	LastIp   string `json:"last_ip"`
	IsRoot   uint8  `json:"is_root"`
	Status   uint8  `json:"status"`
	RoleName string `json:"role_name"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserListInfo
}
