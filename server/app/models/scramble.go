package models

import (
	"gorm.io/gorm"
	"puzzle/utils"
	"time"
)

type Scramble struct {
	ID        int64  `json:"id" gorm:"primaryKey"`    // 主键ID
	Dimension int    `json:"dimension"`               // 阶数 3 | 4 | 5 | 6
	IDx       int64  `json:"idx"`                     // 随机数索引
	Scramble  string `json:"scramble"`                // 打乱公式
	Status    int8   `json:"status" gorm:"default 1"` // 状态 1:启用 2:冻结 3:删除
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

type ScrambleReq struct {
	ID         int64            `json:"id"`        // 主键ID
	Dimension  int              `json:"dimension"` // 阶数 3 | 4 | 5 | 6
	IDx        int64            `json:"idx"`       // 随机数索引
	Scramble   string           `json:"scramble"`  // 打乱公式
	Status     int8             `json:"status"`    // 状态 1:启用 2:冻结 3:删除
	DateRange  []time.Time      `json:"dateRange"` // 时间范围
	Pagination utils.Pagination `gorm:"embedded"`  // 分页
	Sorted     string           `json:"sorted"`    // 排序
}

type ScrambleResp struct {
	ID        int64     `json:"id" gorm:"primaryKey"`    // 主键ID
	Dimension int       `json:"dimension"`               // 阶数 3 | 4 | 5 | 6
	IDx       int64     `json:"idx"`                     // 随机数索引
	Scramble  string    `json:"scramble"`                // 打乱公式
	Status    int8      `json:"status" gorm:"default 1"` // 状态 1:启用 2:冻结 3:删除
	CreatedAt time.Time `json:"createdAt"`               // 创建时间
	UpdatedAt time.Time `json:"updatedAt"`               // 更新时间
}

type ScrambleListResp struct {
	Total   int64          `json:"total"`   // 总数
	Records []ScrambleResp `json:"records"` // 打乱公式列表
}

type CreateScambleReq struct {
	UserID    int64 `json:"userId"`    // 用户ID
	Dimension int   `json:"dimension"` // 阶数 3 | 4 | 5 | 6
}
