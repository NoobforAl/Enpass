package interactor

import (
	"github.com/NoobforAl/Enpass/contract"
)

type interActor struct {
	store contract.Store
	log   contract.Logger
}

func New(
	store contract.Store,
	log contract.Logger,
) interActor {
	return interActor{
		store: store,
		log:   log,
	}
}
