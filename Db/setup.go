package Db

import (
	"sync"

	model "github.com/NoobforAl/Enpass/Model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var onc sync.Once

func InitDB(dsn string) (*gorm.DB, error) {
	var err error

	onc.Do(func() {
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		if err = migrate(); err != nil {
			return
		}

		if err = createUserIfNotExist(); err != nil {
			return
		}
	})
	return db, err
}

func createUserIfNotExist() error {
	var user model.UserPass

	users, err := GetMany(&user)
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
		if err = Insert(&user); err != nil {
			return err
		}
	}
	return nil
}

func migrate() error {
	return db.AutoMigrate(
		&model.UserPass{},
		&model.Service{},
		&model.SavedPassword{},
	)
}
