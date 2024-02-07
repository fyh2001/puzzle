package models

import (
	"puzzle/utils"
	"time"
)

// Record 记录模型
type Record struct {
	Id        int64     `json:"id" gorm:"primaryKey"`             // 主键ID
	UserId    int64     `json:"user_id"`                          // 用户ID
	Dimension int       `json:"dimension"`                        // 阶数 3 | 4 | 5 | 6 | 7 | 8
	Type      int       `json:"type"`                             // 类型 0:练习 1:排行榜 2:对战
	Duration  int       `json:"duration"`                         // 耗时
	Step      int       `json:"step"`                             // 步数
	Status    int       `json:"status"`                           // 状态 0:启用 1:冻结 2:删除
	Scramble  string    `json:"scramble"`                         // 打乱公式
	Solution  string    `json:"solution"`                         // 解法
	Idx       int64     `json:"idx"`                              // 打乱随机数
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"` // 更新时间
}

// RecordReq 记录请求模型
type RecordReq struct {
	Id            int64            `json:"id"`             // 主键ID
	UserId        int64            `json:"user_id"`        // 用户ID
	Dimension     int              `json:"dimension"`      // 阶数 3 | 4 | 5 | 6 | 7 | 8
	Type          int              `json:"type"`           // 类型 0:个人 1:排行榜 2:对战
	Status        int              `json:"status"`         // 状态 0:启用 1:冻结 2:删除
	DurationRange []int            `json:"duration_range"` // 耗时范围
	StepRange     []int            `json:"step_range"`     // 步数范围
	DateRange     []time.Time      `json:"date_range"`     // 日期范围
	Pagination    utils.Pagination `gorm:"embedded"`       // 分页
}

// RecordResp 记录响应模型
type RecordResp struct {
	Id        int64     `json:"id"`         // 主键ID
	UserId    int64     `json:"user_id"`    // 用户ID
	Dimension int       `json:"dimension"`  // 阶数 3 | 4 | 5 | 6 | 7 | 8
	Type      int       `json:"type"`       // 类型 0:个人 1:排行榜 2:对战
	Duration  int       `json:"duration"`   // 耗时
	Step      int       `json:"step"`       // 步数
	Status    int       `json:"status"`     // 状态 0:启用 1:冻结 2:删除
	Scramble  string    `json:"scramble"`   // 打乱公式
	Solution  string    `json:"solution"`   // 解法
	Idx       int64     `json:"idx"`        // 打乱随机数
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}

// RecordListResp 记录列表响应模型
type RecordListResp struct {
	Total   int64        `json:"total"`
	Records []RecordResp `json:"records"`
}
