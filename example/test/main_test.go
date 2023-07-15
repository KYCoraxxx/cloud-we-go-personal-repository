package test

import (
	"fmt"
	"runtime"
	"testing"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello, world")
	}
}

func BenchmarkHelloWorldParallel(b *testing.B) {
	runtime.GOMAXPROCS(4)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fmt.Sprintf("hello, world")
		}
	})
}
