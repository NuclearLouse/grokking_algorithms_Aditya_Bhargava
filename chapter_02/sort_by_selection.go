package main

import "fmt"

func remove(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func selectionSort(arr []int) []int {
	var newArr []int
	// так как я буду удалять из исходного среза элементы,
	// нужна дополнительная переменная чтобы пройтись по всей длине среза.
	larr := len(arr)
	for i := 0; i < larr; i++ {
		// нахожу индекс минимального элемента
		smallest := findSmallest(arr)
		// вставляю этот элемент в новый массив
		newArr = append(newArr, arr[smallest])
		// исключаю из массива этот элемент
		arr = remove(arr, smallest)
	}
	return newArr
}

func main() {
	myArr := []int{5, 3, 6, 2, 10}
	fmt.Println(selectionSort(myArr))
}
