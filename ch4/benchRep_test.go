package main

import (
	"testing"

	"github.com/kaepa3/parallel_golang/ch4/common"
)

func BenchmarkGeneric(b *testing.B) {
	done := make(chan interface{})

	defer close(done)
	b.ResetTimer()
	for range common.ToString(done, common.Take(done, common.Repeat(done, "a"), b.N)) {
	}
}

func BenchmarkTyped(b *testing.B) {
	done := make(chan interface{})
	defer close(done)
	b.ResetTimer()
	for range common.Take(done, common.Repeat(done, "a"), b.N) {
	}
}
