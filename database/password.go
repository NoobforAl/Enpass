package database

import (
	"context"
	"time"
)

type Password struct {
	ID uint `gorm:"primarykey"`

	ServiceID uint
	Service   Service `gorm:"foreignKey:ServiceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	UserID uint
	User   User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	Values

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s Stor) NewPassword(id, serId, userId uint,
	userName, password, note, hash string) *Password {
	return &Password{
		ID:        id,
		ServiceID: serId,
		UserID:    userId,
		Values: Values{
			UserName: Value(userName),
			Password: Value(password),
			Note:     Value(note),
			Hash:     Value(hash),
		},
	}
}

func (s Stor) GetPassword(ctx context.Context, id uint) (*Password, error) {
	ser := s.NewPassword(id, 0, 0, "", "", "", "")
	return ser, get(ctx, ser)
}

func (s Stor) GetManyPassword(ctx context.Context) ([]*Password, error) {
	return getMany(ctx, s.NewPassword(0, 0, 0, "", "", "", ""))
}

func (s Stor) InsertPassword(ctx context.Context, value Password) error {
	return insert(ctx, &value)
}

func (s Stor) InsertManyPassword(ctx context.Context, values []*Password) error {
	return insertMany(ctx, values)
}

func (s Stor) UpdatePassword(ctx context.Context, m Password) error {
	err := get(ctx, s.NewPassword(m.ID, m.ServiceID, m.UserID, "", "", "", ""))
	if err != nil {
		return err
	}
	return update(ctx, &m)
}

func (s Stor) UpdateManyPassword(ctx context.Context, values []*Password) error {
	return updateMany(ctx, values)
}

func (s Stor) DeletePassword(ctx context.Context, id uint) error {
	return delete(ctx, s.NewPassword(id, 0, 0, "", "", "", ""))
}

func (s Stor) DeleteManyPassword(ctx context.Context, values []*Password) error {
	return deleteMany(ctx, values)
}
