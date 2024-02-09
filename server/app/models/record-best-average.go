package models

import (
	"puzzle/utils"
	"time"
)

// RecordBestAverage 最佳平均记录模型
type RecordBestAverage struct {
	UserId                int64     `json:"user_id" gorm:"primaryKey"`        // 用户ID
	Dimension             int       `json:"dimension"`                        // 阶数 3 | 4 | 5 | 6 | 7 | 8
	Type                  int       `json:"type"`                             // 类型 5:5次平均 12:12次平均
	RecordIds             string    `json:"record_ids"`                       // 记录ID
	RecordAverageDuration int       `json:"record_average_duration"`          // 平均耗时
	CreatedAt             time.Time `json:"created_at" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt             time.Time `json:"updated_at" gorm:"autoUpdateTime"` // 更新时间
}

// RecordBestAverageReq 最佳平均记录请求模型
type RecordBestAverageReq struct {
	UserId        int64            `json:"user_id" gorm:"primaryKey"` // 用户ID
	Dimension     int              `json:"dimension"`                 // 阶数 3 | 4 | 5 | 6 | 7 | 8
	Type          int              `json:"type"`                      // 类型 5:5次平均 12:12次平均
	DurationRange []int            `json:"duration_range"`            // 耗时范围
	DateRange     []time.Time      `json:"date_range"`                // 日期范围
	Pagination    utils.Pagination `gorm:"embedded"`                  // 分页
	Sorted        string           `json:"sorted"`                    // 排序
}

// RecordBestAverageResp 最佳平均记录响应模型
type RecordBestAverageResp struct {
	UserId                int64     `json:"user_id" gorm:"primaryKey"`        // 用户ID
	Dimension             int       `json:"dimension"`                        // 阶数 3 | 4 | 5 | 6 | 7 | 8
	Type                  int       `json:"type"`                             // 类型 5:5次平均 12:12次平均
	RecordIds             string    `json:"record_ids"`                       // 记录ID
	RecordAverageDuration int       `json:"record_average_duration"`          // 平均耗时
	CreatedAt             time.Time `json:"created_at" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt             time.Time `json:"updated_at" gorm:"autoUpdateTime"` // 更新时间
}

// RecordBestAverageListResp 最佳平均记录列表响应模型
type RecordBestAverageListResp struct {
	Total   int64                   `json:"total"`
	Records []RecordBestAverageResp `json:"records"`
}
