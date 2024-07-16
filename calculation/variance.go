package calc

import "math"

func Variance(nums []float64) float64 {
	lenNum := len(nums)
	var sum float64

	for _, val := range nums {
		sum += math.Pow(val-Average(nums), 2.0)
	}
	return sum / float64(lenNum)
}
