package main

import "fmt"

func fibonacci() func() int {
    current := 0
    next := 1
    return func() int {
        current, next = next, next + current
        return current
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}