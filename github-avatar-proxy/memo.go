package main

import "time"

type Func func(string) (interface{}, error)

type memo struct {
	requests chan request
}

type request struct {
	key      string
	response chan result
}

type result struct {
	value interface{}
	err   error
}

func NewMemo(f Func) *memo {
	m := memo{
		requests: make(chan request),
	}

	go m.listen(f)

	return &m
}

func (m *memo) Get(key string) (interface{}, error) {
	response := make(chan result)

	m.requests <- request{
		key:      key,
		response: response,
	}

	res := <-response

	return res.value, res.err
}

func (m *memo) listen(f Func) {
	cache := make(map[string]*entry)

	for req := range m.requests {
		e := cache[req.key]
		if e == nil || time.Until(e.expire) < 0 {
			e = &entry{
				ready:  make(chan struct{}),
				expire: time.Now().Add(time.Hour),
			}
			cache[req.key] = e

			go e.call(f, req.key)
		}

		go e.deliver(req.response)
	}
}

type entry struct {
	res    result
	expire time.Time
	ready  chan struct{}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)

	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready

	response <- e.res
}