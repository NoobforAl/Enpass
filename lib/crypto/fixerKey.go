package crypto

func combineWithHash(key string, con int) string {
	hash := HashSha256(key)
	return key + hash[:con]
}

func fixLengthKey(key string) (string, error) {
	s := len(key)
	switch {
	case s < 17:
		return combineWithHash(key, 16-s), nil
	case s < 33:
		return combineWithHash(key, 32-s), nil
	default:
		return "", ErrTextIsShort
	}
}
