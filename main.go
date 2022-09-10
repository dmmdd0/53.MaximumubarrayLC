package main

import "fmt"

func main() {

	arr := []int{-1}
	arr = []int{1}
	arr = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	arr = []int{5, 4, -1, 7, 8}
	arr = []int{-1}
	fmt.Println(maxSubArray(arr))

}

func maxSubArray(nums []int) int {

	max := nums[0]
	summ := 0
	for _, num := range nums {
		summ += num
		if summ < num {
			summ = num
		}
		if summ > max {
			max = summ
		}
	}
	return max
}
