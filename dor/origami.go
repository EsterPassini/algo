/*
Si raggiunge un'altra parte vulcanicamente attiva della grotta.
Sarebbe bello se potessi fare una sorta di imaging termico in modo da poter dire in anticipo quali grotte sono troppo calde per entrare in sicurezza.

Fortunatamente, il sottomarino sembra essere dotato di una termocamera!
Quando lo attivi, sei accolto con:
Congratulazioni per l'acquisto! Per attivare questa termografia a infrarossi
sistema di telecamere, inserire il codice che si trova a pagina 1 del manuale.
Apparentemente, gli Elfi non hanno mai usato questa funzione. Con tua sorpresa, riesci a trovare il manuale;
mentre vai ad aprirlo, la pagina 1 cade. È un grande foglio di carta trasparente!
La carta trasparente è contrassegnata da punti casuali e include istruzioni su come piegarla (input del puzzle). Per esempio

(...)<

La prima sezione è un elenco di punti sulla carta trasparente.
0,0 rappresenta la coordinata in alto a sinistra. Il primo valore, x, aumenta verso destra.
Il secondo valore, y, aumenta verso il basso.
Quindi, la coordinata 3,0 è a destra di 0,0 e la coordinata 0,7 è sotto 0,0.
Le coordinate in questo esempio formano il modello seguente, dove # è un punto sulla carta e . è una posizione vuota e non contrassegnata:

(...)


Quindi, c'è un elenco di istruzioni per la piegatura.
Ogni istruzione indica una linea sulla carta trasparente e richiede di piegare la carta verso l'alto
(per y=... linee orizzontali) o verso sinistra (per x=... linee verticali).
In questo esempio, la prima istruzione di piegatura è piega lungo y=7, che designa la linea formata da tutte le posizioni in cui y è 7 (contrassegnata qui con -):



La carta trasparente è piuttosto grande, quindi per ora concentrati solo sul completamento della prima piega.
Dopo la prima piega nell'esempio sopra, sono visibili 17 punti - i punti che finiscono per sovrapporsi dopo che la piega è stata completata contano come un singolo punto.

Quanti punti sono visibili dopo aver completato solo la prima istruzione di piegatura sulla carta trasparente?
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var f [][]int //matrice bidimensionale
	var input [][2]int
	var yMax, xMax int

	input, xMax, yMax = LeggiInput()

	f = CreaFoglio(input, yMax, xMax)
	fmt.Println("Foglio:")
	for _, s := range f {
		fmt.Println(s)
	}

	fmt.Println("\n Foglio piegato lungo orizzontale:")
	fy := piegay(f, 7)
	for _, s := range fy {
		fmt.Println(s)
	}

	fx := piegax(fy, 5)
	fmt.Println("\n Foglio piegato lungo verticale:")
	for _, s := range fx {
		fmt.Println(s)
	}
}

func LeggiInput() (input [][2]int, xMax int, yMax int) {
	var x, y int

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		//-----------------------------
		p := s.Text()
		punti := strings.Split(p, ",")
		x, _ = strconv.Atoi(punti[0])
		y, _ = strconv.Atoi(punti[1])

		if x > xMax {
			xMax = x
		}
		if y > yMax {
			yMax = y
		}

		input = append(input, [2]int{x, y})
	}

	return
}

func CreaFoglio(punti [][2]int, xMax int, yMax int) (f [][]int) {
	xMax = xMax + 1
	yMax = yMax + 1
	f = make([][]int, xMax)

	for i := 0; i < xMax; i++ {
		f[i] = make([]int, yMax)
		for t := 0; t < yMax; t++ {
			f[i][t] = 0
		}
	}
	//fai buchi
	for _, punto := range punti {
		x := punto[0]
		y := punto[1]

		f[y][x] = 1
	}

	return f
}

func piegay(f [][]int, y int) (fy [][]int) {

	f1 := f[:len(f)/2]
	f2 := f[len(f)/2+1:]

	fy = f1

	r2 := len(f2) - 1
	for r1 := 0; r1 < len(f1); r1++ {
		for c := 0; c < len(f[0]); c++ {
			if f2[r2][c] == 1 {
				fy[r1][c] = 1
			}
		}
		r2--
	}
	return
}

func piegax(f [][]int, x int) (fx [][]int) {

	for riga := 0; riga < len(f); riga++ {
		for col := len(f[riga])/2 + 1; col < len(f[riga]); col++ {

			if f[riga][col] == 1 {
				f[riga][len(f[riga])/2-(col-len(f[riga])/2)] = 1
			}
		}
	}

	fx = make([][]int, len(f))
	for r := 0; r < len(f); r++ {
		fx[r] = append(fx[r], f[r][:len(f[0])/2]...)
	}
	return fx
}

/*
INPUT COORDINATE BUCHI:

6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

*/
