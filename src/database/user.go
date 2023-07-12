package database

import (
	"context"
	"time"

	"github.com/NoobforAl/Enpass/entity"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	EnPass    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func entityToModelUser(
	user entity.User,
) User {
	return User{
		ID:     user.ID,
		EnPass: user.Password,
	}
}

func modelToEntityUser(
	user User,
) entity.User {
	return entity.User{
		ID:       user.ID,
		Password: user.EnPass,
	}
}

func (s Stor) GetUser(
	ctx context.Context,
	u entity.User,
) (entity.User, error) {
	user := entityToModelUser(u)

	err := s.db.Model(&user).
		WithContext(ctx).
		Where("id = ?", user.ID).
		First(&user).Error

	return modelToEntityUser(user), err
}

func (s Stor) UpdateUser(
	ctx context.Context,
	old entity.User,
	new entity.User,
) (entity.User, error) {
	var Pass []*Password
	err := s.db.Model(&Password{}).
		Find(&Pass).Error

	if err != nil {
		return new, err
	}

	for i := range Pass {
		if err = Pass[i].DecryptValues(
			old.Password,
		); err != nil {
			return new, err
		}

		if err = Pass[i].EncryptValues(
			new.Password,
		); err != nil {
			return new, err
		}
	}

	tx := s.db.Begin()

	if err = tx.Model(&User{}).
		Save(&new).Error; err != nil {
		tx.Rollback()
		return new, err
	}

	for i := range Pass {
		if err = tx.Model(&Password{}).
			Save(Pass[i]).Error; err != nil {
			tx.Rollback()
			return new, err
		}
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return new, err
	}

	return new, nil
}
