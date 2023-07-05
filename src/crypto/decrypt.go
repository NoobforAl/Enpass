package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"

	errs "github.com/NoobforAl/Enpass/errors"
)

func Decrypt(key, text string) (string, error) {
	textDecoded, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", errors.Join(errs.ErrDecrypt, err)
	}

	nonceSize := 12
	if len(textDecoded) < nonceSize {
		return "", errors.Join(errs.ErrDecrypt, errs.ErrTextIsShort)
	}
	nonce := textDecoded[:nonceSize]
	textDecoded = textDecoded[nonceSize:]

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", errors.Join(errs.ErrDecrypt, err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.Join(errs.ErrDecrypt, err)
	}

	plaintextBytes, err := aesgcm.Open(nil, nonce, textDecoded, nil)
	if err != nil {
		return "", errors.Join(errs.ErrDecrypt, err)
	}

	plaintext := string(plaintextBytes)
	return plaintext, nil
}
