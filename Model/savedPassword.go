package model

import "time"

type SavedPassword struct {
	ID uint `gorm:"primarykey"`

	ServiceID uint
	Service   Service `gorm:"foreignKey:ServiceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	UserPassID uint
	UserPass   UserPass `gorm:"foreignKey:UserPassID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	Values

	CreatedAt time.Time
	UpdatedAt time.Time
}
