package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    img := make([][]uint8, dy)
    for y := 0; y < len(img); y++ {
        img[y] = make([]uint8, dx)
        for x:= 0; x < len(img[y]); x++ {
            img[y][x] = uint8(x^y)
        }
    }
    return img
}

func main() {
    pic.Show(Pic)
}