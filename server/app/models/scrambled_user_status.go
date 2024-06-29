package models

import (
	"puzzle/utils"
	"time"
)

type ScrambledUserStatus struct {
	Id         int64     `json:"id" gorm:"primaryKey"`            // ID
	UserId     int64     `json:"userId"`                          // 用户ID
	Dimension  int       `json:"dimension"`                       // 阶数 3 | 4 | 5 | 6
	ScrambleId int64     `json:"scrambleId"`                      // 打乱公式ID
	Status     int       `json:"status"`                          // 完成状态 1:未完成 2:已完成
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

type ScrambledUserStatusReq struct {
	Id         int64            `json:"id"`         // ID
	UserId     int64            `json:"userId"`     // 用户ID
	Dimension  int              `json:"dimension"`  // 阶数 3 | 4 | 5 | 6
	ScrambleId int64            `json:"scrambleId"` // 打乱公式ID
	Status     int              `json:"status"`     // 完成状态 1:未完成 2:已完成
	DateRange  []time.Time      `json:"dateRange"`  // 时间范围
	Pagination utils.Pagination `gorm:"embedded"`   // 分页
	Sorted     string           `json:"sorted"`     // 排序
}

type ScrambledUserStatusResp struct {
	Id         string    `json:"id"`                              // ID
	UserId     int64     `json:"userId"`                          // 用户ID
	Dimension  int       `json:"dimension"`                       // 阶数 3 | 4 | 5 | 6
	ScrambleId int64     `json:"scrambleId"`                      // 打乱公式ID
	Status     int       `json:"status"`                          // 完成状态 1:未完成 2:已完成
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

type ScrambledUserStatusListResp struct {
	Total   int64                     `json:"total"`   // 总数
	Records []ScrambledUserStatusResp `json:"records"` // 用户打乱公式状态列表
}
