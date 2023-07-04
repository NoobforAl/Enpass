package database

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
)

type Value string

type Values struct {
	UserName Value
	Password Value
	Note     Value
	Hash     Value
}

func (v Value) EncryptValue(key string) (Value, error) {
	tmp, err := encrypt(key, v.String())
	return Value(tmp), err
}

func (v Value) DecryptValue(key string) (Value, error) {
	tmp, err := decrypt(key, v.String())
	return Value(tmp), err
}

func (v Value) HashSha256() Value {
	h := sha256.Sum224([]byte(v))
	return Value(hex.EncodeToString(h[:]))
}

func (v Value) IsOkHash(s string) bool {
	val := Value(s)
	return val.HashSha256() == v
}

func (v Value) String() string {
	return string(v)
}

func (v *Values) EncryptValues(key string) error {
	var (
		UserName Value
		Password Value
		Note     Value
		err      error
	)

	UserName, err = v.UserName.EncryptValue(key)
	if err != nil {
		return err
	}

	Password, err = v.Password.EncryptValue(key)
	if err != nil {
		return err
	}

	Note, err = v.Note.EncryptValue(key)
	if err != nil {
		return err
	}

	v.Hash = v.UserName + v.Password + v.Note
	v.Hash = v.Hash.HashSha256()

	v.UserName = UserName
	v.Password = Password
	v.Note = Note
	return nil
}

func (v *Values) DecryptValues(key string) error {
	var (
		UserName Value
		Password Value
		Note     Value
		err      error
	)

	UserName, err = v.UserName.DecryptValue(key)
	if err != nil {
		return err
	}

	Password, err = v.Password.DecryptValue(key)
	if err != nil {
		return err
	}

	Note, err = v.Note.DecryptValue(key)
	if err != nil {
		return err
	}

	v.UserName = UserName
	v.Password = Password
	v.Note = Note

	if !v.Hash.IsOkHash(string(UserName + Password + Note)) {
		return ErrHashNotMatch
	}

	return nil
}

func encrypt(key, text string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", errors.Join(ErrEncrypt, err)
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", errors.Join(ErrEncrypt, err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.Join(ErrEncrypt, err)
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(text), nil)
	ciphertext = append(nonce, ciphertext...)

	ciphertextString := base64.StdEncoding.EncodeToString(ciphertext)
	return ciphertextString, nil
}

func decrypt(key, text string) (string, error) {
	textDecoded, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", errors.Join(ErrDecrypt, err)
	}

	nonceSize := 12
	if len(textDecoded) < nonceSize {
		return "", errors.Join(ErrDecrypt, ErrTextIsShort)
	}
	nonce := textDecoded[:nonceSize]
	textDecoded = textDecoded[nonceSize:]

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", errors.Join(ErrDecrypt, err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.Join(ErrDecrypt, err)
	}

	plaintextBytes, err := aesgcm.Open(nil, nonce, textDecoded, nil)
	if err != nil {
		return "", errors.Join(ErrDecrypt, err)
	}

	plaintext := string(plaintextBytes)
	return plaintext, nil
}
