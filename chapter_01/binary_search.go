package main

import "fmt"

type pos int

func (p pos) String() string {
	if p < 0 {
		return "None"
	}
	return fmt.Sprintf("%d", p)
}

func main() {
	myList := []int{1, 3, 5, 7, 9}
	position := pos(binarySearch(myList, 7))

	fmt.Println("Position:", position)
}

func binarySearch(l []int, num int) int {
	low := 0
	high := len(l) - 1
	for low <= high {
		mid := (low + high) / 2
		switch {
		case l[mid] == num:
			return mid
		case l[mid] > num:
			//выбрасываем из дальнейшего поиска все числа выше среднего
			// путем сдвига высшей точки поиска вниз до среднего значения
			high = mid - 1
		case l[mid] < num:
			//выбрасываем из дальнейшего поиска все числа меньше среднего
			//путем сдвига низшей точки поиска вверх до среднего значения
			low = mid + 1
		}
	}
	return -1
}
