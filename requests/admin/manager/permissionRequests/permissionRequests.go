package permissionRequests

import (
	"time"
)

type CreateRequest struct {
	Label         string `json:"label"`
	IsContainMenu uint8  `json:"iscontainmenu"`
	Pid           uint64 `json:"pid"`
	Url           string `json:"url"`
	Level         uint8  `json:"level"`
	Sort          uint64 `json:"sort"`
	Cond          string `json:"cond"`
	Icon          string `json:"icon"`
}

type CreateResponse struct {
	Label string `json:"label"`
}

type UpdateRequest struct {
	Label         string `json:"label" valid:"required"`
	Sort          uint64 `json:"sort"`
	IsContainMenu uint8  `json:"iscontainmenu"`
	Url           string `json:"url"`
	Cond          string `json:"cond"`
	Icon          string `json:"icon"`
	UpdatedAt     time.Time
}
