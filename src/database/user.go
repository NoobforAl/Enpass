package database

import (
	"context"
	"time"

	"github.com/NoobforAl/Enpass/entity"
	"github.com/NoobforAl/Enpass/lib/crypto"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	EnPass    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s Stor) entityToModelUser(user entity.User) User {
	s.log.Debug("Entity To Model User")
	return User{
		ID:     user.ID,
		EnPass: user.Password,
	}
}

func (s Stor) modelToEntityUser(user User) entity.User {
	s.log.Debug("Model To Entity User")
	return entity.User{
		ID:       user.ID,
		Password: user.EnPass,
	}
}

func (s Stor) handelErrorTX(tx *gorm.DB) {
	if err := recover(); err != nil {
		tx.Rollback()
		s.log.Panic(err)
	}
}

func (s Stor) GetUser(
	ctx context.Context,
	u entity.User,
) (entity.User, error) {
	s.log.Debug("Get User")
	user := s.entityToModelUser(u)

	err := s.db.Model(&user).
		WithContext(ctx).
		Where("id = ?", user.ID).
		First(&user).Error

	return s.modelToEntityUser(user), err
}

func (s Stor) UpdateUser(
	ctx context.Context,
	old entity.User,
	new entity.User,
) (entity.User, error) {
	s.log.Debug("Update User")
	var Pass []*Password
	var err error

	s.log.Debug("Get All Password")
	if err = s.db.
		Model(&Password{}).
		Find(&Pass).Error; err != nil {
		return new, err
	}

	s.log.Debug("Decrypt All Password")
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

	s.log.Debug("Create New TX & setup handel Error")
	tx := s.db.WithContext(ctx).Begin()
	defer s.handelErrorTX(tx)

	s.log.Debug("SetUp New User Password")
	newUser := s.entityToModelUser(new)
	newUser.EnPass = crypto.HashSha256(newUser.EnPass)
	newUser.EnPass, err = crypto.Encrypt(
		new.Password,
		newUser.EnPass)

	if err != nil {
		tx.Rollback()
		return new, err
	}

	s.log.Debug("SetUp New User Password On DB")
	if err = tx.Model(&User{}).
		Where("id = ?", new.ID).
		Save(newUser).Error; err != nil {
		tx.Rollback()
		return new, err
	}

	s.log.Debug("Encrypt Passwords with New user password")
	for i := range Pass {
		if err = tx.Model(&Password{}).
			Where("id = ?", Pass[i].ID).
			Save(Pass[i]).Error; err != nil {
			tx.Rollback()
			return new, err
		}
	}

	s.log.Debug("commit Changes")
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return new, err
	}
	return new, nil
}
