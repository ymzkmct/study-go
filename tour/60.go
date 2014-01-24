package main

import (
    "code.google.com/p/go-tour/pic"
    "image"
    "image/color"
)

type Image struct{}

func (img Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, 150, 100)
}

func (img Image) At(x, y int) color.Color {
    return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
    m := Image{}
    pic.ShowImage(m)
}