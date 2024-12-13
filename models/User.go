package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"unique"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time // 在创建时，如果该字段值为零值，则使用当前时间填充
}
