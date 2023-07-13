package interactor

import (
	"context"

	"github.com/NoobforAl/Enpass/caching"
	"github.com/NoobforAl/Enpass/crypto"
	"github.com/NoobforAl/Enpass/entity"
	errs "github.com/NoobforAl/Enpass/errors"
)

func (i interActor) FindUser(
	ctx context.Context,
	user entity.User,
) (entity.User, error) {
	u, err := i.store.GetUser(ctx, user)
	if err != nil {
		return user, err
	}

	p, err := crypto.Decrypt(
		user.Password, u.Password)

	if err != nil {
		return user, err
	}

	if !crypto.IsOkHash(user.Password, p) {
		return user, errs.ErrNotMatchPassword
	}

	if _, err = caching.CachedPass.
		GetPass(user.ID); err != nil {
		go caching.CachedPass.DeletePass(user.ID)
	}

	caching.CachedPass.SetPass(user.ID, user.Password)
	return user, nil
}

func (i interActor) UpdateUser(
	ctx context.Context,
	old entity.User,
	new entity.User,
) (entity.User, error) {
	password, err := caching.CachedPass.
		GetPass(old.ID)

	if err != nil {
		return new, err
	}

	if password != old.Password {
		return new, errs.ErrNotMatchPassword
	}

	return i.store.UpdateUser(
		ctx, old, new)
}
