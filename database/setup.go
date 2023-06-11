package database

import (
	"context"
	"sync"

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
			return
		}

		if err = createUserIfNotExist(); err != nil {
			return
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
	var user User
	users, err := getMany(context.TODO(), &user)
	if err != nil {
		return err
	}

	if len(users) == 0 {
		user.EnPass = "1111111111111111"
		user.EnPass = user.EnPass.HashSha256()
		user.EnPass, err = user.EnPass.EncryptValue("1111111111111111")
		if err != nil {
			return err
		}
		if err = insert(context.TODO(), &user); err != nil {
			return err
		}
	}
	return nil
}
