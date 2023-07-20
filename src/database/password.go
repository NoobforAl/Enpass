package database

import (
	"context"
	"time"

	"github.com/NoobforAl/Enpass/entity"
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

func entityToModelPass(
	pass entity.Password,
	key string,
	decrypt bool,
) (Password, error) {
	var err error

	password := Password{
		ID:        pass.ID,
		UserID:    pass.UserID,
		ServiceID: pass.ServiceID,
		Values: Values{
			UserName: pass.UserName,
			Password: pass.Password,
			Note:     pass.Note,
		},
	}

	if decrypt {
		err = password.EncryptValues(key)
	}

	return password, err
}

func modelToEntityPass(
	pass Password,
	key string,
	decrypt bool,
) (entity.Password, error) {
	var err error
	if decrypt {
		err = pass.DecryptValues(key)
	}

	return entity.Password{
		ID:        pass.ID,
		UserID:    pass.UserID,
		ServiceID: pass.ServiceID,
		UserName:  pass.UserName,
		Password:  pass.Password,
		Note:      pass.Note,
	}, err
}

func (s Stor) GetPassword(
	ctx context.Context,
	pass entity.Password,
	key string,
	decrypt bool,
) (entity.Password, error) {
	password, err := entityToModelPass(
		pass, "", false)

	if err != nil {
		return pass, err
	}

	err = s.db.Model(&password).
		WithContext(ctx).
		Where("id = ?", password.ID).
		First(&password).Error

	if err != nil {
		return pass, err
	}

	return modelToEntityPass(
		password, key, decrypt)
}

func (s Stor) GetManyPassword(
	ctx context.Context,
	key string,
	decrypt bool,
) ([]entity.Password, error) {
	var data []*Password
	var err error

	if err = s.db.Model(&Password{}).
		WithContext(ctx).
		Find(&data).Error; err != nil {
		return nil, err
	}

	passwords := make([]entity.Password, len(data))
	for i := range passwords {
		passwords[i], err = modelToEntityPass(
			*data[i], key, decrypt)

		if err != nil {
			return nil, err
		}
	}

	return passwords, nil
}

func (s Stor) InsertPassword(
	ctx context.Context,
	pass entity.Password,
	key string,
) (entity.Password, error) {
	password, err := entityToModelPass(
		pass, key, true)

	if err != nil {
		return pass, err
	}

	err = s.db.Model(&password).
		WithContext(ctx).
		Create(&password).Error

	if err != nil {
		return pass, err
	}

	return modelToEntityPass(password, key, true)
}

func (s Stor) UpdatePassword(
	ctx context.Context,
	pass entity.Password,
	key string,
) (entity.Password, error) {
	password, err := entityToModelPass(
		pass, key, true)

	if err != nil {
		return pass, err
	}

	err = s.db.Model(&password).
		WithContext(ctx).
		Save(&password).Error

	if err != nil {
		return pass, err
	}

	return modelToEntityPass(password, key, true)
}

func (s Stor) DeletePassword(
	ctx context.Context,
	pass entity.Password,
) (entity.Password, error) {
	password, err := entityToModelPass(
		pass, "", false)

	if err != nil {
		return pass, err
	}

	err = s.db.Model(password).
		Where("id = ?", password.ID).
		First(&password).
		Delete(&password).Error

	if err != nil {
		return pass, err
	}

	return modelToEntityPass(password, "", false)
}
