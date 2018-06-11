package main

import "fmt"

func getPrime(primes *[]int) {
	if len(*primes) != 0 {
		temp := make([]int, 0)
		temp = append(temp, *primes...)
		for i := temp[len(*primes)-1] + 1; ; i++ {
			isPrime := true
			for j := 2; j < i; j++ {
				if i%j == 0 {
					isPrime = false
				}
			}
			if isPrime {
				*primes = append(*primes, i)
				break
			}
		}
	} else {
		*primes = append(*primes, 2)
	}
}
func main() {
	primes := make([]int, 0)
	for i := 0; i < 10001; i++ {
		getPrime(&primes)
	}
	fmt.Println(primes[len(primes)-1])
}
