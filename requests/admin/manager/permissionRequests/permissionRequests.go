package permissionRequests

import "time"

type CreateRequest struct {
	Label         string `json:"label"`
	IsContainMenu uint8  `json:"is_contain_menu"`
	Pid           uint8  `json:"pid"`
	Url           string `json:"url"`
	Level         uint8  `json:"level"`
	Sort          uint64 `json:"sort"`
}

type CreateResponse struct {
	Label string `json:"label"`
}

type UpdateRequest struct {
	Label     string `json:"label" valid:"required"`
	UpdatedAt time.Time
}
