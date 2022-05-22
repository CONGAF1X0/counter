package main

import (
	"sync"
	"time"
)

type Counter struct {
	mu sync.RWMutex
	m  map[string]int
}

func NewCounter() *Counter {
	return &Counter{
		m: make(map[string]int),
	}
}

func (c *Counter) Flush2broker(ttl int, fun func()) {
	go func() {
		ticker := time.NewTicker(time.Duration(ttl) * time.Millisecond)
		for {
			<-ticker.C
			fun()
			c.Reset()
		}
	}()
}

func (c *Counter) Incr(key string, num int) {
	c.mu.Lock()
	c.m[key] += num
	c.mu.Unlock()
}

func (c *Counter) Get(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.m[key]
}

func (c *Counter) Reset() {
	c.mu.Lock()
	c.m = make(map[string]int)
	c.mu.Unlock()
}


