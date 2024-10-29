package main

import "fmt"

func maxProfit(prices []int) int {
	maxProfit := 0
	buyingDay := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[buyingDay] > maxProfit {
			maxProfit = prices[i] - prices[buyingDay]
		} else if prices[i]-prices[buyingDay] < 0 {
			buyingDay = i
		}
	}
	return maxProfit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{1, 2}))
}
