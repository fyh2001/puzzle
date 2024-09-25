package models

import (
	"gorm.io/gorm"
	"puzzle/utils"
	"time"
)

type Notification struct {
	ID         int64  `json:"id" gorm:"primary_key"`       // 主键ID
	UserID     int64  `json:"userId" gorm:"default 0"`     // 用户ID 0:全体通知
	TypeID     int    `json:"typeId"`                      // 类型ID
	Content    string `json:"content"`                     // 内容
	ReadStatus int    `json:"readStatus" gorm:"default 1"` // 状态 1:未读 2: 已读
	Status     int    `json:"status" gorm:"default 1"`     // 状态 1:启用 2:冻结 3:删除
	DeletedAt  gorm.DeletedAt
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

type NotificationReq struct {
	ID         int64            `json:"-"`                           // 主键ID
	IDs        []int64          `json:"-"`                           // 主键ID集合
	UserID     int64            `json:"-"`                           // 用户ID
	TypeID     int              `json:"typeId"`                      // 类型ID
	Content    string           `json:"content"`                     // 内容
	ReadStatus int              `json:"readStatus" gorm:"default 1"` // 状态 1:未读 2: 已读
	Status     int              `json:"status" gorm:"default 1"`     // 状态 1:启用 2:冻结 3:删除
	IDStr      string           `json:"id"`                          // 主键ID
	IDsStr     []string         `json:"ids"`                         // 主键ID集合
	UserIDStr  string           `json:"userId"`                      // 用户ID
	Username   string           `json:"username"`                    // 用户名
	Nickname   string           `json:"nickname"`                    // 昵称
	Sorted     string           `json:"sorted"`                      // 排序
	OrderBy    string           `json:"orderBy"`                     // 排序字段
	OnlyTotal  bool             `json:"onlyTotal"`                   // 只查询总数
	Pagination utils.Pagination `gorm:"embedded"`                    // 分页
	DateRange  []time.Time      `json:"dateRange"`                   // 时间范围
}

type NotificationResp struct {
	ID                   string               `json:"id"`                                                          // 主键ID
	UserID               string               `json:"userId"`                                                      // 用户ID
	TypeID               int                  `json:"typeId"`                                                      // 类型ID
	Content              string               `json:"content"`                                                     // 内容
	ReadStatus           int                  `json:"readStatus" gorm:"default 1"`                                 // 状态 1:未读 2: 已读
	Status               int                  `json:"status"`                                                      // 状态 1:已读 2:未读
	CreatedAt            time.Time            `json:"createdAt"`                                                   // 创建时间
	UpdatedAt            time.Time            `json:"updatedAt"`                                                   // 更新时间
	NotificationTypeInfo NotificationTypeResp `json:"notificationTypeInfo" gorm:"foreignKey:Id;references:TypeId"` // 通知类型信息
}

type NotificationListResp struct {
	Total   int64              `json:"total"`   // 总数
	Records []NotificationResp `json:"records"` // 通知列表
}
