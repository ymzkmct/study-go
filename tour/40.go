package main
import "fmt"

func main() {
    m := make(map[string]int)

    m["Answer"] = 42
    fmt.Println("The Value:", m["Answer"])

    m["Answer"] = 48
    fmt.Println("The Value:", m["Answer"])

    delete(m, "Answer")
    fmt.Println("The Value:", m["Answer"])

    var v, ok = m["Answer"]
    fmt.Println("The Value:", v, "Present?", ok)

    m["Answer"] = 48
    v, ok = m["Answer"]
    fmt.Println("The Value:", v, "Present?", ok) /* "ok" maybe returns "yes" */
}