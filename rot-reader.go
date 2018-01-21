package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	n, e := r.r.Read(b)
	for i, v := range b {
		b[i] = rot13(v)
	}
	return n, e
}

func rot13(b byte) byte {
	key := 'A' | (b & 0x20)
	if b & 0x40 == 0x40 {
		return key + ((b - key) + 13) % 26
	}
	return b
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
