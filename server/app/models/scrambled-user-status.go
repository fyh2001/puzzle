package models

import (
	"puzzle/utils"
	"time"
)

type ScrambledUserStatus struct {
	UserId     int64     `json:"user_id"`                          // 用户ID
	Dimension  int       `json:"dimension"`                        // 阶数 3 | 4 | 5 | 6
	ScrambleId int64     `json:"scramble_id"`                      // 打乱公式ID
	Status     int       `json:"status"`                           // 完成状态 0:未完成 1:已完成
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"` // 更新时间
}

type ScrambledUserStatusReq struct {
	UserId     int64            `json:"user_id"`     // 用户ID
	Dimension  int              `json:"dimension"`   // 阶数 3 | 4 | 5 | 6
	ScrambleId int64            `json:"scramble_id"` // 打乱公式ID
	Status     int              `json:"status"`      // 完成状态 0:未完成 1:已完成
	DateRange  []time.Time      `json:"date_range"`  // 时间范围
	Pagination utils.Pagination `gorm:"embedded"`    // 分页
	Sorted     string           `json:"sorted"`      // 排序
}

type ScrambledUserStatusResp struct {
	UserId     int64     `json:"user_id"`                          // 用户ID
	Dimension  int       `json:"dimension"`                        // 阶数 3 | 4 | 5 | 6
	ScrambleId int64     `json:"scramble_id"`                      // 打乱公式ID
	Status     int       `json:"status"`                           // 完成状态 0:未完成 1:已完成
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"` // 更新时间
}

type ScrambledUserStatusListResp struct {
	Total   int64                     `json:"total"`   // 总数
	Records []ScrambledUserStatusResp `json:"records"` // 用户打乱公式状态列表
}
