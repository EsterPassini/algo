package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Buono struct {
	costo int
	val   int
}

func main() {
	var budget int = 100

	fmt.Println(valoreMassimoBuoni(budget))
}

func leggiBuoni() []Buono {
	var b []Buono
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {

		st := s.Text()
		if st == "" {
			break
		}
		campi := strings.Split(st, " ")
		v, _ := strconv.Atoi(campi[0])
		c, _ := strconv.Atoi(campi[1])
		b = append(b, Buono{c, v})
	}

	return b
}

func valoreMassimoBuoni(bud int) int {
	var b []Buono
	var v []int = make([]int, bud+1, bud+1)

	b = leggiBuoni()

	for i := 0; i < len(v); i = i + 2 {
		for _, bu := range b {

			if (bu.costo + i) >= bud+1 {
				continue
			}

			v[i+(bu.costo)] = maxf(v[i+(bu.costo)], v[i]+bu.val)
		}
	}
	//fmt.Println(v)
	return v[bud]
}

func maxf(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
