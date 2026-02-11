package utils

import "gorm.io/gorm"

func Paginate(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// 1. 统一参数校验
		if page <= 0 {
			page = 1
		}
		switch {
		case limit > 100:
			limit = 100 // 限制单页最大数量，保护数据库
		case limit <= 0:
			limit = 10
		}

		// 2. 计算偏移量
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}
