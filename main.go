package main

import (
	"fmt"
	"math"

	"github.com/timlucca/algo2/montecarlo"
	"github.com/timlucca/algo2/pi"
	"github.com/timlucca/algo2/primes"
	"github.com/timlucca/algo2/queens"
	"github.com/timlucca/algo2/randsearch"
)

func main() {

	// Pi approximation
	fmt.Printf("For n=1000, Pi is approximately: %f\n\n", pi.CalcPi(1000))
	fmt.Printf("For n=10000, Pi is approximately: %f\n\n", pi.CalcPi(10000))
	fmt.Printf("For n=100000, Pi is approximately: %f\n\n", pi.CalcPi(100000))
	fmt.Printf("For n=1000000, Pi is approximately: %f\n\n", pi.CalcPi(1000000))
	//fmt.Printf("For n=1000000000, Pi is approximately: %f\n", pi.CalcPi(1000000000))

	// Primality test
	numbers := []int32{10000, 3333333, 123456789, 2352142, 957298375, 1889 * 23471}
	for _, p := range numbers {
		fmt.Printf("\nTesting if %d is prime\n", p)
		for i := 1; i < 5; i++ {
			k := int(math.Pow10(i))
			if t, n := primes.TestPrime(k, p); t == false {
				fmt.Printf("%d is not prime! Divisible by %d (tested against %d values)\n", p, n, k)
			} else {
				fmt.Printf("%d is probably prime (tested against %d values)\n", p, k)
			}
		}
	}

	// Random search
	fmt.Println("\nNow testing random search...")
	total := 0
	for i := 0; i < 100; i++ {
		if found, j := randsearch.RandomSearch(); found {
			total = total + j
		} else {
			total = total + j
		}
	}
	fmt.Printf("The average tries was %d\n\n", (total / 100))

	// Monte Carlo integration
	monte, t := montecarlo.Darts(func(x float64) float64 {
		return math.Pow(math.E, x*math.Sin(x))
	}, -math.Pi, math.Pi, 7, 1000)

	fmt.Println("Dart throw")
	fmt.Printf("The area of e^(x*sin(x)) from (-π, π): %f\nTook: %s\n", monte, t)

	// Random Mean
	randm, t := montecarlo.Mean(func(x float64) float64 {
		return math.Pow(math.E, x*math.Sin(x))
	}, -math.Pi, math.Pi, 1000)

	fmt.Println("\nRandom mean")
	fmt.Printf("The mean of e^(x*sin(x)) from (-π, π): %f\nTook: %s\n", randm, t)

	// Trapezoidal Method
	trap, t := montecarlo.Trap(func(x float64) float64 {
		return math.Pow(math.E, x*math.Sin(x))
	}, -math.Pi, math.Pi, 1000)

	fmt.Println("\nTrapeziodal method")
	fmt.Printf("The area of e^(x*sin(x)) from (-π, π): %f\nTook: %s\n", trap, t)

	// Gaussian Quadrature
	gauss, t := montecarlo.GQuad(func(x float64) float64 {
		return math.Pow(math.E, x*math.Sin(x))
	}, -math.Pi, math.Pi)
	fmt.Println("\nGaussian Quadrature")
	fmt.Printf("The area of e^(x*sin(x)) from (-π, π): %f\nTook: %s\n", gauss, t)

	queens.RQueens(5)
}
