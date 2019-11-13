package main

import (
	"testing"
)

// go test -bench=.
// go test -bench=. -benchmem

func BenchmarkSlice(b *testing.B) {
	var testSlice []int
	for i := 0; i < 102400; i++ {
		testSlice = append(testSlice, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		returnFromSlice(testSlice, 102399)
	}
}

func BenchmarkMap(b *testing.B) {
	testMap := make(map[int]int)
	for i := 0; i < 102400; i++ {
		testMap[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		returnFromMap(testMap, 102399)
	}
}
