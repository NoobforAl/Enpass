package model

import (
	"crypto/sha256"
	"encoding/hex"
)

func (v Value) HashSha256() Value {
	h := sha256.Sum224([]byte(v))
	return Value(hex.EncodeToString(h[:]))
}

func (hash Value) IsOkHash(s string) bool {
	val := Value(s)
	return val.HashSha256() == hash
}
