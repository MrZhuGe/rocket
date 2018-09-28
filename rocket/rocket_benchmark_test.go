package rocket_test

import (
	"../rocket"
	"testing"
	"time"
)

var n = 100000

func doSomething() {
	time.Sleep(500 * time.Millisecond)
}

func BenchmarkRocketPool(b *testing.B) {
	b.N = n
	b.ReportAllocs()

	for j := 0; j < n; j++ {
		rocket.Add(doSomething)
	}
}

func BenchmarkGoroutine(b *testing.B) {
	b.N = n
	b.ReportAllocs()

	for j := 0; j < n; j++ {
		go doSomething()
	}
}