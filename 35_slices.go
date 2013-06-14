package main

import "code.google.com/p/go-tour/pic"

func f(x, y int) int {
    //return x * y
    //return x ^ y
    return (x + y) / 2
}

func Pic(dx, dy int) [][]uint8 {
    pic := make([][]uint8, dy)
    for y := 0; y < dy; y++ {
        pic[y] = make([]uint8, dx)
        for x := 0; x < dx; x++ {
            pic[y][x] = (uint8)(f(x, y))
        }
    }
    return pic
}

func main() {
    pic.Show(Pic)
}