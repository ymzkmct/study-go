package main
import (
    "fmt"
    "net/http"
)

type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}
func (s Struct) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {

    fmt.Fprintf(w, s.Greeting)
    fmt.Fprintf(w, s.Punct)
    fmt.Fprintf(w, s.Who)
}
func (s String) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {

    fmt.Fprintf(w, string(s))

}
func main() {
    http.Handle("/string", String("I'm a frayed knot."))
    http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
    http.ListenAndServe("localhost:4000", nil)
}