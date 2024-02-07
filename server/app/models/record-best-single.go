package models

import (
	"puzzle/utils"
	"time"
)

// RecordBestSingle 最佳单次记录模型
type RecordBestSingle struct {
	UserId         int64     `json:"user_id" gorm:"primaryKey"`        // 用户ID
	Dimension      int       `json:"dimension"`                        // 阶数 3 | 4 | 5 | 6 | 7 | 8
	RecordId       int64     `json:"record_id"`                        // 记录ID
	RecordDuration int       `json:"record_duration"`                  // 耗时
	RecordStep     int       `json:"record_step"`                      // 步数
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"` // 更新时间
}

// RecordBestSingleReq 最佳单次记录请求模型
type RecordBestSingleReq struct {
	UserId        int64            `json:"user_id" gorm:"primaryKey"` // 用户ID
	Dimension     int              `json:"dimension"`                 // 阶数 3 | 4 | 5 | 6 | 7 | 8
	RecordId      int64            `json:"record_id"`                 // 记录ID
	DurationRange []int            `json:"duration_range"`            // 耗时范围
	StepRange     []int            `json:"step_range"`                // 步数范围
	DateRange     []time.Time      `json:"date_range"`                // 日期范围
	Pagination    utils.Pagination `gorm:"embedded"`                  // 分页
}

// RecordBestSingleResp 最佳单次记录响应模型
type RecordBestSingleResp struct {
	UserId         int64     `json:"user_id" gorm:"primaryKey"`        // 用户ID
	Dimension      int       `json:"dimension"`                        // 阶数 3 | 4 | 5 | 6 | 7 | 8
	RecordId       int64     `json:"record_id"`                        // 记录ID
	RecordDuration int       `json:"record_duration"`                  // 耗时
	RecordStep     int       `json:"record_step"`                      // 步数
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"` // 更新时间
}

// RecordBestSingleListReq 最佳单次记录列表请求模型
type RecordBestSingleListResp struct {
	Total   int64                  `json:"total"`
	Records []RecordBestSingleResp `json:"records"`
}
