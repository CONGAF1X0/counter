package main

import (
	"log"
	"sync"
	"sync/atomic"
)

type Counters struct {
	m sync.Map
}

type cou struct {
	name  string
	value int64
}

func NewCounters() *Counters {
	return &Counters{
		m: sync.Map{},
	}
}

func (c *Counters) Init(key string, fun func(*Counters, string)) {
	c.m.LoadOrStore(key, &cou{key, 0})
	if fun != nil {
		go fun(c, key)
	}
}

func (c *Counters) Incr(key string, num int) {
	co, ok := c.m.Load(key)
	if !ok {
		log.Fatal(key + "not init")
		return
	}
	atomic.AddInt64(&co.(*cou).value, int64(num))
}

func (c *Counters) Get(key string) int {
	n, ok := c.m.Load(key)
	if !ok {
		log.Fatal(key + "not init")
		return 0
	}
	return int(atomic.LoadInt64(&n.(*cou).value))
}

func (c *Counters) Reset() {
	c.m = sync.Map{}
}

func (c *Counters) ResetByKey(key string) int {
	co, ok := c.m.Load(key)
	if !ok {
		log.Fatal(key + "not init")
		return -1
	}
	for {
		old := atomic.LoadInt64(&co.(*cou).value)
		if atomic.CompareAndSwapInt64(&co.(*cou).value, old, 0) {
			return int(old)
		}
	}
}
