/*
Un pesce lanterna crea un nuovo pesce lanterna ogni 7 giorni. Ma il processo non è necessariamente sincronizzato,
dipende dall’età dei pesci
(un pesce lanterna di 5 giorni farà nascere un nuovo pesce fra 2 giorni;
un pesce che ha 3 giorni lo farà nascere fra 4 giorni)

Hai la lista delle età di centinaia di pesci lanterna e devi stabilire quanti pesci ci saranno dopo 80 giorni
*/

package main

import "fmt"

func main() {
	var p []int = []int{3, 4, 3, 1, 2}
	var g int = 80

	for i := 0; i < g; i++ { //giorni che passano
		for tim := len(p) - 1; tim >= 0; tim-- { //scandisce slide pesci al contrario
			if p[tim] == 0 {
				p[tim] = 6
				p = append(p, 8)
			} else {
				p[tim]--
			}
		}
	}

	fmt.Println("dopo 80 giorni ci sono: ", len(p), " pesci")

}
