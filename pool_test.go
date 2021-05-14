package main

import (
	"sync"
	"testing"
)

type Counter struct {
	A int
	B int
}

func NewCounter() *Counter {
	return new(Counter)
}

func (c *Counter) Increase() {
	c.A++
	c.B++
}

var counterPool = sync.Pool{
	New: func() interface{} { return NewCounter() },
}

func BenchmarkWithoutPool(b *testing.B) {
	var c *Counter
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			c = NewCounter()
			b.StopTimer()
			c.Increase()
			b.StartTimer()
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var c *Counter
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			c = counterPool.Get().(*Counter)
			b.StopTimer()
			c.Increase()
			b.StartTimer()
		}
	}
}
