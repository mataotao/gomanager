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

type UpdateRequest struct {
	Name    string   `json:"name" valid:"required"`
	Mobile  uint64   `json:"mobile" valid:"required,length(11|11),numeric"`
	HeadImg string   `json:"head_img"`
	Roles   []uint64 `json:"roles" valid:"required"`
}

type ListRequest struct {
	Page     uint64 `json:"page"`
	Limit    uint64 `json:"limit"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Status   uint8  `json:"name"`
	RoleId   uint64 `json:"role_id"`
}
