// +build darwin,cgo

package main

import (
	"testing"
	"time"
)

func CalculateVarianceInGo(values []uint32) float64 {
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

func Test_CalculateVariance_4096_SpeedTest(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
		return
	}

	var values = make([]uint32, 4096, 4096)
	start := time.Now()
	for i := 0; i < 100_000; i++ {
		CalculateVarianceInGo(values)
	}

	duration := time.Since(start)
	t.Logf("Spent %s calling go code with vector of size 4096 64-bit floats", duration)
}

func Test_CalculateVarianceInGo_SpeedTest(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
		return
	}

	var values = make([]uint32, 512, 512)
	start := time.Now()
	for i := 0; i < 100_000; i++ {
		CalculateVarianceInGo(values)
	}

	duration := time.Since(start)
	t.Logf("Spent %s running calling go code with vector of size 512 64-bit floats", duration)
}

func Test_CalculateVarianceInGo_64_SpeedTest(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
		return
	}

	var values = make([]uint32, 64, 64)
	start := time.Now()
	for i := 0; i < 100_000; i++ {
		CalculateVarianceInGo(values)
	}

	duration := time.Since(start)
	t.Logf("Spent %s running calling go code with vector of size 64 64-bit floats", duration)
}

func Test_CalculateVariance_4096_SpeedTestOnDarwin(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
		return
	}

	var values = make([]uint32, 4096, 4096)
	start := time.Now()
	for i := 0; i < 100_000; i++ {
		CalculateVariance(values)
	}

	duration := time.Since(start)
	t.Logf(`Spent %s calling c code with vector of size 4096 64-bit floats. 
			 Utilizing Apple Accelerate Framwork. Using CGO Foreign Function Interface (FFI).`, duration)
}

func Test_CalculateVariance_SpeedTestOnDarwin(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
		return
	}

	var values = make([]uint32, 512, 512)
	start := time.Now()
	for i := 0; i < 100_000; i++ {
		CalculateVariance(values)
	}

	duration := time.Since(start)
	t.Logf(`Spent %s calling c code with vector of size 512 64-bit floats.
		     Utilizing Apple Accelerate Framwork. Using CGO Foreign Function Interface (FFI).`, duration)
}

func Test_CalculateVariance_64_SpeedTestOnDarwin(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
		return
	}

	var values = make([]uint32, 64, 64)
	start := time.Now()
	for i := 0; i < 100_000; i++ {
		CalculateVariance(values)
	}

	duration := time.Since(start)
	t.Logf(`Spent %s calling c code with vector of size 64 64-bit floats.
		     Utilizing Apple Accelerate Framwork. Using CGO Foreign Function Interface (FFI).`, duration)
}

func Test_CalculateVarianceInGo_InputIsOneOfFortySix(t *testing.T) {
	var values = []uint32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	result := CalculateVarianceInGo(values)

	// 0.021266540642722
	if result < 0.02173913043478260 {
		t.Fatalf("Variance is to small. got %.18f. Calling portable go function.", result)
	}

	if result > 0.02173913043478261 {
		t.Fatalf("Variance is to big. got %.18f. Calling portable go function.", result)
	}
}

func Test_CalculateVariance_InputIsOneOfFortySix(t *testing.T) {
	var values = []uint32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	result := CalculateVariance(values)

	// 0.021266540642722
	if result < 0.02173913043478260 {
		t.Fatalf("Variance is to small. Expected 0.021739130434782608. got %.18f. Utilizing Apple Accelerate Framework throught CGO.", result)
	}

	if result > 0.02173913043478261 {
		t.Fatalf("Variance is to big. Expected 0.021739130434782608. got %.18f. Utilizing Apple Accelerate Framework throught CGO.", result)
	}
}

func Test_CalculateVariance_InputHasEvenNumberOfValues(t *testing.T) {
	var values = []uint32{1, 2, 1, 2}

	result := CalculateVariance(values)

	if result < 0.33333 {
		t.Fatalf("Variance is to small. expected 0.333... got %f. Utilizing Apple Accelerate Framework throught CGO.", result)
	}

	if result > 0.33334 {
		t.Fatalf("Variance is to big. expected 0.333... got %f. Utilizing Apple Accelerate Framework throught CGO.", result)
	}
}

func Test_CalculateVarianceInGo_InputHasEvenNumberOfValues(t *testing.T) {
	var values = []uint32{1, 2, 1, 2}

	result := CalculateVarianceInGo(values)

	if result < 0.33333 {
		t.Fatalf("Variance is to small. expected 0.333... got %f. Calling portable go function.", result)
	}

	if result > 0.33334 {
		t.Fatalf("Variance is to big. expected 0.333... got %f. Calling portable go function.", result)
	}
}

func Test_CalculateVariance_InputHasOddNumberOfValues(t *testing.T) {
	var values = []uint32{1, 2, 1, 2, 1}
	result := CalculateVariance(values)

	if result < 0.3 {
		t.Fatalf("Variance is to small. got %f. Utilizing Apple Accelerate Framework throught CGO.", result)
	}

	if result > 0.3 {
		t.Fatalf("Variance is to big. got %f. Utilizing Apple Accelerate Framework throught CGO.", result)
	}
}
