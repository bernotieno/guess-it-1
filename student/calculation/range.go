package calc

import "math"

func Range(data []float64) (int, int) {
	start := len(data) - 4
	if start < 0 {
		start = 0
	}
	window := data[start:]
	mean := average(window)
	variance := variance(window)
	stddev := standarddev(variance)

	lower := mean - (1.8 * stddev)
	upper := mean + (1.8 * stddev)

	return int(math.Round(lower)), int(math.Round(upper))
}
