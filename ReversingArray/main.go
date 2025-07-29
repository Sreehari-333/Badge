package main

import "fmt"

func main() {

	nums := []int{10, 20, 30, 40, 50}
	start := 0
	end := len(nums) - 1

	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}

	fmt.Println(nums)

}
