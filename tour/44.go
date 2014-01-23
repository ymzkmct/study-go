package main
import (
    "fmt"
)

func hoge(i int) int {
    if i == 0 {
        return 0
    } else if i == 1 {
        return 1
    } else if i > 1 {
        return hoge(i-1) + hoge(i-2)
    } else {
        return 0
    }
}

func fibonacci() func() int {
    count := 0
    return func() int {
        v := hoge(count)
        count++
        return v
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}