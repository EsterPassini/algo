package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Interv struct {
	in int
	fi int
	v  int
}

func main() {
	scheduling()
}

// restituisce una slice di intervalli ordinati per orario di fine crescente
func leggiIntervalli() []Interv {
	var in map[int]Interv = make(map[int]Interv)
	var fine []int

	//input: intervalli, valore es (7 9 5) = intervallo 7-9 valore 5
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		s1 := s.Text()
		if s1 == "" {
			break
		}
		campi := strings.Split(s1, " ")
		a, _ := strconv.Atoi(campi[0])
		b, _ := strconv.Atoi(campi[1])
		c, _ := strconv.Atoi(campi[2])
		inter := Interv{a, b, c}
		in[b] = inter
		fine = append(fine, b)
	}
	sort.Ints(fine)
	var intervalli []Interv = make([]Interv, 0, len(fine)+1)
	intervalli = append(intervalli, Interv{0, 0, 0})
	for _, f := range fine {
		intervalli = append(intervalli, in[f])
	}
	return intervalli
}

func scheduling() {
	var p []int
	var v []int

	intervalli := leggiIntervalli()

	p = trovaP(intervalli)
	fmt.Println(p)
	v = make([]int, len(intervalli), len(intervalli))
	v[0] = 0

	for c, i := range intervalli {
		if c == 0 {
			continue
		}
		v[c] = maxx(v[c-1], i.v+v[p[c]])
	}
	fmt.Println(v)


	/*for i := len(v) - 1; i >= 0; {
		if v[i] == v[i-1] {
			i--
			continue
		} else {
			for t := i - 1; t >= 0; t-- {
				if v[i] == intervalli[i].v+v[p[t]] {
					fmt.Println("intervalloooo: ", intervalli[i])
				}
				i = t + 1
				break
			}
		}
		i--
	}*/

}

func trovaP(inte []Interv) (p []int) {
	p = make([]int, len(inte))
	p[0] = 0

	for i := len(inte) - 1; i >= 0; i-- {
		fmt.Println(i)
		inter := inte[i]
		for t := i - 1; t >= 0; t-- {
			if inter.in >= inte[t].fi {
				p[i] = t
				break
			}
		}
	}
	return p

}

func maxx(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
