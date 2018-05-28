package main

import (
	"fmt"
	"strconv"
)

func main() {
	a1 := 1
	for {
		if a1*11 <= 999 {
			a1++
		} else {
			a1--
			break
		}
	}
	fmt.Println("a1: ", a1*11)
	for i := 999; ; i-- {
		a := int64(a1 * 11 * i)
		b := strconv.

			fmt.Println(b)
	}
}
