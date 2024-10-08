package models

import (
	"gorm.io/gorm"
	"puzzle/utils"
	"time"
)

// RecordBestAverage 最佳平均记录模型
type RecordBestAverage struct {
	ID                    int64  `json:"id" gorm:"primaryKey"`  // 主键ID
	UserID                int64  `json:"userId"`                // 用户ID
	Dimension             int    `json:"dimension"`             // 阶数 3 | 4 | 5 | 6 | 7 | 8
	Type                  int    `json:"type"`                  // 类型 5:5次平均 12:12次平均
	RecordIDs             string `json:"recordIds"`             // 记录ID
	RecordAverageDuration int    `json:"recordAverageDuration"` // 平均耗时
	RecordBreakCount      int    `json:"recordBreakCount"`      // 破纪录次数
	Ranked                int    `json:"ranked"`                // 排名
	DeletedAt             gorm.DeletedAt
	CreatedAt             time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt             time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

// RecordBestAverageReq 最佳平均记录请求模型
type RecordBestAverageReq struct {
	ID               int64 `json:"-"`                // 主键ID
	UserID           int64 `json:"-"`                // 用户ID
	Dimension        int   `json:"dimension"`        // 阶数 3 | 4 | 5 | 6 | 7 | 8
	Type             int   `json:"type"`             // 类型 5:5次平均 12:12次平均
	RecordBreakCount int   `json:"recordBreakCount"` // 破纪录次数

	IDStr            string           `json:"id"`               // 主键ID
	UserIDStr        string           `json:"userId"`           // 用户ID
	Username         string           `json:"username"`         // 用户名
	Nickname         string           `json:"nickname"`         // 昵称
	DurationRange    []int            `json:"durationRange"`    // 耗时范围
	DateRange        []time.Time      `json:"dateRange"`        // 日期范围
	RankRange        []int            `json:"rankRange"`        // 排名范围
	BreakCountRange  []int            `json:"breakCountRange"`  // 破纪录次数范围
	Pagination       utils.Pagination `gorm:"embedded"`         // 分页
	Sorted           string           `json:"sorted"`           // 排序
	OrderBy          string           `json:"orderBy"`          // 排序字段
	NeedUserInfo     bool             `json:"needUserInfo"`     // 是否需要用户信息
	NeedRecordDetail bool             `json:"needRecordDetail"` // 是否需要记录详情
}

// RecordBestAverageResp 最佳平均记录响应模型
type RecordBestAverageResp struct {
	ID                    string         `json:"id" gorm:"primaryKey"`                            // 主键ID
	UserID                string         `json:"userId"`                                          // 用户ID
	Dimension             int            `json:"dimension"`                                       // 阶数 3 | 4 | 5 | 6 | 7 | 8
	Type                  int            `json:"type"`                                            // 类型 5:5次平均 12:12次平均
	RecordBreakCount      int            `json:"recordBreakCount"`                                // 破纪录次数
	RecordIDs             string         `json:"recordIds"`                                       // 记录ID
	RecordAverageDuration int            `json:"recordAverageDuration"`                           // 平均耗时
	Ranked                int            `json:"ranked"`                                          // 排名
	CreatedAt             time.Time      `json:"createdAt"`                                       // 创建时间
	UpdatedAt             time.Time      `json:"updatedAt"`                                       // 更新时间
	UserInfo              UserResp       `json:"userInfo" gorm:"foreignKey:ID;references:UserID"` // 用户信息
	RecordDetail          RecordListResp `json:"recordDetail" gorm:"-"`                           // 记录详情
}

// RecordBestAverageListResp 最佳平均记录列表响应模型
type RecordBestAverageListResp struct {
	Total   int64                   `json:"total"`
	Records []RecordBestAverageResp `json:"records"`
}
