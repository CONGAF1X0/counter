package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu sync.Mutex
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
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.m[key]
}

func (c *Counter) Reset() {
	c.mu.Lock()
	c.m = make(map[string]int)
	c.mu.Unlock()
}

func main() {
	counter := NewCounter()
	counter.Flush2broker(1000, func() {
		fmt.Println("flush")
	})
	for i := 0;i<100000 ; i++ {
		go counter.Incr("kkk", 1)
	}
	for {}
}
