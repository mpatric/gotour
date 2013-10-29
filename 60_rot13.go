package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
    n, err = r.r.Read(p)
    for i := 0; i < n; i++ {
        ch := p[i]
        if ((ch >= 'a' && ch <= 'm') || (ch >= 'A' && ch <= 'M')) {
            p[i] = ch + 13
        } else if ((ch >= 'n' && ch <= 'z') || (ch >= 'N' && ch <= 'Z')) {
            p[i] = ch - 13
        }
    }
    return
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}