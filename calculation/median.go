package calc

import (
	"sort"
)

func Median(nums []float64) float64 {
	sort.Float64s(nums)
	lenNum := len(nums)
	var sum, median float64

	// median for even length of slice
	if lenNum%2 == 0 {
		index := lenNum / 2
		sum = nums[index] + nums[index-1]
		median = sum / 2
		// median for odd length of slice
	} else if lenNum%2 == 1 {
		index := lenNum / 2
		median = nums[index]
	}
	return median
}
