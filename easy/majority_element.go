package main

func majorityElement(nums []int) int {
	stack := make([]int, 1)
	stack[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if len(stack) == 0 {
			stack = append(stack, nums[i])
		}else if stack[len(stack)-1] == nums[i] {
			stack = append(stack, nums[i])
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return stack[0]
}

func main() {
	nums := []int{2,2,1,1,1,2,2}
	majorityElement := majorityElement(nums)
	println("Majority element:", majorityElement)
}
