// +build darwin,cgo

package main

// #cgo CFLAGS: -DCGO -DDARWIN -O3 -march=native -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti
// #cgo CFLAGS: -mavx2
// #cgo LDFLAGS: -framework Accelerate
// #include "stats_darwin.h"
// #include <stdint.h>
import "C"

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
	C.variance_uint32_darwin(buffer, length, resultPointer)
	return res
}
