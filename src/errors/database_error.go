package errors

import (
	"errors"

	"github.com/NoobforAl/Enpass/lib/crypto"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	// transaction data can be zero
	ErrTransaction = errors.New("transaction data can be zero")

	// not found action in this action (update create delete)
	ErrNotFoundAction = errors.New("not found action")

	// not enough record for delete
	ErrNotEnoughRecord = errors.New("not enough record for delete")

	// security validation error
	ErrEncrypt      = crypto.ErrEncrypt
	ErrDecrypt      = crypto.ErrDecrypt
	ErrTextIsShort  = crypto.ErrTextIsShort
	ErrHashNotMatch = crypto.ErrHashNotMatch

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
