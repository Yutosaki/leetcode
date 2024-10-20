package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int)  {
    j:=0
    for i:=m;i<len(nums1);i++{
        nums1[i]=nums2[j]
        j++
    }
    sort.Ints(nums1)
}

func main() {
	// Test cases to validate the solution
	// Test case 1
	// Expected output: [1,2,2,3,5,6]
	nums1 := []int{1,2,3,0,0,0}
	m := 3
	nums2 := []int{2,5,6}
	n := 3
	merge(nums1, m, nums2, n)

	fmt.Println(nums1)
}