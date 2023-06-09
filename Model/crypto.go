package model

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

func encrypt(key, text string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", errors.Join(EncryptError, err)
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", errors.Join(EncryptError, err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.Join(EncryptError, err)
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(text), nil)
	ciphertext = append(nonce, ciphertext...)

	ciphertextString := base64.StdEncoding.EncodeToString(ciphertext)
	return ciphertextString, nil
}

func decrypt(key, text string) (string, error) {
	textDecoded, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", errors.Join(DecryptError, err)
	}

	nonceSize := 12
	if len(textDecoded) < nonceSize {
		return "", errors.Join(DecryptError, TextIsShortError)
	}
	nonce := textDecoded[:nonceSize]
	textDecoded = textDecoded[nonceSize:]

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", errors.Join(DecryptError, err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.Join(DecryptError, err)
	}

	plaintextBytes, err := aesgcm.Open(nil, nonce, textDecoded, nil)
	if err != nil {
		return "", errors.Join(DecryptError, err)
	}

	plaintext := string(plaintextBytes)
	return plaintext, nil
}
