package models

import (
	"puzzle/utils"
	"time"
)

// RecordBestAverage 最佳平均记录模型
type RecordBestAverage struct {
	Id                    int64     `json:"id" gorm:"primaryKey"`            // 主键ID
	UserId                int64     `json:"userId"`                          // 用户ID
	Dimension             int       `json:"dimension"`                       // 阶数 3 | 4 | 5 | 6 | 7 | 8
	Type                  int       `json:"type"`                            // 类型 5:5次平均 12:12次平均
	RecordIds             string    `json:"recordIds"`                       // 记录ID
	RecordAverageDuration int       `json:"recordAverageDuration"`           // 平均耗时
	RecordBreakCount      int       `json:"recordBreakCount"`                // 破纪录次数
	CreatedAt             time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt             time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

// RecordBestAverageReq 最佳平均记录请求模型
type RecordBestAverageReq struct {
	Id               int64 `json:"-"`                // 主键ID
	UserId           int64 `json:"-"`                // 用户ID
	Dimension        int   `json:"dimension"`        // 阶数 3 | 4 | 5 | 6 | 7 | 8
	Type             int   `json:"type"`             // 类型 5:5次平均 12:12次平均
	RecordBreakCount int   `json:"recordBreakCount"` // 破纪录次数

	IdStr            string           `json:"id"`               // 主键ID
	UserIdStr        string           `json:"userId"`           // 用户ID
	Username         string           `json:"username"`         // 用户名
	Nickname         string           `json:"nickname"`         // 昵称
	DurationRange    []int            `json:"durationRange"`    // 耗时范围
	DateRange        []time.Time      `json:"dateRange"`        // 日期范围
	Pagination       utils.Pagination `gorm:"embedded"`         // 分页
	Sorted           string           `json:"sorted"`           // 排序
	OrderBy          string           `json:"orderBy"`          // 排序字段
	NeedUserInfo     bool             `json:"needUserInfo"`     // 是否需要用户信息
	NeedRecordDetail bool             `json:"needRecordDetail"` // 是否需要记录详情
}

// RecordBestAverageResp 最佳平均记录响应模型
type RecordBestAverageResp struct {
	Id     string `json:"id" gorm:"primaryKey"` // 主键ID
	UserId string `json:"userId"`               // 用户ID

	Dimension             int            `json:"dimension"`                       // 阶数 3 | 4 | 5 | 6 | 7 | 8
	Type                  int            `json:"type"`                            // 类型 5:5次平均 12:12次平均
	RecordBreakCount      int            `json:"recordBreakCount"`                // 破纪录次数
	RecordIds             string         `json:"recordIds"`                       // 记录ID
	RecordAverageDuration int            `json:"recordAverageDuration"`           // 平均耗时
	CreatedAt             time.Time      `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt             time.Time      `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
	UserInfo              UserResp       `json:"userInfo" gorm:"-"`               // 用户信息
	RecordDetail          RecordListResp `json:"recordDetail" gorm:"-"`           // 记录详情
}

// RecordBestAverageListResp 最佳平均记录列表响应模型
type RecordBestAverageListResp struct {
	Total   int64                   `json:"total"`
	Records []RecordBestAverageResp `json:"records"`
}
