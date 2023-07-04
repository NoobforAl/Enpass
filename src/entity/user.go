package entity

import (
	"context"
	"errors"

	"github.com/NoobforAl/Enpass/contract"
	"github.com/gin-gonic/gin"
)

type User struct {
	Password    string `form:"password" json:"password" binding:"required"`
	NewPassword string `form:"newpassword" json:"newpassword"`
}

func (u *User) Pars(c *gin.Context) error {
	return parsJsonAndValidate(c, u)
}

func (u *User) FindUser(ctx context.Context, stor contract.Stor) (uint, error) {
	listUserPass, err := stor.GetManyUser(ctx)
	if err != nil {
		return 0, err
	}

	for _, v := range listUserPass {
		tmp, _ := v.EnPass.DecryptValue(u.Password)
		if tmp.IsOkHash(u.Password) {
			return v.ID, nil
		}
	}
	return 0, ErrNotFoundUser
}

func (u *User) UpdateUser(ctx context.Context, stor contract.Stor) error {
	userId, err := u.FindUser(ctx, stor)
	if err != nil {
		return err
	}

	oldUser := stor.NewUser(userId, u.Password)
	oldUser.EnPass, err = oldUser.EnPass.EncryptValue(u.NewPassword)
	if err != nil {
		return err
	}

	user := stor.NewUser(userId, u.NewPassword)
	user.EnPass, err = user.EnPass.EncryptValue(u.NewPassword)
	if err != nil {
		return err
	}

	passwords, err := stor.GetManyPassword(ctx)
	if err != nil {
		return err
	}

	for i := range passwords {
		_ = passwords[i].Values.DecryptValues(u.Password)
		_ = passwords[i].Values.EncryptValues(u.NewPassword)
	}

	if err = stor.UpdateManyPassword(ctx, passwords); err != nil {
		return err
	}

	if err = stor.UpdateUser(ctx, *user); err != nil {
		return errors.Join(ErrUpdatePassDone, err)
	}

	return nil
}
