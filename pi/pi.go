package pi

import (
	"fmt"
	"math/rand"
	"time"
)

// Probabilistically calculates Pi for a given n points
func CalcPi(n int) float64 {
	in := 0
	// random seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	start := time.Now()
	// Calculates N points
	for i := 0; i < n; i++ {
		x := r.Float64()
		y := r.Float64()

		// Checks to see if point is also in M
		if x*x+y*y < 1 {
			in++
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("n=%d took %s\n", n, elapsed)
	// 4M/N
	return 4 * float64(in) / float64(n)
}
