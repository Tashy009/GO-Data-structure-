package main

import "fmt"

func main() {
	pascal(10)
}

func pascal(n int) {
	var arr []int = []int{1}
	var temp []int

	for i := 0; i < n; i++ {

		for j, _ := range arr {

			fmt.Printf(" %d", arr[j])
		}

		fmt.Println("")

		temp = append(temp, 1)

		s := len(arr)

		for j := 0; j < s-1; j++ {

			temp = append(temp, (arr[j] + arr[j+1]))

		}

		temp = append(temp, 1)

		arr = temp

		temp = []int{}

	}
}
