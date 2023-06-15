package main

import (
	"fmt"
)

type listNode struct {
	ele  int
	pros *listNode
}

type linkedList struct {
	head *listNode
}

type grafo1 struct {
	n         int // numero di vertici
	adiacenti []*linkedList
}

// --------------------------------------------------
type vertice struct {
	valore []string
	chiave string
}

type grafo struct {
	vertici   map[string]*vertice
	adiacenti map[string][]*vertice
}

func printList(l linkedList) {
	p := l.head

	for p != nil {

		fmt.Print(p.ele, " ")
		p = p.pros
	}
}

func newNodee(ele int) *listNode {
	return &listNode{ele, nil}
}

func addNewNode(ele int, l *linkedList) {
	nod := newNodee(ele)
	nod.pros = l.head
	l.head = nod
}

func searchList(n int, l linkedList) *listNode {
	var p *listNode = l.head

	for p != nil {
		if p.ele == n {
			return p
		}
		p = p.pros
	}
	return p
}

func main() {
	fmt.Println("=^.^=")

	/*es 1
	grafo := leggiGrafo()
	stampaGrafo(*grafo)
	fmt.Println(arcoDiretto(3, 0, *grafo))
	*/
	var v []*vertice = make([]*vertice, 3)
	v = append(v, &vertice{nil, "a"})
	v = append(v, &vertice{nil, "b"})
	v = append(v, &vertice{nil, "c"})
	fmt.Println(v)
	fmt.Println(*popolaGrafo(v))
}

func graphNew(n int) *grafo {
	var vert map[string]*vertice = make(map[string]*vertice, n)
	var add map[string][]*vertice = make(map[string][]*vertice)

	return &grafo{vert, add}
}

/*func popolaGrafo(v []*vertice) *grafo {
	g := *graphNew(len(v))

	for _, e := range v {
		g.vertici[(*e).chiave] = e
	}
	return &g
}*/

/*func nuovoGrafo(n int) *grafo {
	var g grafo
	g.n = n
	var a []*linkedList = make([]*linkedList, n, n)
	for i := 0; i < n; i++ {
		a[i] = &linkedList{nil}
	}
	g.adiacenti = a
	return &g
}

func leggiGrafo() *grafo {
	// formato atteso:  n che indica il numero di vertici seguito da una serie di coppie di indici, una coppia per riga
	var g *grafo
	scanner := bufio.NewScanner((os.Stdin))
	i := 0
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			break
		}
		//riga in stringaaa
		if i == 0 {
			n, _ := strconv.Atoi(s)
			g = nuovoGrafo(n)
		} else {
			vertici := strings.Split(s, " ")
			v0, _ := strconv.Atoi(vertici[0])
			v1, _ := strconv.Atoi(vertici[1])

			addNewNode(v0, g.adiacenti[v1])
			addNewNode(v1, g.adiacenti[v0])
		}
		i++
	}
	return g
}

func stampaGrafo(g grafo) {
	/*FORMATO:
	l’indice di un vertice su ogni riga, seguito dagli indici dei vertici a lui adiacentI
	fmt.Println("GRAFO: {")
	for ver, l := range g.adiacenti {
		fmt.Print(ver, ": ")
		if l != nil {
			printList(*l)
		}
		fmt.Println("")
	}
	fmt.Println("}")
}

func arcoDiretto(x, y int, g grafo) bool {
	puntatore := searchList(y, *g.adiacenti[x])

	return puntatore != nil
}*/
