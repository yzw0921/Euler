package main

import (
	"fmt"
	"math"
)

//传引用很重要！！！
func prime_factor(x int, pfa *[]int) {
	isPrimeFactor := false
	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			fmt.Println(i)
			isPrimeFactor = true
			*pfa = append(*pfa, i)
			prime_factor(x/i, pfa)
			break
		}
	}
	fmt.Println(x, isPrimeFactor)
	if isPrimeFactor == false {
		*pfa = append(*pfa, x)
	}
}

func main() {
	var prime []int
	prime_factor(600851475143, &prime)
	fmt.Println(prime)
}
