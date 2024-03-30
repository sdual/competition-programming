package chap01

import (
	"fmt"
	"strconv"
	"strings"
)

const maxN = 50

func triangle() {
	var n int
	fmt.Scan(&n)

	var m int
	fmt.Scan(&m)

	var str string
	fmt.Scan(&str)
	numStrs := strings.Split(str, ",")

	k := make([]int, len(numStrs))
	for i, numStr := range numStrs {
		num, _ := strconv.Atoi(numStr)
		k[i] = num
	}

	f := false
	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			for c := 0; c < n; c++ {
				for d := 0; d < n; d++ {
					if k[a]+k[b]+k[c]+k[d] == m {
						f = true
					}
				}
			}
		}
	}

	if f {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func solve() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {

	}
}
