/* ou are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store. */

package main

import "fmt"

func main() {

	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxArea(height []int) int {
	i := 0
	r := len(height) - 1
	area := 0

	for i < r {
		area = Max(area, Min(height[i], height[r])*(r-i))
		if height[i] < height[r] {
			i++
		} else {
			r--
		}
	}

	return area
}
