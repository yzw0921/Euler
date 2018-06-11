package main

import (
	"fmt"
	"math"
)

func factorDivided(f *[]int, x int) {
	if x == 1 {
		return
	} else {
		for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
			if x%i == 0 {
				*f = append(*f, i)
				factorDivided(f, x/i)
				return
			}
		}
		*f = append(*f, x)
		return
	}
}

func addFact(fact *[]int, x int) {
	if len(*fact) != 0 {
		tempFact := make([]int, 0)
		tempXFact := make([]int, 0)
		xFact := make([]int, 0)
		tempFact = append(tempFact, *fact...)
		factorDivided(&xFact, x)
		tempXFact = append(tempXFact, xFact...)
		fmt.Printf("%d is divided into: %d\n", x, xFact)
		fmt.Println("The array fact is ", fact)
		//fmt.Println("The array tempFact is ", tempFact)
		n := 0
		for j := 0; j < len(xFact); j++ {
			for i := 0; i < len(tempFact); i++ {
				if tempFact[i] == xFact[j] {
					tempFact = append(tempFact[:i], tempFact[i+1:]...)
					fmt.Printf("tempFact is %d\n", tempFact)
					tempXFact = append(tempXFact[:j-n], tempXFact[j-n+1:]...)
					fmt.Printf("tempXFact is %d\n", tempXFact)
					n++
					break
				}
			}
		}
		if len(tempXFact) != 0 {
			*fact = append(*fact, tempXFact...)
		}

	} else {
		factorDivided(fact, x)
		fmt.Printf("%d is divided into: %d\n", x, *fact)
		fmt.Println("The array fact is []")
	}
}

func main() {
	fact := make([]int, 0)
	for i := 2; i <= 30; i++ {
		addFact(&fact, i)
	}
	product := 1
	for _, r := range fact {
		product *= r
	}
	fmt.Println(fact)
	fmt.Println(product)
}
