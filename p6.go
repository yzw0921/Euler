package main

import "fmt"

func main() {
	sumOfSq := 0
	sum := 0
	sqOfSum := 0
	for i := 1; i <= 100; i++ {
		sumOfSq += i * i
		sum += i
	}

	sqOfSum = sum * sum

	fmt.Println(sqOfSum - sumOfSq)
}
