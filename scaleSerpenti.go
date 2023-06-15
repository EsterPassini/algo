package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
NODI: caselle

ARCHI: collega 2 nodi se le caselle che modellano sono raggiungibili la seconda dalla prima
con un singolo lancio di dadi

CARATTERISTICHE GRAFO: non connesso, ogni arco ha un etichetta (tiro che mi permette quello spostamento), orientato

PROBLEMA 1: numero di archi di uno dei cammini minimi, e stampare uno dei cammini

ALGORITMO 1: algo gready
strategia: scelgo lancio, cioè arco che mi porta alla casella con il numero maggiore

parto da un nodo:
	1.guardo archi uscenti, scelgo quello che mi conduce al nodo con indice maggiore,
	aggiungo uno al contatore, aggiungo a slice l'etichetta del arco (= lancio che permette spostamento)
	mi sposto in quel nodo

	se l'indice è l'indice massimo (= ultima casella mi fermo) RETURN contatore, stampo la slice

	ripeto dal passo 1

*/

type Mossa [2]int

//mossa[0] = nodo a cui arrivo
//mossa[1] = etichetta di quell'arco (lancio)

type Tabellone []int

type GrafoS [][]Mossa

func leggiGrafoS() (g GrafoS, t Tabellone) {
	s := bufio.NewScanner(os.Stdin)

	for i := 0; s.Scan(); i++ {
		campi := strings.Split(s.Text(), " ")
		c0, _ := strconv.Atoi(campi[0])
		c1, _ := strconv.Atoi(campi[1])
		if i == 0 {
			t = make(Tabellone, c0*c1)
			g = make(GrafoS, c0*c1)
		} else {
			//serpente o scala
			t[c0-1] = c1 - 1
		}
	}

	for c := 0; c < len(t); c++ {
		var arrivo int
		if t[c] != 0 { //su quella casella non arriverò mai
			continue
		}
		for l := 1; l <= 6; l++ {
			if c+l >= 30 {
				break
			}
			if t[c+l] != 0 {
				arrivo = t[c+l]
			} else {
				arrivo = l + c
			}
			mossa := Mossa{arrivo, l}
			g[c] = append(g[c], mossa)
		}
	}

	return
}

func scaleSerpenti(g GrafoS, c int) (cont int, lanci []int) {
	cor := c - 1

	for {
		if cor >= len(g)-1 {
			return
		}
		cMax := g[cor][0][0]
		l := g[cor][0][1]

		for _, mossa := range g[cor] {
			if cMax < (mossa[0]) {
				cMax = mossa[0]
				l = mossa[1]
			}
		}
		cont = cont + 1
		lanci = append(lanci, l)
		cor = cMax
	}
	return

}

func main() {
	fmt.Println("=^.^=")
	g, _ := leggiGrafoS()
	/*for c, s := range g {
		fmt.Println(c, ":  ", s)
	}*/

	fmt.Println(scaleSerpenti(g, 1))
	fmt.Println("=^.^=")
}
