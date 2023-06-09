package model

type Value string

func (v Value) String() string {
	return string(v)
}

func (v Value) EncryptValue(key string) (Value, error) {
	tmp, err := encrypt(key, v.String())
	return Value(tmp), err
}

func (v Value) DecryptValue(key string) (Value, error) {
	tmp, err := decrypt(key, v.String())
	return Value(tmp), err
}
