package database

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/NoobforAl/Enpass/crypto"
	errs "github.com/NoobforAl/Enpass/errors"
)

type Values struct {
	UserName string
	Password string
	Note     string
	Hash     string
}

func HashSha256(v string) string {
	h := sha256.Sum224([]byte(v))
	return hex.EncodeToString(h[:])
}

func IsOkHash(s string, hash string) bool {
	return HashSha256(s) == hash
}

func (v *Values) EncryptValues(key string) error {
	var (
		UserName string
		Password string
		Note     string
		err      error
	)

	UserName, err = crypto.Encrypt(key, v.UserName)
	if err != nil {
		return err
	}

	Password, err = crypto.Encrypt(key, v.Password)
	if err != nil {
		return err
	}

	Note, err = crypto.Encrypt(key, v.Note)
	if err != nil {
		return err
	}

	v.Hash = v.UserName + v.Password + v.Note
	v.Hash = HashSha256(v.Hash)

	v.UserName = UserName
	v.Password = Password
	v.Note = Note
	return nil
}

func (v *Values) DecryptValues(key string) error {
	var (
		UserName string
		Password string
		Note     string
		err      error
	)

	UserName, err = crypto.Decrypt(key, v.UserName)
	if err != nil {
		return err
	}

	Password, err = crypto.Decrypt(key, v.Password)
	if err != nil {
		return err
	}

	Note, err = crypto.Decrypt(key, v.Note)
	if err != nil {
		return err
	}

	v.UserName = UserName
	v.Password = Password
	v.Note = Note

	if IsOkHash(UserName+Password+Note, v.Hash) {
		return errs.ErrHashNotMatch
	}

	return nil
}
