package model

type Values struct {
	UserName Value
	Password Value
	Note     Value
	Hash     Value
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
