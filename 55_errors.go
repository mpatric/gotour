package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("Cannot sqrt a negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return math.NaN(), ErrNegativeSqrt(x)
    }
    z := x
    for {
        next_z := z - ((z * z - x) / (2 * z))
        if math.Abs(next_z - z) < 0.000001 {
            return next_z, nil
        }
        z = next_z
    }
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}