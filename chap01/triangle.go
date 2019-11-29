package main

import (
	"math"
)

func main() {
	n := 5
	a := []int{2, 3, 4, 5, 10}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				len := a[i] + a[j] + a[k]
				ma = max()
			}
		}
	}
}
