package calc

import "math"

func Range(data []float64) (int, int) {
	mean := average(data)
	variance := variance(data)
	stddev := standarddev(variance)

	lower := mean - 1.5*stddev
	upper := mean + 1.5*stddev

	return int(math.Round(lower)), int(math.Round(upper))
}
