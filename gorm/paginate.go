package gorm

import (
	"gorm.io/gorm"

	"github.com/gorm-gen/paginate"
)

// Paginate 分页
func Paginate[T paginate.Int](page, pageSize T) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}
		return db.Offset((int(page) - 1) * int(pageSize)).Limit(int(pageSize))
	}
}
