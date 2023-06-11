package controller

import (
	"sync"
	"time"

	env "github.com/NoobforAl/Enpass/loadEnv"
)

var cachedPass savedPass
var timeDelay = env.GetLifeTime()

type dataPass struct {
	id       uint
	password string
}

type savedPass struct {
	passwords map[uint]string
	mux       sync.Mutex
	ch        chan dataPass
}

func init() {
	// create new password saver
	cachedPass = savedPass{
		passwords: make(map[uint]string, 10),
		ch:        make(chan dataPass),
	}
}

func (sp *savedPass) getPass(id uint) (string, error) {
	sp.mux.Lock()
	defer sp.mux.Unlock()
	str, ok := sp.passwords[id]
	if !ok {
		return "", ErrNotFoundPass
	}
	return str, nil
}

func (sp *savedPass) setPass(id uint, pass string) {
	sp.mux.Lock()
	defer sp.mux.Unlock()
	sp.passwords[id] = pass
	go func() { sp.ch <- dataPass{id: id, password: pass} }()
}

func (sp *savedPass) deletePass(id uint) {
	for {
		select {
		case pass := <-sp.ch:
			if pass.id != id {
				sp.ch <- pass
				continue
			}
			sp.mux.Lock()
			sp.passwords[pass.id] = pass.password
			sp.mux.Unlock()

		case <-time.After(timeDelay):
			sp.mux.Lock()
			defer sp.mux.Unlock()
			delete(sp.passwords, id)
			return
		}
	}
}
