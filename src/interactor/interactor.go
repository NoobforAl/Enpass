package interactor

import (
	"github.com/NoobforAl/Enpass/contract"
)

type interActor struct {
	store contract.Store
}

func NewActorUser(
	store contract.Store,
) interActor {
	return interActor{
		store: store,
	}
}
