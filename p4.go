package main

import (
	"fmt"
	"strconv"
)

func main() {
	palindromic := make([]int, 0)
	for i := 999999; i > 0; i-- {
		temp := strconv.Itoa(i)
		rune := make([]rune, len(temp))
		n := 0
		for _, r := range temp {
			rune[n] = r
			n++
		}
		for j := 0; j < n/2; j++ {
			rune[j], rune[n-1-j] = rune[n-1-j], rune[j]
		}
		r_temp := string(rune)
		if temp == r_temp {
			//fmt.Println("result:", temp)
			palindromic = append(palindromic, i)
		}
	}
	fmt.Println(len(palindromic))
	for _, p := range palindromic {
		for j := 100; j <= 999; j++ {
			if p%j == 0 {
				fact1 := strconv.Itoa(p / j)
				if len(fact1) == 3 {
					fmt.Printf("%d is product by %d, %d \n", p, p/j, j)
				}
			}
		}
	}
}
