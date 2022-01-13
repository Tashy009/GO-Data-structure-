package main

import "fmt"

func main() {
	fmt.Println("this is the one: ", list(100))
}

func list(n int) []int {

	arr := []int{}
	for i := 2; i < n; i++ {
		if isPrime(i) {
			arr = append(arr, i)
		}
	}
	return arr
}

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
