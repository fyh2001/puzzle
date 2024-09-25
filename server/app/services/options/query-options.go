package options

import (
	"gorm.io/gorm"
	"puzzle/utils"
	"time"
)

type QueryOption func(*gorm.DB)

func ApplyQueryFilters(db *gorm.DB, opts ...QueryOption) {
	for _, opt := range opts {
		opt(db) // 调用每个过滤器，应用到db上
	}
}

func WithQueryID32(id int32) QueryOption {
	return func(db *gorm.DB) {
		if id != 0 {
			db.Where("id = ?", id)
		}
	}
}

func WithQueryIDs32(ids []int) QueryOption {
	return func(db *gorm.DB) {
		if len(ids) > 0 {
			db.Where("id IN (?)", ids)
		}
	}
}

func WithQueryID64(id int64) QueryOption {
	return func(db *gorm.DB) {
		if id != 0 {
			db.Where("id = ?", id)
		}
	}
}

func WithQueryIDs64(ids []int64) QueryOption {
	return func(db *gorm.DB) {
		if len(ids) > 0 {
			db.Where("id IN (?)", ids)
		}
	}
}

func WithQueryUsername(username string) QueryOption {
	return func(db *gorm.DB) {
		if username != "" {
			db.Where("username LIKE ?", "%"+username+"%")
		}
	}
}

func WithQueryNickname(nickname string) QueryOption {
	return func(db *gorm.DB) {
		if nickname != "" {
			db.Where("nickname LIKE ?", "%"+nickname+"%")
		}
	}
}

func WithQueryAccoladeID(accoladeID int) QueryOption {
	return func(db *gorm.DB) {
		if accoladeID != 0 {
			db.Where("accolade_id = ?", accoladeID)
		}
	}
}

func WithQueryEmail(email string) QueryOption {
	return func(db *gorm.DB) {
		if email != "" {
			db.Where("email = ?", email)
		}
	}
}

func WithQueryPhone(phone string) QueryOption {
	return func(db *gorm.DB) {
		if phone != "" {
			db.Where("phone = ?", phone)
		}
	}
}

func WithQueryStatus(status int) QueryOption {
	return func(db *gorm.DB) {
		if status != 0 {
			db.Where("status = ?", status)
		}
	}
}

func WithQueryDateRange(dateRange []time.Time) QueryOption {
	return func(db *gorm.DB) {
		if len(dateRange) <= 2 {
			if !dateRange[0].IsZero() {
				db.Where("created_at > ?", dateRange[0])
			}
			if !dateRange[1].IsZero() {
				db.Where("created_at < ?", dateRange[1])
			}
		}
	}
}

func WithQueryPagination(pagination utils.Pagination) QueryOption {
	return func(db *gorm.DB) {
		if pagination.Page > 0 && pagination.PageSize > 0 {
			db.Scopes(utils.Paginate(&pagination))
		}
	}
}

func WithQueryOrderBy(orderBy, sorted string) QueryOption {
	return func(db *gorm.DB) {
		db.Order(orderBy + " " + sorted)
	}
}

func WithQueryUserID(userId int64) QueryOption {
	return func(db *gorm.DB) {
		if userId != 0 {
			db.Where("user_id = ?", userId)
		}
	}
}

func WithQueryDimension(dimension int) QueryOption {
	return func(db *gorm.DB) {
		if dimension != 0 {
			db.Where("dimension = ?", dimension)
		}
	}
}

func WithQueryType(types int) QueryOption {
	return func(db *gorm.DB) {
		if types != 0 {
			db.Where("type = ?", types)
		}
	}
}

func WithQueryDurationRange(durationRange []int) QueryOption {
	return func(db *gorm.DB) {
		if len(durationRange) <= 2 {
			if durationRange[0] != 0 {
				db.Where("duration >= ?", durationRange[0])
			}
			if durationRange[1] != 0 {
				db.Where("duration <= ?", durationRange[1])
			}
		}
	}
}

func WithQueryStepRange(stepRange []int) QueryOption {
	return func(db *gorm.DB) {
		if len(stepRange) <= 2 {
			if stepRange[0] != 0 {
				db.Where("step >= ?", stepRange[0])
			}
			if stepRange[1] != 0 {
				db.Where("step <= ?", stepRange[1])
			}
		}
	}
}

func WithQueryIdx(idx int64) QueryOption {
	return func(db *gorm.DB) {
		if idx != 0 {
			db.Where("idx = ?", idx)
		}
	}
}

func WithQueryUserData(needUserData bool) QueryOption {
	return func(db *gorm.DB) {
		if needUserData {
			db.Preload("UserData")
		}
	}
}

func WithQueryRecordData(needRecordData bool) QueryOption {
	return func(db *gorm.DB) {
		if needRecordData {
			db.Preload("RecordData")
		}
	}
}

func WithQueryRecordID(recordId int64) QueryOption {
	return func(db *gorm.DB) {
		if recordId != 0 {
			db.Where("record_id = ?", recordId)
		}
	}
}

func WithQueryRecordBreakCount(count int) QueryOption {
	return func(db *gorm.DB) {
		if count != 0 {
			db.Where("record_break_count = ?", count)
		}
	}
}

func WithQueryBreakCountRange(breakCountRange []int) QueryOption {
	return func(db *gorm.DB) {
		if len(breakCountRange) <= 2 {
			if breakCountRange[0] != 0 {
				db.Where("record_break_count >= ?", breakCountRange[0])
			}
			if breakCountRange[1] != 0 {
				db.Where("record_break_count <= ?", breakCountRange[1])
			}
		}
	}
}

func WithQueryRankRange(rankRage []int) QueryOption {
	return func(db *gorm.DB) {
		if len(rankRage) <= 2 {
			if rankRage[0] != 0 {
				db.Where("ranked >= ?", rankRage[0])
			}
			if rankRage[1] != 0 {
				db.Where("ranked <= ?", rankRage[1])
			}
		}
	}
}
