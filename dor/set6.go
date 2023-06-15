package main

import (
	"fmt"
)

type nodo struct {
	left  *nodo
	right *nodo
	val   int
}
type bitree struct {
	radice *nodo
}

func main() {
	/*var a bitree

	fmt.Println(" =^.^=")

	a.radice = newNode(1)
	a.radice.left = newNode(2)
	a.radice.right = newNode(3)
	a.radice.left.left = newNode(4)
	a.radice.left.right = newNode(5)
	a.radice.right.left = newNode(6)
	preorder(a.radice)
	fmt.Println(" =^.^=")
	preorder1(a.radice)
	fmt.Println(" =^.^=")
	inorder(a.radice)
	fmt.Println(" =^.^=")
	postorder(a.radice)
	fmt.Println(" =^.^=")
	stampaAlberoASommario(a.radice, 0)
	fmt.Println(" =^.^=")
	stampaAlbero(a.radice)
	fmt.Println(" =^.^=")
	stampaAlberoASommario(arr2tree([]int{69, 89, 28, 39, 66, 44, 12, 2, 71}, 0), 0)*/
	fmt.Println(" =^.^=   ORBITE: ", contaOrbite())
}
func newNode(val int) *nodo {
	return &nodo{nil, nil, val}
}

func preorder(a *nodo) {
	if a.right == nil && a.left == nil {
		fmt.Print(a.val, " ")
		return
	}
	fmt.Print(a.val, " ")
	if a.left != nil {
		preorder(a.left)
	}
	if a.right != nil {
		preorder(a.right)
	}
	return
}

func preorder1(a *nodo) {
	if a == nil {
		return
	}
	fmt.Print(a.val, " ")
	preorder(a.left)
	preorder(a.right)
	return
}

func inorder(a *nodo) {
	if a == nil {
		return
	}
	inorder(a.left)
	fmt.Print(a.val, " ")
	inorder(a.right)
}

func postorder(a *nodo) {
	if a == nil {
		return
	}
	inorder(a.left)
	inorder(a.right)
	fmt.Print(a.val, " ")
}

// 1.2 SOMMARIO
// preorder
func stampaAlberoASommario(node *nodo, spaces int) {

	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}
	fmt.Print("*")
	if node == nil {
		fmt.Println()
		return
	}
	fmt.Println(node.val)

	if node.left != nil || node.right != nil {
		stampaAlberoASommario(node.right, spaces+1)
		stampaAlberoASommario(node.left, spaces+1)
	}
}

// 1.3 STAMPA ALTERNATIVA
func stampaAlbero(node *nodo) {
	if node == nil {
		return
	}
	fmt.Print(node.val, "")
	if node.left == nil && node.right == nil {
		return
	}
	fmt.Print(" [")
	if node.left != nil {
		stampaAlbero(node.left)
	} else {
		fmt.Print("-")
	}
	fmt.Print(", ")
	if node.right != nil {
		stampaAlbero(node.right)
	} else {
		fmt.Print("-")
	}
	fmt.Print("]")
}

// 1.4 dal vettore all'albero
func arr2tree(a []int, i int) (root *nodo) {
	if i >= len(a) {
		return nil
	}
	root = newNode(a[i])
	root.left = arr2tree(a, 2*i+1)
	root.right = arr2tree(a, 2*i+2)
	return root

}

//ES 2

//ES 3
/* una slice
informazione | //in questo caso interi
	padre    | //vettore che fa riferimento alla posizione nel primo array del padre
*/

//3.1 advent of code 2019 day six
/*
1. Modellate la situazione con un albero: cosa rappresentano i nodi dell’albero?
cosa rappresenta la relazione padre/figlio?
NODI: corpi celesti
RELAZ PADRE FIGLIO: => figlio ruota attorno a padre

2. Riformulate i problemi usando la terminologia degli alberi:
a) PARTE 1: Cosa sono le orbite dirette?
		sono il padre diretto di ogni nodo
E le orbite indirette?
		antenati di un nodo, tolto il padre
Come può essere descritto, in termini di alberi, il numero di orbite indirette di un oggetto?
		livello dalla radice
b) PARTE 2 - Come può essere descritta, in termini di alberi, La distanza tra gli oggetti
attorno a cui orbitano YOU e SAN?
		che sonoooo?


3. Progettate una soluzione per calcolare:
a) Parte 1 - il numero di orbite indirette:
		rappresento m
b) Parte 2 - il numero di trasferimenti di orbita necessari
*/

func contaOrbite() (o int) {
	var corpiCel []string
	var padri []int
	var orbite []int

	input := [][]string{{"COM", "B"}, {"B", "C"}, {"C", "D"}, {"D", "E"}, {"E", "F"}, {"B", "G"}, {"G", "H"}, {"D", "I"}, {"E", "J"}, {"J", "K"}, {"K", "L"}}
	corpiCel = append(corpiCel, "COM")
	padri = append(padri, 0)
	orbite = append(orbite, 0)

	for _, corpi := range input {
		/*if trovaInSlice(corpi[0], corpiCel) == -1{
			corpiCel = append(corpiCel, corpi[0])
		}*/
		if trovaInSlice(corpi[1], corpiCel) == -1 {
			//fmt.Println("corpiCel: ", corpi[1], ", padri: ", trovaInSlice(corpi[0], corpiCel), ", orbite: ", orbite[trovaInSlice(corpi[0], corpiCel)]+1)
			corpiCel = append(corpiCel, corpi[1])
			padri = append(padri, trovaInSlice(corpi[0], corpiCel))
			orbite = append(orbite, orbite[trovaInSlice(corpi[0], corpiCel)]+1)
		}
	}
	for _, orb := range orbite {
		o = o + orb
	}
	return o
}

func trovaInSlice(st string, s []string) (pos int) { //-1 se non c'è
	for pos, c := range s {
		if c == st {
			return pos
		}
	}
	return -1
}
