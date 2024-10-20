package main

import "fmt"

func removeElement(nums []int, val int) int {
	counter := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[counter] = nums[i]
			counter++
		}
	}
	nums = nums[:counter]
	fmt.Println(nums)
	return counter
}

func main() {
	// Test cases to validate the solution
	// Test case 1
	// Expected output: 5
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(nums, val))
	fmt.Println(nums)
}
