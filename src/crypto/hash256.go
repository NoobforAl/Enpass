package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashSha256(
	v string,
) string {
	h := sha256.Sum256([]byte(v))
	return hex.EncodeToString(h[:])
}

func IsOkHash(
	s string, hash string,
) bool {
	return HashSha256(s) == hash
}
