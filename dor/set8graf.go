package main

import (
	"fmt"
)

// --------------------------------------------------
type vertice struct {
	valore []string
	chiave string
}

type grafo struct {
	vertici   map[string]*vertice
	adiacenti map[string][]*vertice
}

func main() {
	fmt.Println("=^.^=")

	var v []*vertice = make([]*vertice, 0)

	h1 := []string{"h1", "h2", "h3", "h4", "h5", "h6"}
	h2 := []string{"h4"}
	h3 := []string{"h5", "h6"}

	v = append(v, &vertice{h1, "a"})
	v = append(v, &vertice{h2, "b"})
	v = append(v, &vertice{h3, "c"})

	l1 := []string{"a", "b"}
	l2 := []string{"b", "a"}
	l3 := []string{"a", "c"}
	l := [][]string{l1, l2, l3}

	g := *popolaGrafo(v, l)

	for _, v := range g.vertici {
		stampaVertice(*v)
		fmt.Println("")
	}
	fmt.Println("=^.^=")
	StampaGrafo(g)

	Hobbyseguiti("a", g)

}

func graphNew(n int) *grafo {
	var vert map[string]*vertice = make(map[string]*vertice, n)
	var add map[string][]*vertice = make(map[string][]*vertice)

	return &grafo{vert, add}
}

func popolaGrafo(v []*vertice, c [][]string) *grafo {
	g := *graphNew(len(v))

	for _, e := range v {
		g.vertici[(*e).chiave] = e
	}

	for _, s := range c {
		g.adiacenti[s[0]] = append(g.adiacenti[s[0]], g.vertici[s[1]])
	}
	return &g
}

func stampaVertice(v vertice) {
	fmt.Print("vert{", v.chiave, " ,hobbies: ", v.valore, "} ")
}

func StampaGrafo(g grafo) {
	for _, v := range g.vertici {
		stampaVertice(*v)
		fmt.Print("segue:")
		for _, ve := range g.adiacenti[v.chiave] {
			fmt.Print(" ", ve.chiave, " ")
		}
		fmt.Print("\n")
	}

}

/*
Scrivete una funzione che data una stringa A stampi gli hobby dell’utente
di nome A e l’elenco di tutti gli hobby delle persone seguite da A.
*/
func HobbySeguiti(a string, g grafo) {
	fmt.Print("hobby di ", a, ": ", g.vertici[a].valore, ", ")
	fmt.Print("hobby di seguiti:")
	for _, s := range g.adiacenti[a] {
		fmt.Print(g.vertici[s.chiave].valore, " ")
	}
}

/*Scrivete una funzione che data una stringa A stampi gli hobby dell’utente 
di nome A e l’elenco di tutti gli hobby delle persone che seguono A*/

func HobbyCheSeguono(a string, g grafo) {
	fmt.Print("hobby di ", a, ": ", g.vertici[a].valore, ", ")
	fmt.Print("hobby di persono che seguono:")
	
	//DA SCRIVERE, NN HO SBATTI ORA 
}

/*
 Quale è la più complessa tra le ultime 2 operazioni? 
 stampare elenco di tutti gli hobby delle persone che seguono A
*/  
