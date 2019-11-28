// +build cgo

// #cgo CFLAGS: -O3 -march=native -mstackrealign -mllvm -inline-threshold=1000 -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti
// #cgo CFLAGS: -mavx2
package main

// #include "stats_blocking.h"
import "C"

// CalculateVarianceLoopUnrolled gets the variance of the data set
func CalculateVarianceLoopUnrolled(values []uint32) float64 {
	if len(values) == 0 || len(values) == 1 {
		return 0
	}

	var (
		buffer = (*C.uint32_t)(&values[0])
		length = C.int(len(values))
		res    float64
	)

	resultPointer := (*C.double)(&res)
	C.variance_uint32_loopunrolled(buffer, length, resultPointer)
	return res
}
