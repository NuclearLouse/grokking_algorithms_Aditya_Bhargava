package main

// Из среза будет возвращать индекс
func returnFromSlice(l []int, e int) int {
	for i := range l {
		if l[i] == e {
			return i
		}
	}
	return 0
}

// Из карты будет возвращать значение
func returnFromMap(m map[int]int, k int) int {
	if val, ok := m[k]; ok {
		return val
	}
	return 0
}

// func main() {
// 	// 	l := []int{1, 25, 2, 4, 18, 3, 67, 5, 9, 10, 33, 99, 6, 8, 3, 2, 12, 15, 17, 19, 98, 56, 77, 88, 43, 44, 45, 23, 21, 40, 60, 70, 20, 47, 54}
// 	// 	fmt.Println(returnFromSlice(l, 77))
// 	// 	m := map[int]int{1: 1, 2: 25, 3: 2, 4: 4, 5: 18, 6: 3, 7: 67, 8: 5, 9: 9, 10: 10, 11: 33, 12: 99, 13: 6, 14: 8, 15: 3, 16: 2, 17: 12, 18: 15, 19: 17, 20: 19, 21: 98, 22: 56, 23: 77, 24: 88, 25: 43, 26: 44, 27: 45, 28: 23, 29: 21, 30: 40, 31: 60, 32: 70, 33: 20, 34: 47, 35: 54}
// 	// 	fmt.Println(returnFromMap(m, 23))
// 	set := setMap{tMap: make(map[string]int)}
// 	for i := 0; i < 102; i++ {
// 		set.addMap(strconv.Itoa(i), i)
// 	}
// 	fmt.Println(set.tMap)
// 	fmt.Println(returnFromMap(set.tMap, "1023"))
// }
