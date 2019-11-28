// +build !cgo

package main

func CalculateVariance(values []uint32) float64 {
	var sum uint32 = 0
	for _, value := range values {
		sum += value
	}

	mean := float64(sum) / float64(len(values))

	var accumulator float64
	for _, value := range values {
		diff := float64(value) - mean
		accumulator += diff * diff
	}

	return accumulator / float64(len(values)-1)
}
