package database

import (
	"context"

	"gorm.io/gorm"
)

type model interface {
	*Password |
		*Service |
		*User
}

func get[T model](ctx context.Context, m T) error {
	return withContextModel(ctx, m).Where(m).First(m).Error
}

func getMany[T model](ctx context.Context, m T) ([]T, error) {
	var values []T
	return values, withContextModel(ctx, m).Find(&values).Error
}

func insert[T model](ctx context.Context, value T) error {
	return withContextModel(ctx, value).Create(value).Error
}

func insertMany[T model](ctx context.Context, values []T) error {
	return transaction("create", ctx, values)
}

func update[T model](ctx context.Context, m T) error {
	return withContextModel(ctx, m).Updates(m).Error
}

func updateMany[T model](ctx context.Context, values []T) error {
	return transaction("update", ctx, values)
}

func delete[T model](ctx context.Context, m T) error {
	return withContextModel(ctx, m).Model(m).Delete(m).Error
}

func deleteMany[T model](ctx context.Context, values []T) error {
	return transaction("delete", ctx, values)
}

func withContextModel[T model](ctx context.Context, m T) *gorm.DB {
	return stor.db.WithContext(ctx).Model(m)
}

func transaction[T model](t string, ctx context.Context, values []T) error {
	if len(values) == 0 {
		return ErrTransaction
	}

	tx := withContextModel(ctx, values[0]).Begin()
	var err error

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		}
	}()

	for _, v := range values {
		switch t {
		case "create":
			if err = tx.Create(v).Error; err != nil {
				tx.Rollback()
				return err
			}
		case "delete":
			if err = tx.Delete(v).Error; err != nil {
				tx.Rollback()
				return err
			}
		case "update":
			if err = tx.Updates(v).Error; err != nil {
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
