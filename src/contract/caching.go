package contract

type Caching interface {
	SetPass(id uint, pass string)
	DeletePass(id uint)
	GetPass(id uint) (string, error)
}
