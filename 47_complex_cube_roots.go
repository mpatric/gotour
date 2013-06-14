package main

import (
    "fmt"
    "math/cmplx"
)

func Cbrt(x complex128) complex128 {
    z := x
    for {
        next_z := z - (((z * z * z) - x) / (3 * z * z))
        if cmplx.Abs(next_z - z) < 0.00001 {
            return next_z
        }
        z = next_z
    }
}

func main() {
    fmt.Printf("Mine: %f\n", Cbrt(2))
    fmt.Printf("Math: %f\n", cmplx.Pow(2, (1.0 / 3)))
    
    fmt.Printf("Mine: %f\n", Cbrt(-2))
    fmt.Printf("Math: %f\n", cmplx.Pow(-2, (1.0 / 3)))
}