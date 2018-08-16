package user

type CreateRequest struct {
	Username string   `json:"username" valid:"required"`
	Name     string   `json:"name" valid:"required"`
	Mobile   uint64   `json:"mobile" valid:"required,length(11|11),numeric"`
	Password string   `json:"password" valid:"required"`
	HeadImg  string   `json:"head_img"`
	Roles    []uint64 `json:"roles" valid:"required"`
}

type FreezeRequest struct {
	Status uint8 `json:"status" valid:"required,length(1|1),numeric"`
}

type PwdRequest struct {
	Password string `json:"password" valid:"required,length(6|30)"`
}
