package models

import (
	"gorm.io/gorm"
	"puzzle/app/types"
	"puzzle/utils"
	"time"
)

// Record 记录模型
type Record struct {
	ID        int64              `json:"id" gorm:"primaryKey"`            // 主键ID
	UserID    int64              `json:"userId" check:"true"`             // 用户ID
	Dimension int                `json:"dimension" check:"true"`          // 阶数 3 | 4 | 5 | 6 | 7 | 8
	Type      types.RecordType   `json:"type" check:"true"`               // 类型 1:练习 2:排行榜 3:对战
	Duration  int                `json:"duration" check:"true"`           // 耗时
	Step      int                `json:"step" check:"true"`               // 步数
	Status    types.RecordStatus `json:"status"`                          // 状态 1:启用 2:冻结 3:删除
	Scramble  string             `json:"scramble" check:"true"`           // 打乱公式
	Solution  string             `json:"solution" check:"true"`           // 解法
	IDx       int64              `json:"idx" check:"true"`                // 打乱随机数
	DeletedAt gorm.DeletedAt     `json:"-"`                               // 删除标记
	CreatedAt time.Time          `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time          `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

// RecordReq 记录请求模型
type RecordReq struct {
	ID            int64              `json:"-"`             // 主键ID
	IDs           []int64            `json:"-"`             // 主键ID列表
	UserID        int64              `json:"-"`             // 用户ID
	Dimension     int                `json:"dimension"`     // 阶数 3 | 4 | 5 | 6 | 7 | 8
	Type          types.RecordType   `json:"type"`          // 类型 1:练习 2:排行榜 3:对战
	Duration      int                `json:"duration"`      // 耗时
	Step          int                `json:"step"`          // 步数
	Status        types.RecordStatus `json:"status"`        // 状态 1:启用 2:冻结 3:删除
	Scramble      string             `json:"scramble"`      // 打乱公式
	Solution      string             `json:"solution"`      // 解法
	IDx           int64              `json:"-"`             // 打乱随机数
	Pagination    utils.Pagination   `gorm:"embedded"`      // 分页
	DurationRange []int              `json:"durationRange"` // 耗时范围
	StepRange     []int              `json:"stepRange"`     // 步数范围
	DateRange     []time.Time        `json:"dateRange"`     // 日期范围
	IDStr         string             `json:"id"`            // 主键ID
	IDsStr        []string           `json:"ids"`           // 主键ID列表
	UserIDStr     string             `json:"userId"`        // 用户ID
	Username      string             `json:"username"`      // 用户名
	Nickname      string             `json:"nickname"`      // 昵称
	IDxStr        string             `json:"idx"`           // 打乱随机数
	Sorted        string             `json:"sorted"`        // 排序
	OrderBy       string             `json:"orderBy"`       // 排序字段
	NeedUserData  bool               `json:"needUserData"`  // 是否需要用户信息
}

// RecordResp 记录响应模型
type RecordResp struct {
	ID        string             `json:"id"`                                              // 主键ID
	UserID    string             `json:"userId"`                                          // 用户ID
	Dimension int                `json:"dimension"`                                       // 阶数 3 | 4 | 5 | 6 | 7 | 8
	Type      types.RecordType   `json:"type"`                                            // 类型 1:练习 2:排行榜 3:对战
	Duration  int                `json:"duration"`                                        // 耗时
	Step      int                `json:"step"`                                            // 步数
	Status    types.RecordStatus `json:"status"`                                          // 状态 1:启用 2:冻结 3:删除
	Scramble  string             `json:"scramble"`                                        // 打乱公式
	Solution  string             `json:"solution"`                                        // 解法
	IDx       string             `json:"idx"`                                             // 打乱随机数
	CreatedAt time.Time          `json:"createdAt"`                                       // 创建时间
	UpdatedAt time.Time          `json:"updatedAt"`                                       // 更新时间
	UserData  UserResp           `json:"userData" gorm:"foreignKey:ID;references:UserID"` // 用户信息
}

// RecordListResp 记录列表响应模型
type RecordListResp struct {
	Total   int64        `json:"total"`
	Records []RecordResp `json:"records"`
}

func (recordResp RecordResp) TableName() string {
	return "record"
}

func (r *Record) BeforeCreate(tx *gorm.DB) (err error) {

	r.ID = (&utils.Snowflake{}).NextVal()
	r.Status = types.RecordStatusNormal

	return
}
