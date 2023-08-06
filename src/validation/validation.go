package validation

import "github.com/NoobforAl/Enpass/contract"

type validator struct {
	log contract.Logger
}

func New(log contract.Logger) validator {
	return validator{log: log}
}
