package model

import "time"

type Service struct {
	ID        uint  `gorm:"primarykey;uniq"`
	Name      Value `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
