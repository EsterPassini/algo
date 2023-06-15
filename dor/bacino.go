/*
https://adventofcode.com/2021/day/9
Giorno 9: Smoke Basin --- 2021

Queste grotte sembrano essere tubi di lava. Le parti sono ancora vulcanicamente attive;
Piccole bocche idrotermali rilasciano fumo nelle grotte che lentamente si deposita come pioggia.

Se riesci a modellare come il fumo scorre attraverso le grotte, potresti essere in grado di evitarlo.
Il sottomarino genera una mappa dell'altezza del pavimento
Ogni numero corrisponde all'altezza di una particolare posizione
primo obiettivo è trovare i punti più bassi -
le posizioni che sono più basse di qualsiasi posizione adiacente.
La maggior parte delle località ha quattro posizioni adiacenti (su, giù, sinistra e destra)

Il livello di rischio di un punto basso è 1 più la sua altezza.

Trova tutti i punti più bassi sulla tua mappa di altezza.
Qual è la somma dei livelli di rischio di tutti i punti bassi sulla tua mappa di altezza?

2199943210
3987894921
9856789892
8767896789
9899965678

punti bassi (4):
prima riga: primo 1 e lo 0
terza riga: 5
quinta riga(ultima): 5

*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("=^.^=")
}

func trovaPuntiBassi(m [][]int) []int {
	//caso base
	if len(m) == 1 && len(m[0]) == 1 {
		return []int{m[0][0]}
	}

	return nil
}
