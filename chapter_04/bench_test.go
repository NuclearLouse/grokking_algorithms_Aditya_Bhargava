package main

// go test -bench=.

import "testing"

func BenchmarkSorting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := []int{1, 25, 2, 4, 18, 3, 67, 5, 9, 10, 33, 99, 6, 8, 3, 2, 12, 15, 17, 19, 98, 56, 77, 88, 43, 44, 45, 23, 21, 40, 60, 70, 20, 47, 54}
		sorting(m)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := []int{1, 25, 2, 4, 18, 3, 67, 5, 9, 10, 33, 99, 6, 8, 3, 2, 12, 15, 17, 19, 98, 56, 77, 88, 43, 44, 45, 23, 21, 40, 60, 70, 20, 47, 54}
		quickSort(m)
	}
}

func BenchmarkQuickSortRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := []int{1, 25, 2, 4, 18, 3, 67, 5, 9, 10, 33, 99, 6, 8, 3, 2, 12, 15, 17, 19, 98, 56, 77, 88, 43, 44, 45, 23, 21, 40, 60, 70, 20, 47, 54}
		quickSortRand(m)
	}
}
