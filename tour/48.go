package main

import (
    "fmt"
    "math/cmplx"
)
func Cbrt(x complex128) complex128 {
    var z, e = complex128(1.0), complex128(0.0)
    for {
        z = z - (z*z*z-x) / (3*z*z)
        if cmplx.Abs(z - e) < 1.0e-16 {
            break
        }
        e = z
    }
    return z
}
func main() {
    fmt.Println(cmplx.Pow(2, 1.0 / 3.0))
    fmt.Println(Cbrt(2))
}
