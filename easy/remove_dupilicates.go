package main

func removeDuplicates(nums []int) int {
	counter := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[counter] = nums[i]
			counter++
		}
	}

	return counter
}

func main() {
	// Test cases to validate the solution
	// Test case 1
	// Expected output: 5
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	println(removeDuplicates(nums))
}
