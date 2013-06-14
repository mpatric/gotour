package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := x
    for {
        next_z := z - ((z * z - x) / (2 * z))
        if math.Abs(next_z - z) < 0.000001 {
        	return next_z
        }
        z = next_z
    }
}

func main() {
	fmt.Println("Math: %f", math.Sqrt(2))
    fmt.Println("Mine: %f", Sqrt(2))
}