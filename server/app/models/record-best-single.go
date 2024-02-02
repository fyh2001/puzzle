package models

import "time"

type RecordBestSingle struct {
	UserId         int64     `json:"user_id" gorm:"primaryKey"`        // 用户ID
	Event          string    `json:"event"`                            // p15 | p24 | p35 | p48 | p63 | p80
	RecordId       string    `json:"record_id"`                        // 记录ID
	RecordDuration int       `json:"record_duration"`                  // 耗时
	RecordStep     int       `json:"record_step"`                      // 步数
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"` // 更新时间
}
