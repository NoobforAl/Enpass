package model

import "time"

type UserPass struct {
	ID        uint `gorm:"primarykey"`
	EnPass    Value
	CreatedAt time.Time
	UpdatedAt time.Time
}
