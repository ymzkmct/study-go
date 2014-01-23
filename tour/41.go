package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    r := make(map[string] int)
    for _, v := range strings.Fields(s) {
        r[v]++
    }
    return r
}

func main() {
    wc.Test(WordCount)
}