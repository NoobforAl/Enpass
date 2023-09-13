package caching

import (
	"sync"
	"time"
)

type SendCachePass struct {
	id       uint
	password string
}

type savedPass struct {
	passwords map[uint]string
	mux       sync.Mutex
	ch        chan SendCachePass
	timeDelay time.Duration
}

func New(timeDelay time.Duration) *savedPass {
	return &savedPass{
		passwords: make(map[uint]string, 1),
		ch:        make(chan SendCachePass),
		timeDelay: timeDelay,
	}
}

func (sp *savedPass) GetPass(id uint) (string, error) {
	sp.mux.Lock()
	defer sp.mux.Unlock()
	str, ok := sp.passwords[id]
	if !ok {
		return "", ErrNotFoundPass
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
	timer := time.NewTicker(sp.timeDelay)
	for {
		select {
		case pass := <-sp.ch:
			if pass.id != id {
				sp.ch <- pass
				continue
			}
			sp.mux.Lock()
			timer.Reset(sp.timeDelay)
			sp.passwords[pass.id] = pass.password
			sp.mux.Unlock()

		case <-timer.C:
			sp.mux.Lock()
			defer sp.mux.Unlock()
			delete(sp.passwords, id)
			return
		}
	}
}
