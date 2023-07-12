package database

import (
	"context"
	"errors"
	"sync"

	"github.com/NoobforAl/Enpass/crypto"
	"github.com/NoobforAl/Enpass/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Stor struct {
	db *gorm.DB
}

var stor Stor
var onc sync.Once

func New(dsn string) (Stor, error) {
	var err error
	onc.Do(func() {
		if stor.db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}); err != nil {
			panic(err)
		}

		if err = migrate(); err != nil {
			panic(err)
		}

		if err = createUserIfNotExist(); err != nil {
			panic(err)
		}
	})
	return stor, err
}

func migrate() error {
	return stor.db.AutoMigrate(
		&User{},
		&Service{},
		&Password{},
	)
}

func createUserIfNotExist() error {
	var user entity.User
	user, err := stor.GetUser(context.TODO(), user)

	if !errors.Is(
		err, gorm.ErrRecordNotFound,
	) && err != nil {
		return err
	}

	if user.Password == "" {
		defaultPassword := "1111111111111111"
		user.Password = defaultPassword
		user.Password = crypto.HashSha256(user.Password)
		user.Password, err = crypto.Encrypt(
			defaultPassword, user.Password)

		if err != nil {
			return err
		}

		u := entityToModelUser(user)
		if err = stor.db.Model(&u).
			Save(&u).Error; err != nil {
			return err
		}
	}
	return nil
}
