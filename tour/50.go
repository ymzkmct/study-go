package main
import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Say() {
    fmt.Println(v)
}

func main() {
    v := &Vertex{3, 4}
    v.Say()
    fmt.Println(v.Abs())
}