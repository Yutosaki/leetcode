package main

import "fmt"

func rotate(nums []int, k int) {
	tmp := nums[len(nums)-k:]
	nums = nums[:len(nums)-k]
	tmp = append(tmp, nums...)
	copy(nums, tmp) // nums = tmp としても、元の配列の参照が変わらないので、copy() で中身をコピーする
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Original array:", nums)
	rotate(nums, 3)
	fmt.Println("Rotated array:", nums)
}
