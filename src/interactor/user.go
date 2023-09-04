package interactor

import (
	"context"

	"github.com/NoobforAl/Enpass/entity"
	errs "github.com/NoobforAl/Enpass/errors"
	"github.com/NoobforAl/Enpass/lib/caching"
	"github.com/NoobforAl/Enpass/lib/crypto"
)

func (i interActor) FindUser(
	ctx context.Context,
	user entity.User,
) (entity.User, error) {
	i.log.Debug("Find User")
	u, err := i.store.GetUser(ctx, user)
	if err != nil {
		return user, err
	}

	i.log.Debug("check decrypt password")
	p, err := crypto.Decrypt(
		user.Password, u.Password)

	if err != nil {
		return user, err
	}

	i.log.Debug("check hash password")
	if !crypto.IsOkHash(user.Password, p) {
		return user, errs.ErrNotMatchPassword
	}

	if _, err = caching.CachedPass.
		GetPass(user.ID); err != nil {
		i.log.Debug("cache pass delete not found," +
			"start new delete password cache," +
			"after set timed.")
		go caching.CachedPass.DeletePass(user.ID)
	}

	i.log.Debug("set password in cache")
	caching.CachedPass.SetPass(user.ID, user.Password)
	return user, nil
}

func (i interActor) UpdateUser(
	ctx context.Context,
	old, new entity.User,
) (entity.User, error) {
	var pass string
	var err error

	i.log.Debug("update user password")
	if pass, err = caching.
		CachedPass.
		GetPass(old.ID); err != nil {
		return new, err
	}

	if pass != old.Password {
		return new, errs.ErrNotMatchPassword
	}

	return i.store.UpdateUser(ctx, old, new)
}
