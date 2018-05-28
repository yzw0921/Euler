package main

import "fmt"

func main() {
	a1 := 1
	a2 := 2
	a3 := 0
	fi := []int{a1, a2}
	fmt.Println(a1)
	fmt.Println(a2)
	for {
		a3 = a1 + a2
		if a3 <= 4000000 {
			fmt.Println(a3)
			fi = append(fi, a3)
			a1, a2 = a2, a3
		} else {
			break
		}
	}
	sum := 0
	fmt.Println(fi)
	for _, f := range fi {
		if f%2 == 0 {
			sum += f
		}
	}

	fmt.Println(sum)
}
