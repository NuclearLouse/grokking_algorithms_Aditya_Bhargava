package main

import "fmt"

// * Нахождение суммы элементов
func sum(m []int) int {
	if len(m) == 0 {
		return 0
	}
	return m[0] + sum(m[1:])
}

// * Рекурсивный подсчет числа элементов
func quantity(m []int) int {
	if fmt.Sprintf("%v", m) == "[]" {
		return 0
	}
	return 1 + quantity(m[1:])
}

// * Наибольшее число в списке
func maxNum(m []int) int {
	if len(m) == 2 {
		if m[0] > m[1] {
			return m[0]
		}
		return m[1]
	}
	max := maxNum(m[1:])
	if m[0] > max {
		return m[0]
	}
	return max
}

func main() {
	m := []int{1, 25, 2, 4, 18, 3, 67, 5, 9, 10, 33, 99, 6, 8, 3, 2, 12, 15, 17, 19, 98, 56, 77, 88, 43, 44, 45, 23, 21, 40, 60, 70, 20, 47, 54}
	fmt.Println("Количество элементов:", quantity(m), "(", len(m), ")")
	fmt.Println("Сумма всех элементов:", sum(m))
	fmt.Println("Максимальный элемент в списке:", maxNum(m))
}

// * Значение O при выводе значения каждого элемента
// *
// * O(n)

// * Значение O при удвоении значения каждого элемента
// *
// * O(n)

// * Значение O при удвоении значения первого элемента
// *
// * O(1)

// * Массив состоит из элементов [1, 3, 5, 7]. Сеачала каждый умножается на 2, затем на 3 и тд
// *
// * O(n^2)
