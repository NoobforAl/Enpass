package database

import (
	"github.com/NoobforAl/Enpass/crypto"
	errs "github.com/NoobforAl/Enpass/errors"
)

type Values struct {
	UserName string
	Password string
	Note     string
	Hash     string
}

func (v *Values) EncryptValues(
	key string,
) error {
	var (
		UserName string
		Password string
		Note     string
		err      error
	)

	UserName, err = crypto.Encrypt(
		key, v.UserName)

	if err != nil {
		return err
	}

	Password, err = crypto.Encrypt(
		key, v.Password)

	if err != nil {
		return err
	}

	Note, err = crypto.Encrypt(
		key, v.Note)

	if err != nil {
		return err
	}

	v.Hash = v.UserName +
		v.Password +
		v.Note

	v.Hash = crypto.HashSha256(v.Hash)
	v.UserName = UserName
	v.Password = Password
	v.Note = Note
	return nil
}

func (v *Values) DecryptValues(
	key string,
) error {
	var (
		UserName string
		Password string
		Note     string
		err      error
	)

	UserName, err = crypto.Decrypt(
		key, v.UserName)

	if err != nil {
		return err
	}

	Password, err = crypto.Decrypt(
		key, v.Password)

	if err != nil {
		return err
	}

	Note, err = crypto.Decrypt(
		key, v.Note)

	if err != nil {
		return err
	}

	v.UserName = UserName
	v.Password = Password
	v.Note = Note

	if crypto.IsOkHash(
		UserName+Password+Note,
		v.Hash,
	) {
		return errs.ErrHashNotMatch
	}

	return nil
}
