package main

import "fmt"

func countdown(n int) {
	fmt.Println(n)
	if n <= 0 {
		return
	}
	countdown(n - 1)
}

func main() {
	countdown(10)
}
