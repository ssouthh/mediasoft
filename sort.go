package main

import (
	"fmt"
	"sort"
)

func bubble(a []int) []int {
	b := append([]int{}, a...)
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b)-1-i; j++ {
			if b[j] > b[j+1] {
				b[j], b[j+1] = b[j+1], b[j]
			}
		}
	}
	return b
}

func main() {
	fmt.Println("sort tests:")
	cases := [][]int{
		{5,1,4,2,8},
		{0,-1,9,-5,2},
		{},
		{1},
		{3,3,2,1,0,-1},
	}

	okAll := true
	for i,c := range cases {
		g := bubble(c)
		w := append([]int{}, c...)
		sort.Ints(w)
		ok := len(g)==len(w)
		if ok {
			for k:=0;k<len(w);k++ {
				if g[k] != w[k] {
					ok = false
					break
				}
			}
		}
		if ok {
			fmt.Println("case", i+1, "ok")
		} else {
			okAll = false
			fmt.Println("case", i+1, "fail", g, "!=", w)
		}
	}
	if okAll { fmt.Println("sort pass") } else { fmt.Println("sort fail") }
}
