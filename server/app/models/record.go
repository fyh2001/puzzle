package models

import (
	"puzzle/utils"
	"time"
)

// Record 记录模型
type Record struct {
	Id        int64     `json:"id" gorm:"primaryKey"`            // 主键ID
	UserId    int64     `json:"userId"`                          // 用户ID
	Dimension int       `json:"dimension"`                       // 阶数 3 | 4 | 5 | 6 | 7 | 8
	Type      int       `json:"type"`                            // 类型 1:练习 2:排行榜 3:对战
	Duration  int       `json:"duration"`                        // 耗时
	Step      int       `json:"step"`                            // 步数
	Status    int       `json:"status"`                          // 状态 1:启用 2:冻结 3:删除
	Scramble  string    `json:"scramble"`                        // 打乱公式
	Solution  string    `json:"solution"`                        // 解法
	Idx       int64     `json:"idx"`                             // 打乱随机数
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

// RecordReq 记录请求模型
type RecordReq struct {
	Id        int64   `json:"-"`         // 主键ID
	Ids       []int64 `json:"-"`         // 主键ID列表
	UserId    int64   `json:"-"`         // 用户ID
	Dimension int     `json:"dimension"` // 阶数 3 | 4 | 5 | 6 | 7 | 8
	Type      int     `json:"type"`      // 类型 1:练习 2:排行榜 3:对战
	Duration  int     `json:"duration"`  // 耗时
	Step      int     `json:"step"`      // 步数
	Status    int     `json:"status"`    // 状态 1:启用 2:冻结 3:删除
	Scramble  string  `json:"scramble"`  // 打乱公式
	Solution  string  `json:"solution"`  // 解法
	Idx       int64   `json:"-"`         // 打乱随机数

	Pagination    utils.Pagination `gorm:"embedded"`      // 分页
	DurationRange []int            `json:"durationRange"` // 耗时范围
	StepRange     []int            `json:"stepRange"`     // 步数范围
	DateRange     []time.Time      `json:"dateRange"`     // 日期范围
	IdStr         string           `json:"id"`            // 主键ID
	IdsStr        []string         `json:"ids"`           // 主键ID列表
	UserIdStr     string           `json:"userId"`        // 用户ID
	Username      string           `json:"username"`      // 用户名
	Nickname      string           `json:"nickname"`      // 昵称
	IdxStr        string           `json:"idx"`           // 打乱随机数
	Sorted        string           `json:"sorted"`        // 排序
	OrderBy       string           `json:"orderBy"`       // 排序字段
	NeedUserInfo  bool             `json:"needUserInfo"`  // 是否需要用户信息
}

// RecordResp 记录响应模型
type RecordResp struct {
	Id        string    `json:"id"`                                              // 主键ID
	UserId    string    `json:"userId"`                                          // 用户ID
	UserInfo  UserResp  `json:"userInfo" gorm:"foreignKey:UserId;references:Id"` // 用户信息
	Dimension int       `json:"dimension"`                                       // 阶数 3 | 4 | 5 | 6 | 7 | 8
	Type      int       `json:"type"`                                            // 类型 1:练习 2:排行榜 3:对战
	Duration  int       `json:"duration"`                                        // 耗时
	Step      int       `json:"step"`                                            // 步数
	Status    int       `json:"status"`                                          // 状态 1:启用 2:冻结 3:删除
	Scramble  string    `json:"scramble"`                                        // 打乱公式
	Solution  string    `json:"solution"`                                        // 解法
	Idx       string    `json:"idx"`                                             // 打乱随机数
	CreatedAt time.Time `json:"createdAt"`                                       // 创建时间
	UpdatedAt time.Time `json:"updatedAt"`                                       // 更新时间
}

// RecordListResp 记录列表响应模型
type RecordListResp struct {
	Total   int64        `json:"total"`
	Records []RecordResp `json:"records"`
}
