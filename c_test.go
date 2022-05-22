package main

import (
	"testing"
)

/*
goos: windows
goarch: amd64
pkg: counter
cpu: Intel(R) Core(TM) i5-8300H CPU @ 2.30GHz
BenchmarkCounter
BenchmarkCounter-8    	 3785401	       329.5 ns/op
BenchmarkCounter2
BenchmarkCounter2-8   	 4908016	       251.4 ns/op
PASS
*/
func BenchmarkCounter(b *testing.B) {
	counter := NewCounter()

	for i := 0; i < b.N; i++ {
		go counter.Incr("key", 1)

	}
}

func BenchmarkCounter2(b *testing.B) {
	counter2 := NewCounters()
	counter2.Init("key", nil)
	for i := 0; i < b.N; i++ {
		go counter2.Incr("key", 1)
	}
}
