package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Сортировка где базовый элемент берется первым в списке
func sorting(m []int) []int {
	if len(m) < 2 {
		return m
	}
	base := m[0]
	// fmt.Printf("Base index:%d | Base element:%d\n", 0, base)
	var leftM, rightM []int
	for _, e := range m[1:] {
		switch {
		case e <= base:
			leftM = append(leftM, e)
		case e > base:
			rightM = append(rightM, e)
		}
	}
	// fmt.Printf("LeftM:%v | RightM:%v\n", leftM, rightM)
	return append(append(sorting(leftM), base), sorting(rightM)...)
}

func transposition(m []int) []int {
	if m[0] > m[1] {
		m = []int{m[1], m[0]}
	}
	return m
}

// Сортировка где базовый элемент берется средним элементом списка
// а когда список доходит до двух элементов применяется транспозиция
func quickSort(m []int) []int {
	if len(m) < 2 {
		return m
	}
	if len(m) == 2 {
		return transposition(m)
	}
	base := len(m) / 2
	// fmt.Printf("Base index:%d | Base element:%d\n", base, m[base])
	var leftM, rightM []int
	for _, e := range m[:base] {
		switch {
		case e <= m[base]:
			leftM = append(leftM, e)
		case e > m[base]:
			rightM = append(rightM, e)
		}
	}
	for _, e := range m[base+1:] {
		switch {
		case e <= m[base]:
			leftM = append(leftM, e)
		case e > m[base]:
			rightM = append(rightM, e)
		}
	}
	// fmt.Printf("LeftM:%v | RightM:%v\n", leftM, rightM)
	return append(append(quickSort(leftM), m[base]), quickSort(rightM)...)
}

func random(limit int) int {
	sr := rand.NewSource(time.Now().UnixNano())
	r := rand.New(sr)
	return r.Intn(limit)
}

// Сортировка где базовым элементом берется случайный элемент списка
func quickSortRand(m []int) []int {
	if len(m) < 2 {
		return m
	}
	base := random(len(m))
	var leftM, rightM []int
	for _, e := range m[:base] {
		switch {
		case e <= m[base]:
			leftM = append(leftM, e)
		case e > m[base]:
			rightM = append(rightM, e)
		}
	}
	for _, e := range m[base+1:] {
		switch {
		case e <= m[base]:
			leftM = append(leftM, e)
		case e > m[base]:
			rightM = append(rightM, e)
		}
	}
	return append(append(quickSort(leftM), m[base]), quickSort(rightM)...)
}

func main() {
	m := []int{1, 25, 2, 4, 18, 3, 67, 5, 9, 10, 33, 99, 6, 8, 3, 2, 12, 15, 17, 19, 98, 56, 77, 88, 43, 44, 45, 23, 21, 40, 60, 70, 20, 47, 54}
	// fmt.Println("Отсортированный список:", sorting(m))
	fmt.Println("Отсортированный список:", quickSort(m))
	// fmt.Println("Отсортированный список:", quickSortRand(m))
}
