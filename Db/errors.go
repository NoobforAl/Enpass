package Db

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	// not found action in this action (update create delete)
	ErrNotFoundAction = errors.New("not found action")
	// not enough record for delete
	ErrNotEnoughRecord = errors.New("not enough record for delete")

	// all database error gorm

	ErrInvalidDB                     = gorm.ErrInvalidDB
	ErrEmptySlice                    = gorm.ErrEmptySlice
	ErrRegistered                    = gorm.ErrRegistered
	ErrInvalidData                   = gorm.ErrInvalidData
	ErrInvalidValue                  = gorm.ErrInvalidValue
	ErrInvalidField                  = gorm.ErrInvalidField
	ErrDuplicatedKey                 = gorm.ErrDuplicatedKey
	ErrNotImplemented                = gorm.ErrNotImplemented
	ErrRecordNotFound                = logger.ErrRecordNotFound
	ErrSubQueryRequired              = gorm.ErrSubQueryRequired
	ErrUnsupportedDriver             = gorm.ErrUnsupportedDriver
	ErrPreloadNotAllowed             = gorm.ErrPreloadNotAllowed
	ErrPrimaryKeyRequired            = gorm.ErrPrimaryKeyRequired
	ErrModelValueRequired            = gorm.ErrModelValueRequired
	ErrMissingWhereClause            = gorm.ErrMissingWhereClause
	ErrInvalidTransaction            = gorm.ErrInvalidTransaction
	ErrUnsupportedRelation           = gorm.ErrUnsupportedRelation
	ErrInvalidValueOfLength          = gorm.ErrInvalidValueOfLength
	ErrDryRunModeUnsupported         = gorm.ErrDryRunModeUnsupported
	ErrModelAccessibleFieldsRequired = gorm.ErrModelAccessibleFieldsRequired
)
