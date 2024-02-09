package models

import (
	"puzzle/utils"
	"time"
)

type Scramble struct {
	Id        int64     `json:"id" gorm:"primaryKey"`            // 主键ID
	Dimension int       `json:"dimension"`                       // 阶数 3 | 4 | 5 | 6
	Idx       int64     `json:"idx"`                             // 随机数索引
	Scramble  string    `json:"scramble"`                        // 打乱公式
	Status    int       `json:"status" gorm:"default 0"`         // 状态 0:启用 1:冻结 2:删除
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

type ScrambleReq struct {
	Id         int64            `json:"id"`        // 主键ID
	Dimension  int              `json:"dimension"` // 阶数 3 | 4 | 5 | 6
	Idx        int64            `json:"idx"`       // 随机数索引
	Scramble   string           `json:"scramble"`  // 打乱公式
	Status     int              `json:"status"`    // 状态 0:启用 1:冻结 2:删除
	DateRange  []time.Time      `json:"dateRange"` // 时间范围
	Pagination utils.Pagination `gorm:"embedded"`  // 分页
	Sorted     string           `json:"sorted"`    // 排序
}

type ScrambleResp struct {
	Id        int64     `json:"id" gorm:"primaryKey"`            // 主键ID
	Dimension int       `json:"dimension"`                       // 阶数 3 | 4 | 5 | 6
	Idx       int64     `json:"idx"`                             // 随机数索引
	Scramble  string    `json:"scramble"`                        // 打乱公式
	Status    int       `json:"status" gorm:"default 0"`         // 状态 0:启用 1:冻结 2:删除
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

type ScrambleListResp struct {
	Total   int64          `json:"total"`   // 总数
	Records []ScrambleResp `json:"records"` // 打乱公式列表
}

type GetNewScambleReq struct {
	UserId    int64 `json:"userId"`    // 用户ID
	Dimension int   `json:"dimension"` // 阶数 3 | 4 | 5 | 6
}
