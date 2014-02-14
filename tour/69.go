package main

import "code.google.com/p/go-tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    _walk(t, ch)
    close(ch)
}
// walk recursive
func _walk(t *tree.Tree, ch chan int) {
    if t == nil {return}
    _walk(t.Left, ch)
    ch <- t.Value
    _walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    ch1 := make(chan int)
    go Walk(t1, ch1)

    ch2 := make(chan int)
    go Walk(t2, ch2)

    for i := range ch1 {
        if i != <- ch2 {
            return false
        }
    }
    return true
}

func main() {
    ch := make(chan int)
    go Walk(tree.New(2), ch)
    for i := range ch {
        fmt.Print(i, " ")
    }
    fmt.Println("")

    test(true, Same(tree.New(1), tree.New(1)))
    test(false, Same(tree.New(1), tree.New(2)))
}

func test(expected bool, actual bool) {
    if (expected == actual) {
        fmt.Println("success!")
    } else {
        fmt.Println("failure")
    }
}