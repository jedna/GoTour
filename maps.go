package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	w := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range w {
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

