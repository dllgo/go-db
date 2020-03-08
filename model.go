package gormplus

import (
	"errors"
	"time"

	"github.com/dllgo/go-utils"
	"github.com/jinzhu/gorm"
)

type Model struct {
	ID        uint       `gorm:"column:id;primary_key;auto_increment;" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index;" json:"deleted_at"`
}

type ModelID struct {
	ID uint `gorm:"column:id;primary_key;auto_increment;" json:"id"`
}

// FindPage 查询分页数据
func FindPage(db *gorm.DB, pageIndex, pageSize int64, out interface{}) (int64, error) {
	var count int64
	result := db.Count(&count)
	if err := result.Error; err != nil {
		return 0, err
	} else if count == 0 {
		return 0, nil
	}
	// 如果分页大小小于0，则不查询数据
	if pageSize < 0 || pageIndex < 0 {
		return count, nil
	}
	if pageIndex > 0 && pageSize > 0 {
		db = db.Offset((pageIndex - 1) * pageSize)
	}
	if pageSize > 0 {
		db = db.Limit(pageSize)
	}
	if err := db.Find(out).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// FindOne 查询单条数据
func FindOne(db *gorm.DB, out interface{}) (bool, error) {
	if err := db.First(out).Error; err != nil {
		return false, errors.New("数据不存在")
	}
	return true, nil
}

// Check 检查数据是否存在
func Check(db *gorm.DB) (bool, error) {
	var count int
	result := db.Count(&count)
	if err := result.Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func ToString(v interface{}) string {
	return utils.JSONMarshalToString(v)
}
