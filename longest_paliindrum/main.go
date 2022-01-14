package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	Lpalindrum("acaac")

}

func updatastring(n string) string {
	r := []rune(n)
	new := []rune{'#'}
	for _, val := range r {
		new = append(new, val)
		new = append(new, '#')

	}

	return string(new)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MaxA(n []int) int {
	max := math.MinInt32
	for _, val := range n {
		if val > max {
			max = val
		}
	}
	return max
}

func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func Lpalindrum(n string) {
	s := updatastring(n)
	LPS := make([]int, len(s))
	LPS[0] = 0
	C := 0
	R := 0
	for i := 0; i < len(s); i++ {
		mirror := 2*C - i
		if i < R {
			LPS[i] = Min(R-i, LPS[mirror])

		} else {

			LPS[i] = 0
		}
		for {
			if i+1+LPS[i] == len(s) {
				break
			} else if i-1-LPS[i] == -1 {
				break
			} else if s[i+1+LPS[i]] == s[i-1-LPS[i]] {

				LPS[i]++
			} else {
				break
			}

		}

		if i+LPS[i] > R {
			C = i
			R = i + LPS[i]
		}
	}

	r, c := MaxA(LPS), indexOf(MaxA(LPS), LPS)
	newS := s[c-r+1 : c+r]
	newS = strings.Replace(newS, "#", "", -1)

	fmt.Println(newS)

}
