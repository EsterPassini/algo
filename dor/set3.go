package main

import (
	"fmt"
)

/*func insertionSort() (s []int) {
	var n, i int
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}

		s = append(s, n)
		for i = len(s) - 2; i >= 0 && s[i] > n; i-- {
			s[i+1] = s[i]
		}

		s[i+1] = n

	}
	return

}*/

func insertionSort() (num []int) {
	var n int
	fmt.Scan(&n)
	for {
		fmt.Scan(&n)
		fmt.Println("voglio inserire:", n)

		if n == 0 {
			break
		}

		if len(num) == 0 {
			num = append(num, n)
		} else {
			//len !=0
			for t := 0; t < len(num); t++ {
				fmt.Println(len(num))
				fmt.Println(t, "aaaa")

				//è da inserire in fondo
				if t == len(num) {
					fmt.Println("inserisco:", n)
					num = append(num, n)
					break
				} else {
					//devo spostare gli elem
					num = append(num, num[len(num)-1])
					for j := len(num) - 2; j >= t; j-- {
						num[j] = num[j+1]
					}

					//inserisco elem
					num[t] = n
				}
			}
		}
	}
	return num
}

func selectionSort(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		//cerca minimo parte nn ordinata
		min := s[i]
		m := i
		for j := i + 1; j < len(s); j++ {
			if s[j] < min {
				min = s[j]
				m = j
			}
		}

		s[i], s[m] = s[m], s[i]
	}
	return s
}

func selectionSortNNSOLEGGEREEEE(s []int) []int {
	for i := len(s) - 1; i > 0; i-- {
		//cerca max parte nn ordinata
		max := s[i]
		M := i
		for j := 0; j < i; j++ {
			if s[j] > max {
				max = s[j]
				M = j
			}
		}

		s[i], s[M] = s[M], s[i]
	}
	return s
}

func selectionSortRec(s []int) {
	if len(s) == 0 || len(s) == 1 {
		return
	}

	//trova max e lo mette in ultima posizione
	max := s[len(s)-1]
	M := len(s) - 1
	for j := 0; j < len(s)-1; j++ {
		if s[j] > max {
			max = s[j]
			M = j
		}
	}

	s[len(s)-1], s[M] = s[M], s[len(s)-1]

	selectionSortRec(s[:len(s)-1])
	return

}

/*func selectionSortRec2(s []int) []int{
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	//trova max e lo mette in ultima posizione
	max := s[len(s)-1]
	M := len(s) - 1
	for j := 0; j < len(s)-1; j++ {
		if s[j] > max {
			max = s[j]
			M = j
		}
	}

	s[len(s)-1], s[M] = s[M], s[len(s)-1]

	return append(s, selectionSortRec(s[:len(s)-1]) )

}*/

func mergeSort(s []int) []int {
	//BASE
	if len(s) <= 1 {
		return s
	}

	//PASSO
	m := len(s) / 2

	s1 := mergeSort(s[:m])
	s2 := mergeSort(s[m:])

	s = merge(s1, s2)

	return s
}

func merge(s1, s2 []int) (s []int) {
	var i1, i2 int

	for i1 < len(s1) && i2 < len(s2) {
		if s1[i1] <= s2[i2] {
			s = append(s, s1[i1])
			i1++
		} else {
			s = append(s, s2[i2])
			i2++
		}
	}

	if i1 < len(s1) {
		for i1 < len(s1) {
			s = append(s, s1[i1])
			i1++
		}
	}

	if i2 < len(s2) {
		for i2 < len(s2) {
			s = append(s, s2[i2])
			i2++
		}
	}

	return s
}

func main() {
	var s []int = []int{1, 7, 2, 3, 4, 0}
	//fmt.Println(insertionSort())
	//fmt.Println(selectionSortNNSOLEGGEREEEE(s))
	//selectionSortRec(s)
	//fmt.Println(s)

	fmt.Println(mergeSort(s))

}
