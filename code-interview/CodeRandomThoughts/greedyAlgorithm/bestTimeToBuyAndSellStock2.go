package greedyalgorithm

// 122. Best Time to Buy and Sell Stock II

// You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
// On each day, you may decide to buy and/or sell the stock.
// You can only hold at most one share of the stock at any time.
// However, you can sell and buy the stock multiple times on the same day,
// ensuring you never hold more than one share of the stock.
// Find and return the maximum profit you can achieve.

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
