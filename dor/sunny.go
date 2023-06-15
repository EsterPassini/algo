package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// H: ad ogni svincolo sceglie sempre e comunque la galleria meno luminosa

/*
N = numero svincoli (1... N) H: harmony, S: sarah
M = numero gallerie (1... M)

l := luminosità galleria

Problema Sapendo che non esistono due gallerie egualmente luminose, bisogna determinare
se Harmony possa raggiungere  Sarah e, in caso affermativo, quante gallerie
le sono necessarie per arrivare.
*/

/*MODELLAZIONE
NODI: sono gli svincoli delle gallerie

ARCHI: sono le gallerie

CARATTERISTICHE GRAFO: non orientato, pesato(?)

PROBLEMA: esistenza del cammino tra H e S scegliend archi che rispettano la strategia
	gready di harmony (ad ogni nodo sceglie sempre l'arco con "peso" minore), SE ESISTE
	restituire il numero di archi che lo compongono


ALGORITMO:
	associo ad ogni nodo un flag che sarà false se non ho ancora visitato il nodo, true altrimenti,
	dichiaro p una variabile numerica intera (associata al numero di archi attraversati), la pongo = 0.
	parto da H,
	1. scelgo l'arco con peso minore (nodo in cui sono, y) mi sposto in y, aggiorno p= p+1.
		se y è S:  ESISTE IL CAMMINO, restituisco p
		se il flag di y era true:   NON ESISTE IL CAMMINO
		else: aggiorno il flag di y a true e riparto da 1
*/

/*rappresento grafo con una slice di slice:
la slice associa ad ogni indice l'indice di un nodo, e contiene la slice
che include gli indici dei nodi a lui collegati, poiche i nodi partono da 1
e gli indici da 0, per passare da nodo a indice faremo (indice = nodo-1).

inoltre memorizzerò in 2 variabili intere il nodo di partenza H e quello di arrivo S.
utilizzerò una slice di booleani di lunghezza uguale alla slice con i nodi,
che segnerà i nodi già visitati e quelli non, come descritto dall'algoritmo.
*/

/*type Grafo struct {
	listaAdd [][]int
	pesi []int
	visitati []bool
	inizio   int
	fine     int
}

/*func cammino() (p int) {

}*/

/*func leggiGrafo() (g Grafo) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		for i := 0; ; i++ {
			if i == 0 {
				campi := strings.Split(s.Text(), " ")
				H, _ := strconv.Atoi(campi[2])
				S, _ := strconv.Atoi(campi[3])
				g.inizio = H - 1
				g.fine = S - 1
				N, _ := strconv.Atoi(campi[0]) //num nodi
				g.listaAdd = make([][]int, N)
				g.visitati = make([]bool, N)
				//M, _ := strconv.Atoi(campi[1]) //num archi
			} else{
				campi := strings.Split(s.Text(), " ")
				n1, _ := strconv.Atoi(campi[0])
				n2, _ := strconv.Atoi(campi[1])
				l, _ := strconv.Atoi(campi[2])
			}
		}
	}
	return
}*/

//--------------------------------------------------------------------------------------------------

/*MODELLAZIONE
NODI: sono gli svincoli delle gallerie

ARCHI: sono le gallerie con l'indice di luce minore (1 per nodo, 2 nodi possono anche avere la stessa)

CARATTERISTICHE GRAFO:  orientato, pesato(?), non per forza connesso

PROBLEMA: esistenza del cammino tra H e S scegliend archi che rispettano la strategia
	gready di harmony (ad ogni nodo sceglie sempre l'arco con "peso" minore), SE ESISTE
	restituire il numero di archi che lo compongono


ALGORITMO:
	associo ad ogni nodo un flag che sarà false se non ho ancora visitato il nodo, true altrimenti,
	dichiaro p una variabile numerica intera (associata al numero di archi attraversati), la pongo = 0.
	parto da H,
	1. scelgo l'arco collegato al mio nodo (nodo in cui sono, y) mi sposto in y, aggiorno p= p+1.
		se y è S:  ESISTE IL CAMMINO, restituisco p
		se il flag di y era true:   NON ESISTE IL CAMMINO
		else: aggiorno il flag di y a true e riparto da 1
*/

type ArcoMigliore [2]int

//pos0 : arco di arrivo
//pos1 : luminosità

type Grafo []ArcoMigliore

func leggiGrafo() (g Grafo, H int, S int) {
	s := bufio.NewScanner(os.Stdin)

	for i := 0; s.Scan(); i++ {
		if i == 0 {
			campi := strings.Split(s.Text(), " ")
			H, _ = strconv.Atoi(campi[2])
			S, _ = strconv.Atoi(campi[3])
			H = H - 1
			S = S - 1
			N, _ := strconv.Atoi(campi[0]) //num nodi
			g = make([]ArcoMigliore, N)

			for n, _ := range g {
				g[n][1] = 9999
			}

			//M, _ := strconv.Atoi(campi[1]) //num archi
		} else {
			campi := strings.Split(s.Text(), " ")
			n1, _ := strconv.Atoi(campi[0])
			n2, _ := strconv.Atoi(campi[1])
			l, _ := strconv.Atoi(campi[2])

			if g[n1-1][1] > l {
				a := ArcoMigliore{n2 - 1, l}
				g[n1-1] = a
			}
			if g[n2-1][1] > l {
				a := ArcoMigliore{n1 - 1, l}
				g[n2-1] = a
			}
		}
	}

	return
}

func cammino() (p int) {
	g, H, S := leggiGrafo()
	var visitati []bool = make([]bool, len(g))

	n := H

	for {
		if n == S {
			return p
		}
		if visitati[n] {
			return -1
		}
		visitati[n] = true
		p++
		n = g[n][0]
	}
	return
}

func main() {
	fmt.Println("=^.^=")
	fmt.Println(cammino())
	fmt.Println("=^.^=")
}
