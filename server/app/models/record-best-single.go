package models

import (
	"puzzle/utils"
	"time"
)

// RecordBestSingle 最佳单次记录模型
type RecordBestSingle struct {
	UserId         int64     `json:"userId" gorm:"primaryKey"`        // 用户ID
	Dimension      int       `json:"dimension"`                       // 阶数 3 | 4 | 5 | 6 | 7 | 8
	RecordId       int64     `json:"recordId"`                        // 记录ID
	RecordDuration int       `json:"recordDuration"`                  // 耗时
	RecordStep     int       `json:"recordStep"`                      // 步数
	CreatedAt      time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt      time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

// RecordBestSingleReq 最佳单次记录请求模型
type RecordBestSingleReq struct {
	UserId        int64            `json:"userId" gorm:"primaryKey"` // 用户ID
	Dimension     int              `json:"dimension"`                // 阶数 3 | 4 | 5 | 6 | 7 | 8
	RecordId      int64            `json:"recordId"`                 // 记录ID
	DurationRange []int            `json:"duration_range"`           // 耗时范围
	StepRange     []int            `json:"stepRange"`                // 步数范围
	DateRange     []time.Time      `json:"dateRange"`                // 日期范围
	Pagination    utils.Pagination `gorm:"embedded"`                 // 分页
	Sorted        string           `json:"sorted"`                   // 排序
}

// RecordBestSingleResp 最佳单次记录响应模型
type RecordBestSingleResp struct {
	UserId         string    `json:"userId" gorm:"primaryKey"`        // 用户ID
	UserInfo       UserResp  `json:"userInfo" gorm:"embedded"`        // 用户信息
	Dimension      int       `json:"dimension"`                       // 阶数 3 | 4 | 5 | 6 | 7 | 8
	RecordId       string    `json:"recordId"`                        // 记录ID
	RecordDuration int       `json:"recordDuration"`                  // 耗时
	RecordStep     int       `json:"recordStep"`                      // 步数
	CreatedAt      time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt      time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

// RecordBestSingleListReq 最佳单次记录列表请求模型
type RecordBestSingleListResp struct {
	Total   int64                  `json:"total"`
	Records []RecordBestSingleResp `json:"records"`
}
