package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("es calcolatrice: ", valuta("5 3 - 2 *"))
	fmt.Println("es conversione: ", converti("( ( 5 - 3 ) * 2 )"))
	fmt.Println("sam  <3  =^.^=")
	fmt.Print("es html: ")
	htmlBenFormato("<a> <b> </b> </c>")
	html()
}

// INPUT:una stringa
func htmlBenFormato(doc string) {
	var pila []string
	tag := strings.Split(doc, " ")

	for pos, t := range tag {
		if string(t[1]) != "/" { //tag di apertura
			pila = append(pila, string(t[1]))
		} else { //tag chiusura
			t2 := pila[len(pila)-1]
			if t2 != string(t[2]) {
				fmt.Println("err in posizione ", pos+1)
				return
			}
			pila = pila[:len(pila)-1]
		}
	}
	fmt.Print("sono rimasti aperti i seguenti tag: ")
	for _, t := range pila {
		fmt.Print("<", t, ">  ")
	}
	return
}

// INPUT:un tag per riga, il programma termina leggendo una riga vuota
func html() {
	var pila []string
	var pos int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		pos++
		if s.Text() == "" {
			break
		}
		//-----------------------------
		t := s.Text()
		if string(t[1]) != "/" { //tag di apertura
			pila = append(pila, string(t[1]))
		} else { //tag chiusura
			t2 := pila[len(pila)-1]
			if t2 != string(t[2]) {
				fmt.Println("err in posizione ", pos)
				return
			}
			pila = pila[:len(pila)-1]
		}
	}
	if len(pila) == 0 {
		return
	}
	fmt.Print("sono rimasti aperti i seguenti tag: ")
	for _, t := range pila {
		fmt.Print("<", t, ">  ")
	}
	return
}

func converti(e string) (s string) {
	var pila []string
	t := strings.Split(e, " ")

	for _, tok := range t {
		_, err := strconv.Atoi(tok)
		if err == nil { //è un numero
			s = s + tok + " "
		} else {
			switch tok {
			case ")":
				s = s + pila[len(pila)-1] + " "
				pila = pila[:len(pila)-1]
			case "(":

			default:
				pila = append(pila, tok)
			}

		}
	}
	return
}

func valuta(e string) int {
	var pila []int
	t := strings.Split(e, " ")

	for _, tok := range t {
		num, err := strconv.Atoi(tok)
		if err == nil { //è un numero
			pila = append(pila, num)
		} else {
			var newElem int
			switch tok {
			case "+":
				newElem = pila[len(pila)-1] + pila[len(pila)-1]
			case "-":
				newElem = pila[len(pila)-1] - pila[len(pila)-1]
			case "*":
				newElem = pila[len(pila)-1] * pila[len(pila)-1]
			case "/":
				newElem = pila[len(pila)-1] / pila[len(pila)-1]

			}
			pila = pila[:len(pila)-2]
			pila = append(pila, newElem)
		}
	}

	return pila[0]
}
