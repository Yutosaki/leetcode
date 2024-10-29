package main

func maxProfit(prices []int) int {
	max := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1]-prices[i] > 0 {
			max += prices[i+1] - prices[i]
		}
	}
	return max
}

func main() {
	println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	println(maxProfit([]int{7, 6, 4, 3, 1}))
	println(maxProfit([]int{1, 2}))
}
