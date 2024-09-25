package models

import (
	"gorm.io/gorm"
	"puzzle/utils"
	"time"
)

type ScrambledUserStatus struct {
	ID         int64 `json:"id" gorm:"primaryKey"` // ID
	UserID     int64 `json:"userId"`               // 用户ID
	Dimension  int   `json:"dimension"`            // 阶数 3 | 4 | 5 | 6
	ScrambleID int64 `json:"scrambleId"`           // 打乱公式ID
	Status     int   `json:"status"`               // 完成状态 1:未完成 2:已完成
	DeletedAt  gorm.DeletedAt
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

type ScrambledUserStatusReq struct {
	ID         int64            `json:"id"`         // ID
	UserID     int64            `json:"userId"`     // 用户ID
	Dimension  int              `json:"dimension"`  // 阶数 3 | 4 | 5 | 6
	ScrambleID int64            `json:"scrambleId"` // 打乱公式ID
	Status     int              `json:"status"`     // 完成状态 1:未完成 2:已完成
	DateRange  []time.Time      `json:"dateRange"`  // 时间范围
	Pagination utils.Pagination `gorm:"embedded"`   // 分页
	Sorted     string           `json:"sorted"`     // 排序
}

type ScrambledUserStatusResp struct {
	ID         string    `json:"id"`         // ID
	UserID     int64     `json:"userId"`     // 用户ID
	Dimension  int       `json:"dimension"`  // 阶数 3 | 4 | 5 | 6
	ScrambleID int64     `json:"scrambleId"` // 打乱公式ID
	Status     int       `json:"status"`     // 完成状态 1:未完成 2:已完成
	CreatedAt  time.Time `json:"createdAt"`  // 创建时间
	UpdatedAt  time.Time `json:"updatedAt"`  // 更新时间
}

type ScrambledUserStatusListResp struct {
	Total   int64                     `json:"total"`   // 总数
	Records []ScrambledUserStatusResp `json:"records"` // 用户打乱公式状态列表
}
