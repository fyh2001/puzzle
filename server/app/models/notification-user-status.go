package models

import (
	"puzzle/utils"
	"time"
)

type NotificationUserStatus struct {
	Id             int64     `json:"id" gorm:"primary_key"`           // 主键ID
	UserId         int64     `json:"userId"`                          // 用户ID
	NotificationId int64     `json:"notificationId"`                  // 通知ID
	Status         int       `json:"status"`                          // 状态 1:已读 0:未读
	CreatedAt      time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt      time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

type NotificationUserStatusReq struct {
	Id                int64            `json:"-"`              // 主键ID
	Ids               []int64          `json:"-"`              // 主键ID集合
	UserId            int64            `json:"-"`              // 用户ID
	NotificationId    int64            `json:"-"`              // 通知ID
	Status            int              `json:"status"`         // 状态 1:已读 2:未读
	IdStr             string           `json:"id"`             // 主键ID
	IdsStr            string           `json:"ids"`            // 主键ID集合
	NotificationIdStr string           `json:"notificationId"` // 通知ID
	UserIdStr         string           `json:"userId"`         // 用户ID
	Username          string           `json:"username"`       // 用户名
	Nickname          string           `json:"nickname"`       // 昵称
	Sorted            string           `json:"sorted"`         // 排序
	OrderBy           string           `json:"orderBy"`        // 排序字段
	Pagination        utils.Pagination `gorm:"embedded"`       // 分页
	DateRange         []time.Time      `json:"dateRange"`      // 时间范围
}

type NotificationUserStatusResp struct {
	Id             string    `json:"id"`             // 主键ID
	UserId         string    `json:"userId"`         // 用户ID
	NotificationId string    `json:"notificationId"` // 通知ID
	Status         int       `json:"status"`         // 状态 1:已读 2:未读
	CreatedAt      time.Time `json:"createdAt" `     // 创建时间
	UpdatedAt      time.Time `json:"updatedAt" `     // 更新时间
}

type NotificationUserStatusListResp struct {
	Total   int64                    `json:"total"`   // 总数
	Records []NotificationUserStatus `json:"records"` // 用户通知状态列表
}

func (NotificationUserStatusResp) TableName() string {
	return "notification_user_status"
}
