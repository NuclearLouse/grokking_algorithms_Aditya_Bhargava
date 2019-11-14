package main

import (
	"testing"
)

// go test -bench=.
// go test -bench=. -benchmem

// func BenchmarkSlice(b *testing.B) {
// 	var testSlice []int
// 	for i := 0; i < 102400; i++ {
// 		testSlice = append(testSlice, i)
// 	}
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		returnFromSlice(testSlice, 102399)
// 	}
// }

// func BenchmarkMap(b *testing.B) {
// 	testMap := make(map[int]int)
// 	for i := 0; i < 102400; i++ {
// 		testMap[i] = i
// 	}
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		returnFromMap(testMap, 102399)
// 	}
// }

func BenchmarkDijckstra(b *testing.B) {
	g := newGraph()

	a1 := g.addNode("1")
	b2 := g.addNode("2")
	c3 := g.addNode("3")
	d4 := g.addNode("4")
	e5 := g.addNode("5")
	f6 := g.addNode("6")

	g.addEdge(a1, b2, 7)
	g.addEdge(a1, c3, 9)
	g.addEdge(a1, f6, 14)
	g.addEdge(b2, c3, 10)
	g.addEdge(b2, d4, 15)
	g.addEdge(c3, d4, 11)
	g.addEdge(f6, c3, 2)
	g.addEdge(e5, f6, 9)
	g.addEdge(d4, e5, 6)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.dijkstra(a1)
	}
}
