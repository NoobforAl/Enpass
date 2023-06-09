package Db

import (
	model "github.com/NoobforAl/Enpass/Model"
)

type Model interface {
	*model.Service | *model.SavedPassword | *model.UserPass
}

func Get[T Model](m T) error {
	return db.Where(m).First(m).Error
}

func GetMany[T Model](m T) ([]T, error) {
	var values []T
	return values, db.Model(m).Find(&values).Error
}

func Insert[T Model](value T) error {
	return db.Model(value).Create(value).Error
}

func InsertMany[T Model](values []T) error {
	return transaction("create", values)
}

func Update[T Model](m T) error {
	return db.Model(m).Updates(m).Error
}

func UpdateMany[T Model](values []T) error {
	return transaction("update", values)
}

func Delete[T Model](m T) error {
	return db.Model(m).Delete(m).Error
}

func DeleteMany[T Model](values []T) error {
	return transaction("delete", values)
}

func transaction[T Model](t string, values []T) error {
	if len(values) == 0 {
		return nil
	}

	tx := db.Begin()
	var err error

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		}
	}()

	action := tx.Model(values[0])
	for _, v := range values {
		switch t {
		case "create":
			if err = action.Create(v).Error; err != nil {
				tx.Rollback()
				return err
			}
		case "delete":
			if err = action.Delete(v).Error; err != nil {
				tx.Rollback()
				return err
			}
		case "update":
			if err = action.Updates(v).Error; err != nil {
				tx.Rollback()
				return err
			}
		default:
			return ErrNotFoundAction
		}
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
