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

func (s Stor) entityToModelPass(
	pass entity.Password,
	key string,
	decrypt bool,
) (Password, error) {
	s.log.Debug("Convert Entity To Model Pass")

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

func (s Stor) modelToEntityPass(
	pass Password,
	key string,
	decrypt bool,
) (entity.Password, error) {
	s.log.Debug("Convert Model To Entity Pass")
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
	s.log.Debugf("Get Password, with decrypt? %b", decrypt)
	password, _ := s.entityToModelPass(pass, "", false)

	if err := s.db.Model(&password).
		WithContext(ctx).
		Where("id = ?", password.ID).
		First(&password).Error; err != nil {
		return pass, err
	}

	return s.modelToEntityPass(password, key, decrypt)
}

func (s Stor) GetManyPassword(
	ctx context.Context,
	key string,
	decrypt bool,
) ([]entity.Password, error) {
	s.log.Debugf("Get Many Password, with decrypt? %b", decrypt)
	var data []*Password
	var err error

	if err = s.db.Model(&Password{}).
		WithContext(ctx).
		Find(&data).Error; err != nil {
		return nil, err
	}

	passwords := make([]entity.Password, len(data))
	for i := range passwords {
		passwords[i], err = s.modelToEntityPass(
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
	s.log.Debug("Insert New Password")
	password, err := s.entityToModelPass(pass, key, true)

	if err != nil {
		return pass, err
	}

	if err = s.db.Model(&password).
		WithContext(ctx).
		Create(&password).Error; err != nil {
		return pass, err
	}

	return s.modelToEntityPass(password, key, true)
}

func (s Stor) UpdatePassword(
	ctx context.Context,
	pass entity.Password,
	key string,
) (entity.Password, error) {
	s.log.Debug("Update Password")
	password, err := s.entityToModelPass(pass, key, true)

	if err != nil {
		return pass, err
	}

	if err = s.db.Model(&password).
		WithContext(ctx).
		Save(&password).Error; err != nil {
		return pass, err
	}

	return s.modelToEntityPass(password, key, true)
}

func (s Stor) DeletePassword(
	ctx context.Context,
	pass entity.Password,
) (entity.Password, error) {
	s.log.Debug("Delete Password")
	password, err := s.entityToModelPass(pass, "", false)

	if err != nil {
		return pass, err
	}

	if err = s.db.Model(password).
		Where("id = ?", password.ID).
		First(&password).
		Delete(&password).Error; err != nil {
		return pass, err
	}

	return s.modelToEntityPass(password, "", false)
}
