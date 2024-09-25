package models

import (
	"gorm.io/gorm"
	"puzzle/utils"
	"time"
)

type Accolade struct {
	ID        int    `json:"id" gorm:"primaryKey"` // 主键ID
	Name      string `json:"name"`                 // 称号名称
	Icon      string `json:"icon"`                 // 称号图标
	Status    int8   `json:"status"`               // 状态 1:启用 2:冻结 3:删除
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

type AccoladeReq struct {
	ID         int              `json:"-"`         // 主键ID
	Name       string           `json:"name"`      // 称号名称
	Icon       string           `json:"icon"`      // 称号图标
	Status     int8             `json:"status"`    // 状态 1:启用 2:冻结 3:删除
	Pagination utils.Pagination `gorm:"embedded"`  // 分页
	DateRange  []time.Time      `json:"dateRange"` // 时间范围
	Ids        []int            `json:"ids"`       // 主键ID集合
}

type AccoladeResp struct {
	ID        int       `json:"id"`        // 主键ID
	Name      string    `json:"name"`      // 称号名称
	Icon      string    `json:"icon"`      // 称号图标
	Status    int8      `json:"status"`    // 状态 1:启用 2:冻结 3:删除
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 更新时间
}

type AccoladeListResp struct {
	Total   int64          `json:"total"`   // 总数
	Records []AccoladeResp `json:"records"` // 称号列表
}
