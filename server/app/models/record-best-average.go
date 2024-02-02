package models

import "time"

type RecordBestAverage struct {
	UserId                int64     `json:"user_id" gorm:"primaryKey"`        // 用户ID
	Event                 string    `json:"event"`                            // p15 | p24 | p35 | p48 | p63 | p80
	Type                  int       `json:"type"`                             // 类型 5:5次平均 12:12次平均
	RecordIds             string    `json:"record_ids"`                       // 记录ID
	RecordAverageDuration int       `json:"record_average_duration"`          // 平均耗时
	CreatedAt             time.Time `json:"created_at" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt             time.Time `json:"updated_at" gorm:"autoUpdateTime"` // 更新时间
}
