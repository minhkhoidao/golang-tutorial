package main

import (
	"fmt"
	"math"
)

func forContinue() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		prevZ := z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		if math.Abs(prevZ-z) < 1e-6 {
			break
		}
	}
	return z
}

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	forContinue()
	fmt.Println("Sum:", sum)
	fmt.Println(Sqrt(2))
}
