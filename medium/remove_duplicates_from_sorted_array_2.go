package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

    k:=2
    for i:=2;i<len(nums);i++{
        if nums[i] != nums[k-2]{
           nums[k] = nums[i]
           k++
        }
    }
    return k
}

func main() {
    nums := []int{1, 1, 2, 2, 2, 3}
    newLength := removeDuplicates(nums)
    fmt.Println("New length:", newLength)
    fmt.Println("Modified array:", nums[:newLength])
}