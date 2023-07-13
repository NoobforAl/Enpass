package interactor

import (
	"github.com/NoobforAl/Enpass/contract"
)

type interActor struct {
	store contract.Store
}

func New(
	store contract.Store,
) interActor {
	return interActor{
		store: store,
	}
}
