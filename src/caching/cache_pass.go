package caching

import (
	"sync"
	"time"

	env "github.com/NoobforAl/Enpass/config_loader"
	errs "github.com/NoobforAl/Enpass/errors"
)

var CachedPass savedPass
var timeDelay = env.GetLifeTime()

type SendCachePass struct {
	id       uint
	password string
}

type savedPass struct {
	passwords map[uint]string
	mux       sync.Mutex
	ch        chan SendCachePass
}

func init() {
	// create new password saver
	CachedPass = savedPass{
		passwords: make(map[uint]string, 1),
		ch:        make(chan SendCachePass),
	}
}

func (sp *savedPass) GetPass(id uint) (string, error) {
	sp.mux.Lock()
	defer sp.mux.Unlock()
	str, ok := sp.passwords[id]
	if !ok {
		return "", errs.ErrNotFoundPass
	}
	return str, nil
}

func (sp *savedPass) SetPass(id uint, pass string) {
	sp.mux.Lock()
	defer sp.mux.Unlock()
	sp.passwords[id] = pass
	go func() { sp.ch <- SendCachePass{id: id, password: pass} }()
}

func (sp *savedPass) DeletePass(id uint) {
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
