package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

const nonceSize = 12

func Encrypt(key, text string) (string, error) {
	key, err := fixLengthKey(key)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "",
			errors.Join(ErrEncrypt, err)
	}

	nonce := make([]byte, nonceSize)
	if _, err := io.ReadFull(
		rand.Reader, nonce,
	); err != nil {
		return "", errors.Join(ErrEncrypt, err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "",
			errors.Join(ErrEncrypt, err)
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(text), nil)
	ciphertext = append(nonce, ciphertext...)

	return base64.StdEncoding.
		EncodeToString(ciphertext), nil
}
