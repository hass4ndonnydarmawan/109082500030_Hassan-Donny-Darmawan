package main

import (
	"fmt"
	"math"
	"strings"
)

const KapasitasMaksimal = 100

func main() {
	var arr [KapasitasMaksimal]int
	var n int

	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&n)

	if n <= 0 || n > KapasitasMaksimal {
		fmt.Println("Jumlah elemen tidak valid atau melebihi kapasitas!")
		return
	}

	fmt.Println("Masukkan", n, "bilangan bulat (pisahkan dengan spasi/enter):")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(strings.Repeat("-", 40))

	fmt.Print("a. Keseluruhan isi array: ")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("b. Elemen indeks ganjil  : ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("c. Elemen indeks genap   : ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("d. Masukkan nilai x (untuk kelipatan indeks): ")
	fmt.Scan(&x)
	fmt.Printf("   Elemen indeks kelipatan %d: ", x)
	for i := 0; i < n; i++ {
		if x != 0 && i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var hapusIdx int
	fmt.Print("e. Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&hapusIdx)
	
	for i := hapusIdx; i < n-1; i++ {
		arr[i] = arr[i+1]
	}

	fmt.Print("   Isi array setelah dihapus: ")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var total float64
	for i := 0; i < n; i++ {
		total += float64(arr[i])
	}
	rataRata := total / float64(n)
	fmt.Printf("f. Rata-rata bilangan     : %.2f\n", rataRata)

	var totalVarian float64
	for i := 0; i < n; i++ {
		totalVarian += math.Pow(float64(arr[i])-rataRata, 2) 
	}
	standarDeviasi := math.Sqrt(totalVarian / float64(n))
	fmt.Printf("g. Standar Deviasi        : %.2f\n", standarDeviasi)

	var cari, frekuensi int
	fmt.Print("h. Masukkan bilangan untuk dicari frekuensinya: ")
	fmt.Scan(&cari)
	for i := 0; i < n; i++ {
		if arr[i] == cari {
			frekuensi++
		}
	}
	fmt.Printf("   Bilangan %d muncul sebanyak %d kali.\n", cari, frekuensi)
}