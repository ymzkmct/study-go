package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

/*
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}
*/
func (v Vertex) Scale(f float64) Vertex {
    vv := v
    vv.X = v.X * f
    vv.Y = v.Y * f
    return vv
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    v := &Vertex{3, 4}
    vv := v.Scale(5)
    fmt.Println(vv.Abs())
}