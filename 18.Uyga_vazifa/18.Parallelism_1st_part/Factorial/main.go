package main

import (
	"fmt"
)

func factorial(n int, ch chan int) {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	ch <- result
}

func FillSlice(numbers *[]int, length int) {
	for i := 0; i < length; i++ {
		var new int
		fmt.Printf("%v - son: ", i+1)
		fmt.Scanln(&new)
		*numbers = append(*numbers, new)
	}
	fmt.Println()
}

func main() {
	ch := make(chan int)
	var numbers []int
	var length int

	fmt.Print("Nechta son kiritasiz: ")
	fmt.Scanln(&length)

	FillSlice(&numbers, length)

	for _, n := range numbers {
		go factorial(n, ch)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Faktorial:", <-ch)
	}
	close(ch)
}
