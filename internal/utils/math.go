package utils

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

// Computes the minimum and maximum value of a slice of integers
func GetExtrema(values []int) (int, int) {
	min := MaxInt
	max := MinInt
	for i := 0; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
		if values[i] < min {
			min = values[i]
		}
	}
	return min, max
}

// Determines which of two values is smallest
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Computes the sum of the first n positive integers
// https://math.stackexchange.com/questions/2260/proof-1234-cdotsn-fracn-timesn12
func Summorial(n int) int {
	return (n * (n + 1)) / 2
}
