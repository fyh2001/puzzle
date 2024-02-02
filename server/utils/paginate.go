package utils

import "gorm.io/gorm"

// 分页模型
type Pagination struct {
	Page     int `form:"page" json:"page"`
	PageSize int `form:"pageSize" json:"pageSize"`
	Type     int `form:"type" json:"type"` // 1: Record 2: BestRecord
}

// Paginate 分页封装
func Paginate(pagination *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := pagination.Page
		if page == 0 {
			page = 1
		}
		pageSize := pagination.PageSize
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize

		if pagination.Type == 1 {
			pageSize += 11
		}

		return db.Offset(offset).Limit(pageSize)
	}
}
