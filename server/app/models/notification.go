package models

import (
	"puzzle/utils"
	"time"
)

type Notification struct {
	Id        int64     `json:"id" gorm:"primary_key"`           // 主键ID
	UserId    int64     `json:"userId" gorm:"default 0"`         // 用户ID 0:全体通知
	TypeId    int       `json:"typeId"`                          // 类型ID
	Content   string    `json:"content"`                         // 内容
	Status    int       `json:"status" gorm:"default 1"`         // 状态 1:启用 2:冻结 3:删除
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

type NotificationReq struct {
	Id         int64            `json:"-"`                       // 主键ID
	Ids        []int64          `json:"-"`                       // 主键ID集合
	UserId     int64            `json:"-"`                       // 用户ID
	TypeId     int              `json:"typeId"`                  // 类型ID
	Content    string           `json:"content"`                 // 内容
	Status     int              `json:"status" gorm:"default 1"` // 状态 1:启用 2:冻结 3:删除
	IdStr      string           `json:"id"`                      // 主键ID
	IdsStr     []string         `json:"ids"`                     // 主键ID集合
	UserIdStr  string           `json:"userId"`                  // 用户ID
	Username   string           `json:"username"`                // 用户名
	Nickname   string           `json:"nickname"`                // 昵称
	Sorted     string           `json:"sorted"`                  // 排序
	OrderBy    string           `json:"orderBy"`                 // 排序字段
	Pagination utils.Pagination `gorm:"embedded"`                // 分页
	DateRange  []time.Time      `json:"dateRange"`               // 时间范围
}

type NotificationResp struct {
	Id                         string                     `json:"id"`                                                                        // 主键ID
	UserId                     string                     `json:"userId"`                                                                    // 用户ID
	TypeId                     int                        `json:"typeId"`                                                                    // 类型ID
	Content                    string                     `json:"content"`                                                                   // 内容
	Status                     int                        `json:"status"`                                                                    // 状态 1:已读 2:未读
	CreatedAt                  time.Time                  `json:"createdAt"`                                                                 // 创建时间
	UpdatedAt                  time.Time                  `json:"updatedAt"`                                                                 // 更新时间
	NotificationUserStatusInfo NotificationUserStatusResp `json:"notificationUserStatusInfo" gorm:"foreignKey:NotificationId;references:Id"` // 通知用户状态信息
}

type NotificationListResp struct {
	Total   int64              `json:"total"`   // 总数
	Records []NotificationResp `json:"records"` // 通知列表
}
