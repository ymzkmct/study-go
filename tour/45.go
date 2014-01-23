package main
import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println("Go runs on ")
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X.")
        fallthrough /* breakしたくない場合にfallthroughを入れる */
    case "linux":
        fmt.Println("Linux.")
    default:
        fmt.Println("%s.", os)
    }
    /* 自動的にbreakする */
}