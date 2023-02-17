package utils

func WrapRange(value, start, end int) int {
	if end < start {
		start, end = end, start
	}
	length := end - start
	value = (value - start) % length
	if value < 0 {
		value += length
	}
	return value + start
}
