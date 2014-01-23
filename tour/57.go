package main
import (
    "fmt"
    "net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP (
    w http.ResponseWriter,
    r *http.Request) {

    fmt.Fprintf(w, "hello!")
}

func main() {
    var h Hello
    http.ListenAndServe("localhost:4000", h)
}