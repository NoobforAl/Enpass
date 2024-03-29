package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

func Decrypt(key, text string) (string, error) {
	key, err := fixLengthKey(key)
	if err != nil {
		return "", err
	}

	textDecoded, err := base64.
		StdEncoding.
		DecodeString(text)

	if err != nil {
		return "",
			errors.Join(ErrDecrypt, err)
	}

	if len(textDecoded) < nonceSize {
		return "",
			errors.Join(ErrDecrypt, ErrTextIsShort)
	}

	nonce := textDecoded[:nonceSize]
	textDecoded = textDecoded[nonceSize:]

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "",
			errors.Join(ErrDecrypt, err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "",
			errors.Join(ErrDecrypt, err)
	}

	plaintextBytes, err := aesgcm.Open(
		nil, nonce, textDecoded, nil)

	if err != nil {
		return "",
			errors.Join(ErrDecrypt, err)
	}

	plaintext := string(plaintextBytes)
	return plaintext, nil
}
