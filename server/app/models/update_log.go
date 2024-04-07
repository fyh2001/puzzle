package models

import (
	"puzzle/utils"
	"time"
)

type UpdateLog struct {
	Id          int64     `json:"id" gorm:"primary_key"`           // 主键ID
	Version     string    `json:"version"`                         // 版本号
	Type        int       `json:"type"`                            // 类型 1:日常更新 2:重大更新
	Content     string    `json:"content"`                         // 更新内容
	Description string    `json:"description"`                     // 描述
	Status      int       `json:"status"`                          // 状态 0:启用 2:冻结 3:删除
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

type UpdateLogReq struct {
	Id      int64  `json:"-"`       // 主键ID
	Version string `json:"version"` // 版本号

	Type        int    `json:"type"`        // 类型 1:日常更新 2:重大更新
	Content     string `json:"content"`     // 更新内容
	Description string `json:"description"` // 描述
	Status      int    `json:"status"`      // 状态 0:启用 2:冻结 3:删除

	IdStr      string           `json:"id"`        // 主键ID
	Pagination utils.Pagination `gorm:"embedded"`  // 分页
	DateRange  []time.Time      `json:"dateRange"` // 日期范围
	Sorted     string           `json:"sorted"`    // 排序
	OrderBy    string           `json:"orderBy"`   // 排序字段
}

type UpdateLogResp struct {
	Id          string    `json:"id"`          // 主键ID
	Version     string    `json:"version"`     // 版本号
	Type        int       `json:"type"`        // 类型 1:日常更新 2:重大更新
	Content     string    `json:"content"`     // 更新内容
	Description string    `json:"description"` // 描述
	Status      int       `json:"status"`      // 状态 0:启用 2:冻结 3:删除
	CreatedAt   time.Time `json:"createdAt"`   // 创建时间
	UpdatedAt   time.Time `json:"updatedAt"`   // 更新时间
}

type UpdateLogListResp struct {
	Records []UpdateLogResp `json:"record"` // 列表
	Total   int64           `json:"total"`  // 总数
}
