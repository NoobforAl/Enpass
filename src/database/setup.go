package database

import (
	"context"
	"errors"
	"sync"

	"github.com/NoobforAl/Enpass/contract"
	"github.com/NoobforAl/Enpass/entity"
	"github.com/NoobforAl/Enpass/lib/crypto"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Stor struct {
	db  *gorm.DB
	log contract.Logger
}

var stor Stor
var onc sync.Once

func New(dsn string, log contract.Logger) (Stor, error) {
	var err error
	onc.Do(func() {
		stor.log = log
		log.Debugf("setUp new database with this DSN: %s", dsn)
		if stor.db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}); err != nil {
			log.Panic(err)
		}

		log.Debug("Migrate Database")
		if err = migrate(); err != nil {
			log.Panic(err)
		}

		log.Info("Create New User If not Exist!")
		if err = createUserIfNotExist(); err != nil {
			log.Panic(err)
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
	stor.log.Debug("Get All Users")

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	var user entity.User
	user, err := stor.GetUser(ctx, user)

	if !errors.Is(
		err, gorm.ErrRecordNotFound,
	) && err != nil {
		return err
	}

	if user.Password == "" {
		stor.log.Warn("User Not Found Create New User")
		defaultPassword := "1111"
		user.Password = defaultPassword
		user.Password = crypto.HashSha256(user.Password)
		user.Password, err = crypto.Encrypt(
			defaultPassword, user.Password)

		if err != nil {
			return err
		}

		u := stor.entityToModelUser(user)
		if err = stor.db.
			WithContext(ctx).
			Model(&u).
			Save(&u).Error; err != nil {
			return err
		}
	}
	return nil
}
