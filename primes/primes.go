package primes

import (
	"math"
	"math/rand"
	"time"
)

// tests the primatliy of a given number, prime, against k random numbers between 0 and sqrt(prime)
func TestPrime(k int, prime int32) (bool, int32) {
	// new random seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Uncomment following lines for seeing the time of the operation
	// start := time.Now()
	// defer fmt.Printf("Time to check if %d is prime with %d tests: %s\n", prime, k, time.Since(start))
	for i := 0; i < k; i++ {

		// using floored sqrt(prime) = n, since you only need to check from (1, sqrt(prime)) to see if it's prime or not
		n := r.Int31n(int32(math.Sqrt(float64(prime))))

		// if divisible by something other than itself, not prime!
		if n != 0 && n != 1 && prime%n == 0 {
			return false, n
		}
	}

	// otherwise, prime (within the search space)
	return true, 0
}
