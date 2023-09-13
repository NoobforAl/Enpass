package interactor

import (
	"github.com/NoobforAl/Enpass/contract"
)

type interActor struct {
	store contract.Store
	log   contract.Logger
	cache contract.Caching
}

func New(
	store contract.Store,
	log contract.Logger,
	cache contract.Caching,
) interActor {
	return interActor{
		store: store,
		log:   log,
		cache: cache,
	}
}
