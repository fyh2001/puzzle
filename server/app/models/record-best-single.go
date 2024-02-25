package models

import (
	"puzzle/utils"
	"time"
)

// RecordBestSingle 最佳单次记录模型
type RecordBestSingle struct {
	Id               int64     `json:"id" gorm:"primaryKey"`            // 主键ID
	UserId           int64     `json:"userId" gorm:"primaryKey"`        // 用户ID
	Dimension        int       `json:"dimension"`                       // 阶数 3 | 4 | 5 | 6 | 7 | 8
	RecordId         int64     `json:"recordId"`                        // 记录ID
	RecordDuration   int       `json:"recordDuration"`                  // 耗时
	RecordStep       int       `json:"recordStep"`                      // 步数
	RecordBreakCount int       `json:"recordBreakCount"`                // 破纪录次数
	Rank             int       `json:"rank"`                            // 排名
	CreatedAt        time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt        time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

// RecordBestSingleReq 最佳单次记录请求模型
type RecordBestSingleReq struct {
	Id               int64 `json:"-"`                // 主键ID
	UserId           int64 `json:"-"`                // 用户ID
	Dimension        int   `json:"dimension"`        // 阶数 3 | 4 | 5 | 6 | 7 | 8
	RecordId         int64 `json:"-"`                // 记录ID
	RecordBreakCount int   `json:"recordBreakCount"` // 破纪录次数

	IdStr            string           `json:"id"`               // 主键ID
	UserIdStr        string           `json:"userId"`           // 用户ID
	Username         string           `json:"username"`         // 用户名
	Nickname         string           `json:"nickname"`         // 昵称
	RecordIdStr      string           `json:"recordId"`         // 记录ID
	DurationRange    []int            `json:"durationRange"`    // 耗时范围
	StepRange        []int            `json:"stepRange"`        // 步数范围
	DateRange        []time.Time      `json:"dateRange"`        // 日期范围
	RankRange        []int            `json:"rankRange"`        // 排名范围
	Pagination       utils.Pagination `gorm:"embedded"`         // 分页
	Sorted           string           `json:"sorted"`           // 排序
	OrderBy          string           `json:"orderBy"`          // 排序字段
	NeedUserInfo     bool             `json:"needUserInfo"`     // 是否需要用户信息
	NeedRecordDetail bool             `json:"needRecordDetail"` // 是否需要记录详情
}

// RecordBestSingleResp 最佳单次记录响应模型
type RecordBestSingleResp struct {
	Id               string         `json:"id" gorm:"primaryKey"`            // 主键ID
	UserId           string         `json:"userId"`                          // 用户ID
	Dimension        int            `json:"dimension"`                       // 阶数 3 | 4 | 5 | 6 | 7 | 8
	RecordId         string         `json:"recordId"`                        // 记录ID
	RecordDuration   int            `json:"recordDuration"`                  // 耗时
	RecordStep       int            `json:"recordStep"`                      // 步数
	RecordBreakCount int            `json:"recordBreakCount"`                // 破纪录次数
	Rank             int            `json:"rank"`                            // 排名
	CreatedAt        time.Time      `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt        time.Time      `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
	UserInfo         UserResp       `json:"userInfo" gorm:"-"`               // 用户信息
	RecordDetail     RecordListResp `json:"recordDetail" gorm:"-"`           // 记录详情
}

// RecordBestSingleListReq 最佳单次记录列表请求模型
type RecordBestSingleListResp struct {
	Total   int64                  `json:"total"`
	Records []RecordBestSingleResp `json:"records"`
}
