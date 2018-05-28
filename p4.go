package main

import (
	"fmt"
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
		break
	}
}
