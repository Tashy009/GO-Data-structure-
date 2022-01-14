package main

import "fmt"

func main() {
	fmt.Printf("the ans is %v", substring("bbbbbb"))
}
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func substring(n string) int {
	startidx := 0
	maxm := 0
	xmap := make(map[string]int)
	x := len(n)

	for i := 0; i < x; i++ {

		if val, ok := xmap[string(n[i])]; ok {

			startidx = Max(startidx, val+1)

		}

		maxm = Max(maxm, i-startidx+1)

		xmap[string(n[i])] = i
	}

	return maxm

}
