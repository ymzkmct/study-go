package main
import (
    "io"
    "os"
    "strings"
    //"fmt"
)

type rot13Reader struct {
    r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
    n, err = r.r.Read(p)

    // http://ja.wikipedia.org/wiki/ROT13
    for i := 0; i < n; i++ {
        if 65 <= int(p[i]) && int(p[i]) <= 77 {
            p[i] += 13
        } else if 78 <= int(p[i]) && int(p[i]) <= 89 {
            p[i] -= 13
        } else if 97 <= int(p[i]) && int(p[i]) <= 109 {
            p[i] += 13
        } else if 110 <= int(p[i]) && int(p[i]) < 122 {
            p[i] -= 13
        }
    }
    return
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}