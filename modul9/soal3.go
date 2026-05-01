package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var hasilPertandingan []string 

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	pertandinganKe := 1

	for {
		var skorA, skorB int

		fmt.Printf("Pertandingan %d : ", pertandinganKe)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasilPertandingan = append(hasilPertandingan, klubA)
		} else if skorB > skorA {
			hasilPertandingan = append(hasilPertandingan, klubB)
		} else {
			hasilPertandingan = append(hasilPertandingan, "Draw")
		}

		pertandinganKe++
	}

	for i := 0; i < len(hasilPertandingan); i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, hasilPertandingan[i])
	}
	
	fmt.Println("Pertandingan selesai")
}