package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	m := make([][]uint8, dy)
	for i := range m {
		m[i] = make([]uint8, dx)

		for j := range m[i] {
			m[i][j] = uint8(j^i)
		}
	}

	return m
}

func main() {
	pic.Show(Pic)
}
