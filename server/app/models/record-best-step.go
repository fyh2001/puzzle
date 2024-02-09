package models

import (
	"puzzle/utils"
	"time"
)

// RecordBestStep 用户最佳步数记录模型
type RecordBestStep struct {
	UserId     int64     `json:"userId"`                          // 用户ID
	Dimension  int       `json:"dimension"`                       // 阶数 3 | 4 | 5 | 6 | 7 | 8
	RecordId   int64     `json:"recordId"`                        // 记录ID
	RecordStep int       `json:"recordStep"`                      // 步数
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

// RecordBestStepReq 用户最佳步数记录请求模型
type RecordBestStepReq struct {
	UserId     int64            `json:"userId"`    // 用户ID
	Dimension  int              `json:"dimension"` // 阶数 3 | 4 | 5 | 6 | 7 | 8
	RecordId   int64            `json:"recordId"`  // 记录ID
	StepRange  []int            `json:"stepRange"` // 步数范围
	DateRange  []time.Time      `json:"dateRange"` // 日期范围
	Pagination utils.Pagination `gorm:"embedded"`  // 分页
	Sorted     string           `json:"sorted"`    // 排序
}

// RecordBestStepResp 用户最佳步数记录响应模型
type RecordBestStepResp struct {
	UserId     string    `json:"userId"`                          // 用户ID
	UserInfo   UserResp  `json:"userInfo" gorm:"embedded"`        // 用户信息
	Dimension  int       `json:"dimension"`                       // 阶数 3 | 4 | 5 | 6 | 7 | 8
	RecordId   string    `json:"recordId"`                        // 记录ID
	RecordStep int       `json:"recordStep"`                      // 步数
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

// RecordBestStepListResp 用户最佳步数记录列表响应模型
type RecordBestStepListResp struct {
	Total   int64                `json:"total"`
	Records []RecordBestStepResp `json:"records"`
}
