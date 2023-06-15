package main

import (
	"fmt"
	"strconv"
)

type listNode struct {
	ele  int
	pros *listNode
}

type linkedList struct {
	head *listNode
}

func newNode1(ele int) *listNode {
	return &listNode{ele, nil}
}

func newNode2(ele int) *listNode {
	nodo := new(listNode)
	nodo.ele = ele
	return nodo
}

func addNewNode(ele int, l *linkedList) {
	n := newNode2(ele)
	n.pros = l.head
	l.head = n
}

func printList(l linkedList) {
	p := l.head

	for p != nil {
		fmt.Print(p.ele, " ")
		p = p.pros
	}
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

func RemoveItem(n int, l *linkedList) {
	var p, prec *listNode = l.head, nil

	for p != nil {
		if p.ele == n {
			if prec == nil {
				l.head = l.head.pros
			} else {
				prec.pros = p.pros
			}
			return
		}

		prec = p
		p = p.pros
	}

}

func sizeList(l linkedList) int {
	var p *listNode = l.head
	var s int

	if l.head == nil {
		return 0
	}

	s = 1

	for p.pros != nil {
		p = p.pros
		s++
	}
	return s
}

func reversePrintList(l linkedList) {
	s := sizeList(l)
	p := l.head

	for i := s - 1; i >= 0; i-- { //numero di volte che devo stampare
		for j := 0; j != i; j++ {
			p = p.pros
		}
		fmt.Print(p.ele, " ")
		p = l.head
	}
}

func removeAll(l *linkedList) {
	l.head = nil
	return
}

func main() {
	var l linkedList

	var ist string

	for {
		fmt.Scan(&ist)

		switch ist {

		case "+":
			fmt.Scan(&ist)
			ele, _ := strconv.Atoi(ist)
			fmt.Println("elementoo: ", ele)
			if searchList(ele, l) == nil {
				addNewNode(ele, &l)
				fmt.Printf("aggiunto %d alla lista\n", ele)
			} else {
				fmt.Printf("%d appartiene già alla lista\n", ele)
			}

		case "-":
			fmt.Scan(&ist)
			ele, _ := strconv.Atoi(ist)
			if searchList(ele, l) != nil {
				RemoveItem(ele, &l)
				fmt.Printf("rimosso %d dalla lista\n", ele)
			} else {
				fmt.Printf("%d non è contenuto nella lista\n", ele)
			}

		case "?":
			fmt.Scan(&ist)
			ele, _ := strconv.Atoi(ist)
			if searchList(ele, l) != nil {
				fmt.Printf("%d appartiene alla lista\n", ele)
			} else {
				fmt.Printf("%d non appartiene alla lista\n", ele)
			}

		case "f":
			fmt.Println("esecuzione fermata")
			return

		case "c":
			fmt.Println("taglia: ", sizeList(l))

		case "p":
			fmt.Print("lista: [ ")
			printList(l)
			fmt.Println("]")

		case "o":
			fmt.Print("lista al contrario: [ ")
			reversePrintList(l)
			fmt.Println("]")

		case "d":
			removeAll(&l)
			fmt.Println("rimossi tutti gli elementi dalla lista")
		}
	}

}

/*ESERCIZIO 2:
l'efficienza cambia in base all'operazione che va eseguita
operazioni più onerose che possono essere rese più efficienti con altre strutture:
-stampa al contrario O(n!)
-inserimento in coda (non era da implementare) O(n)

strutture alternative:
lista con puntatore alla coda:
	renderebbe l'inserimento in coda più efficiente O(1), al costo di
	memorizzare un puntatore in più
lista bidirezionale:
	al costo di memorizzare il doppio dei puntatori rende più efficiente la
	stampa al contrario O(n)
*/

/*ESERCIZIO 3.1
a) assegna ad un puntatore al nodo un intero (errore di tipo)
b) non riassegna il puntatore alla coda che rimane a puntare il penultimo nodo
c) non riassegna il puntatore di quello che ora è il penultimo nodo che rimane null
	la lista è dunque divisa in 2 pezzi
d) CORRETTO
e) l.head non centra nulla
*/

/*ESERCIZIO 3.2
quale lista è la migliore implementazione?
	a) LL: il puntatore alla coda è inutile in una serch di un elemento,
		per cui si risparmia la memoria di questo puntatore
	b) LLWT: l'operazione ha complessità O(1), a differenza che con una LL che ha O(n),
		visto che bisogna eseguire una scansione della lista fino all'ultimo elemento
	c) LL: il puntatore alla coda è inutile in una serch di un elemento,
		per cui si risparmia la memoria di questo puntatore
	d) LLWT: l'operazione ha complessità O(1), a differenza che con una LL che ha O(n),
		visto che bisogna eseguire una scansione della lista fino all'ultimo elemento

*/

/* ESERCIZIO 4
b)
(...)
*/
