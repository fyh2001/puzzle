package models

import (
	"puzzle/utils"
	"time"
)

// RecordBestAverage 最佳平均记录模型
type RecordBestAverage struct {
	UserId                int64     `json:"userId"`                          // 用户ID
	Dimension             int       `json:"dimension"`                       // 阶数 3 | 4 | 5 | 6 | 7 | 8
	Type                  int       `json:"type"`                            // 类型 5:5次平均 12:12次平均
	RecordIds             string    `json:"recordIds"`                       // 记录ID
	RecordAverageDuration int       `json:"recordAverageDuration"`           // 平均耗时
	CreatedAt             time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt             time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

// RecordBestAverageReq 最佳平均记录请求模型
type RecordBestAverageReq struct {
	UserId        int64            `json:"userId"`        // 用户ID
	Dimension     int              `json:"dimension"`     // 阶数 3 | 4 | 5 | 6 | 7 | 8
	Type          int              `json:"type"`          // 类型 5:5次平均 12:12次平均
	DurationRange []int            `json:"durationRange"` // 耗时范围
	DateRange     []time.Time      `json:"dateRange"`     // 日期范围
	Pagination    utils.Pagination `gorm:"embedded"`      // 分页
	Sorted        string           `json:"sorted"`        // 排序
}

// RecordBestAverageResp 最佳平均记录响应模型
type RecordBestAverageResp struct {
	UserId                string    `json:"userId"`                          // 用户ID
	UserInfo              UserResp  `json:"userInfo" gorm:"embedded"`        // 用户信息
	Dimension             int       `json:"dimension"`                       // 阶数 3 | 4 | 5 | 6 | 7 | 8
	Type                  int       `json:"type"`                            // 类型 5:5次平均 12:12次平均
	RecordIds             string    `json:"recordIds"`                       // 记录ID
	RecordAverageDuration int       `json:"recordAverageDuration"`           // 平均耗时
	CreatedAt             time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt             time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

// RecordBestAverageListResp 最佳平均记录列表响应模型
type RecordBestAverageListResp struct {
	Total   int64                   `json:"total"`
	Records []RecordBestAverageResp `json:"records"`
}
