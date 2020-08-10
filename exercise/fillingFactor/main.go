package main

import (
	"fmt"
	"os"
)

func main() {
	var primes [52]int
	var j uint8
	for i := 0; i < 1000; i++ {
		if isPrime(i) {
			if j > 51 {
				break
			}
			primes[j] = i
			j++
		}
	}

	var sum int
	var index uint
	position := make(map[int]int)
	for i := 1; i < len(os.Args); i++ {
		for _, s := range os.Args[i] {
			if byte(s) >= 'A' && byte(s) <= 'Z' {
				sum += primes[byte(s)-'A']
			}

			if byte(s) >= 'a' && byte(s) <= 'z' {
				sum += primes[byte(s)-'a'+26]
			}
		}
		_, ok := position[sum%10]
		if ok == false {
			index++
		}
		position[sum%10]++
	}
	fmt.Println("index:", index)
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}

	for i := 3; i < n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
