package slidingwindow

import (
	"math"
)

func MaxDifference(prices []int) int {
	minPrice := math.MaxInt
	maxProfit := 0
	for _, price := range prices {
		if minPrice > price {
			minPrice = price
		}
		if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return -1
}
