package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threesum2(nums))
}

func threeSum(nums []int) [][]int {

	sort.Ints(nums)
	n := len(nums)
	output := [][]int{}

	for i := 0; i < n-1; i++ {

		l := i + 1
		r := n - 1
		x := nums[i]
		for l < r {

			if (x + nums[l] + nums[r]) == 0 {

				output = append(output, []int{x, nums[l], nums[r]})

				l += 1
				r -= 1
			} else if x+nums[l]+nums[r] < 0 {
				l += 1
			} else {
				r -= 1
			}

		}
	}
	return output

}

type List []int

func threesum2(list []int) []List {
	var ansList []List
	lenth := len(list)
	sort.Ints(list)
	for i := 0; i < lenth-2; i++ {

		low, high := i+1, lenth-1
		sum := 0 - list[i]
		for {
			if low >= high {
				break
			}

			switch {
			case list[low]+list[high] == sum:
				var ans List = []int{list[i], list[low], list[high]}
				fmt.Println(ans)
				ansList = append(ansList, ans)
				low++
				high = high - 1
			case list[low]+list[high] < sum:
				low++
			case list[low]+list[high] > sum:
				high = high - 1
			}

		}
	}
	return ansList
}
