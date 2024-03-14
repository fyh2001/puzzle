package models

import (
	"puzzle/utils"
	"time"
)

type NotificationType struct {
	Id        int       `json:"id" gorm:"primary_key"`           // 主键ID
	Name      string    `json:"name"`                            // 名称
	Icon      string    `json:"icon"`                            // 图标
	Status    int       `json:"status"`                          // 状态 1:启用 2:冻结 3:删除
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

type NotificationTypeReq struct {
	Id         int              `json:"-"`        // 主键ID
	Ids        []int            `json:"-"`        // 主键ID集合
	Name       string           `json:"name"`     // 名称
	Icon       string           `json:"icon"`     // 图标
	Status     int              `json:"status"`   // 状态 1:启用 2:冻结 3:删除
	IdStr      string           `json:"id"`       // 主键ID
	IdsStr     string           `json:"ids"`      // 主键ID集合
	Sorted     string           `json:"sorted"`   // 排序
	OrderBy    string           `json:"orderBy"`  // 排序字段
	Pagination utils.Pagination `gorm:"embedded"` // 分页
}

type NotificationTypeResp struct {
	Id        string    `json:"id"`        // 主键ID
	Name      string    `json:"name"`      // 名称
	Icon      string    `json:"icon"`      // 图标
	Status    int       `json:"status"`    // 状态 1:启用 2:冻结 3:删除
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 更新时间
}

type NotificationTypeListResp struct {
	Total   int64              `json:"total"`   // 总数
	Records []NotificationType `json:"records"` // 通知类型列表
}

func (n NotificationTypeResp) TableName() string {
	return "notification_type"
}
