package main

import (
	"fmt"
)

type grafoo map[string][]string

func main() {

	var g grafoo = make(map[string][]string)
	g["a"] = []string{"b", "c", "d"}
	g["b"] = []string{"a", "c"}
	g["c"] = []string{"d"}
	g["d"] = []string{"a", "b", "c"}

	fmt.Println("=^.^=")
	var aux = make(map[string]bool)
	aux["a"] = false
	aux["b"] = false
	aux["c"] = false
	aux["d"] = false

	dfs1(g, "grafo", aux)

}

func dfs1(g grafoo, v string, aux map[string]bool) {
	fmt.Println(v)
	aux[v] = true
	for _, v2 := range g[v] {
		if aux[v2] != true {
			dfs1(g, v2, aux)
		}
	}
}

func bfs1(g grafo, v string, aux map[string]bool) {
	coda := []string{v}
	aux[v] = true
	for len(coda) > 0 {
		v := coda[0]
		coda = coda[1:]
		fmt.Println("\t", v)
		for _, v2 := range g[v] {
			if !aux[v2] {
				coda = append(coda, v2)
				aux[v2] = true
			}
		}
	}
}
