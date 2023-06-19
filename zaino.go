package main

import (
	"fmt"
)

type oggetto struct {
	p int
	v int
}

func main() {
	o1 := oggetto{3, 3}
	o2 := oggetto{5, 2}
	o3 := oggetto{7, 8}

	s1 := oggetto{3, 3}
	s2 := oggetto{2, 4}
	s3 := oggetto{5, 4}
	s4 := oggetto{2, 1}
	s5 := oggetto{6, 7}

	var P int = 10
	var P2 int = 10

	var ogg []oggetto = []oggetto{o1, o2, o3}
	var ogg2 []oggetto = []oggetto{s1, s2, s3, s4, s5}

	fmt.Println("=^.^=")
	fmt.Println(valMax(ogg, P))
	fmt.Println("M: ", valMaxNoInfinitiOgg(ogg2, P2))
}

func valMaxNoInfinitiOgg(ogg []oggetto, P int) [][]int {
	var m [][]int = make([][]int, len(ogg)+1)
	for i, c := range m {
		c = make([]int, P+1)
		m[i] = c
	}

	//RIEMPIAMO PER COLONNE
	for c := 1; c < len(ogg)+1; c++ {
		for r := 1; r < len(m[c]); r++ {
			if ogg[c-1].p <= r { //POTREI PRENDERE OGG
				m[c][r] = max(m[c-1][r-ogg[c-1].p]+ogg[c-1].v, m[c-1][r])
			} else {
				m[c][r] = m[c-1][r] //NON PRENDO OGG, TROPPO PESANTE
			}
		}

	}
	return m
}

func valMax(ogg []oggetto, P int) int {
	var a []int = make([]int, P+1)

	for i := 0; i < P+1; i++ {
		for _, o := range ogg {
			if i+o.p >= P+1 {
				continue
			}
			a[i+o.p] = max(a[i+o.p], a[i]+o.v)
		}

	}
	return a[P]
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
