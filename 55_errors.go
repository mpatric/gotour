package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("Cannot Sqrt negative number: %f", e)
}

func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, ErrNegativeSqrt(f)
    } else {
    	return math.Pow(f, 0.5), nil
    }
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}