package main

import "fmt"

func main() {
	fmt.Println(reverse_string("tashy"))
	fmt.Println(vowel_count("tashy"))
}

func reverse_string(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func vowel_count(s string) int {
	sum := 0
	for _, v := range s {
		switch v {
		case 'a', 'e', 'i', 'o', 'u':
			sum++
		}
	}
	return sum
}
