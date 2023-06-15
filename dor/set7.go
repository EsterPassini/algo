package main

import (
	"fmt"
)

/*
STRUTTURA DATI: foresta di alberi (non binari)
a) stampa i figli
b) conta foglie
c) trovare padre dato figlio
d) stampa cammino da un nodo a radice in ordine
e) stampa ogni nodo indicandone il padre (tranne per le radici)

slice di capi per individuare i vari alberi
*/
type Algore struct {
	capi  []string
	dipe  map[string][]string //[supervisore] => dipendentI
	super map[string]string   //[dipendente] => supervisorE
}

func main() {
	fmt.Println("=^.^=")
	var d map[string][]string = make(map[string][]string)
	var s map[string]string = make(map[string]string)
	var c []string = []string{"a", "e"}
	var alg Algore

	dipA := []string{"b", "c", "d"}
	dipE := []string{"f", "g"}
	dipB := []string{"h"}

	d["a"] = dipA
	d["b"] = dipB
	d["e"] = dipE

	s["b"] = "a"
	s["c"] = "a"
	s["d"] = "a"
	s["f"] = "e"
	s["g"] = "e"
	s["h"] = "b"

	alg.capi = c
	alg.dipe = d
	alg.super = s

	stampaSubordinati("h", alg)
	fmt.Println("quantiSenzaSubordinati: ", quantiSenzaSubordinati(alg))
	fmt.Println(supervisore("h", alg))
	fmt.Println("=^.^=")
	stampaImpiegatiSopra("h", alg)
	fmt.Println("=^.^=")
	stampaImpiegatiConSupervisore(alg)
}

func stampaSubordinati(sup string, a Algore) {
	for _, sub := range a.dipe[sup] {
		fmt.Println(sub)
	}
}

func quantiSenzaSubordinati(a Algore) (d int) {
	for _, dip := range a.super {
		_, ok := a.super[dip]
		if !ok {
			d++
		}
	}
	return d
}

func supervisore(d string, a Algore) (sup string) {
	return a.super[d]
}

func stampaImpiegatiSopra(d string, a Algore) {
	for {
		fmt.Println(a.super[d])
		d = a.super[d]

		if InSlice(d, a.capi) {
			//fmt.Println(d) tolto perchè il capo lo stampava 2 volte :(
			return
		}
	}
}

func stampaImpiegatiConSupervisore(a Algore) {
	for chiave, dipendenti := range a.dipe {
		for _, dipendente := range dipendenti {
			fmt.Println(dipendente, " suo supervisore: ", chiave)
		}
	}
	for _, capo := range a.capi {
		fmt.Println(capo, " dipendente di max livello")
	}
	return
}

func InSlice(st string, s []string) bool {
	for _, c := range s {
		if c == st {
			return true
		}
	}
	return false
}
