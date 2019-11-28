// +build !darwin,cgo

package main

// #cgo CFLAGS: -DCGO -std=c99 -O2 -pedantic -ffast-math -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti
// #cgo CFLAGS: -mavx2
// #cgo amd64 386 arm armbe arm64 arm64be CFLAGS: -march=native
// #include "stats_cgo.h"
import "C"

// Note the -std=gnu99. Using -std=c99 will not work.

// CalculateVariance gets the variance of the data set
func CalculateVariance(values []uint32) float64 {
	if len(values) == 0 || len(values) == 1 {
		return 0
	}

	var (
		buffer = (*C.uint32_t)(&values[0])
		length = C.int(len(values))
		res    float64
	)

	resultPointer := (*C.double)(&res)
	C.variance_uint32(buffer, length, resultPointer)
	return res
}
