# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">Hassan Donny Darmawan - 109082500030</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main
import "fmt"

func factorial(n int)int {
	hasil := 1
	for i:= 1; i<=n;{
		hasil*=i
		i++
	}
	return hasil
}


func permutation (n,r int)int {
	if n>=r {
		var hasil int
		hasil = factorial(n)/factorial(n-r)
		return hasil
	}
	return 0
}


func combination (n,r int)int {
	if n>=r {
		var hasil int
		hasil = factorial(n)/(factorial(r) * factorial(n-r))
		return hasil
	}
	return 0
}

func main () {
	var a,b,c,d int
	fmt.Scan(&a,&b,&c,&d)
	fmt.Printf("%d %d\n",permutation(a,c),combination(a,c))
	fmt.Printf("%d %d\n",permutation(b,d),combination(b,d))

}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul3/output/output-modul3-soal1.png)
[penjelasan]
<p align="justify">program ini bertujuan untuk melakukan pertukaran nilai variable. proses pertukaran ini menggunakan variable tambahan yaitu variable temp yang berfungsi sebagai tempat penyimpan nilai sementara agar saat nilai awal sedang melakukan pertukaran, nilai tersebut tidak hilang melainkan disimpan pada varible temp. lalu saat sudah selesai prosesnya maka user akan mendapatkan nilai variable yang memiliki urutan baru</p> 


### 2. [Soal]
#### soal2.go

```go
package main

import (
	"fmt"
	"math"
)

func fogoh (n int)int {
	h := n+1
	return int(float64(math.Pow(float64(h)-float64(2),2)))
}

func gohof (n int)int{
	f:= math.Pow(float64(n),float64(2))
	return int(f+1)-2
}

func hofog (n int)int{
	g := n-2
	return int((math.Pow(float64(g),float64(2)))+1)
}

func main (){
	var a,b,c int
	fmt.Scan(&a,&b,&c)
	fmt.Printf("%d\n", fogoh(a))
	fmt.Printf("%d\n", gohof(b))
	fmt.Printf("%d\n", hofog(c))
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul3/output/output-modul3-soal2.png)
[penjelasan]
<p align="justify">Program ini digunakan untuk membantu user apakah percobaan nya berhasil atau tidak, pada awal program ini berjalan user dimintakan input 4 warna hasil dari reaksi nantinya program ini akan mengecek apakah warna yang diinputkan user pada setiap percobaan nya sesuai urutan atau tidak, jika ada satu saja percobaan yang tidak urut warna nya maka program yang ada didalam if akan berjalan yaitu mengubah nilai status yang awalnya true menjadi false. diakhir status nya akan ditampilkan pada layar user.</p>



### 3. [Soal]
#### soal3.go

```go
package main 
import "fmt"

func main (){
var kg_parsel, sisa, parsel int
var biayautama, biayatambahan, total float64

fmt.Print("Berat parsel(gram): ")
fmt.Scanln(&parsel)
kg_parsel = parsel / 1000
sisa = parsel % 1000

fmt.Println("Detail berat:", kg_parsel, "kg +", sisa, "gr")
biayautama = float64(kg_parsel)*10000
if kg_parsel > 10{
	biayatambahan = 0
} else if sisa >=500 {
	biayatambahan = float64(sisa)*5
}else {
	biayatambahan = float64(sisa)*15
}

fmt.Println("Detail biaya: Rp.",biayautama, "+ Rp.",biayatambahan)
total = biayautama + biayatambahan
fmt.Println("Total biaya: Rp.",total)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul3/output/output-modul3-soal1.png)
[penjelasan]
<p align="justify">program ini untuk menghitung biaya pengiriman parsel, pada awal program user diminta menginputkan berat suatu parsel dalam gram lalu disimpan pada variable parsel, lalu program akan mencari berat utama (kg) dan berat yang lebih (sisa). </p>
<p>untuk berat utama dicari dengan nilai variable parsel dibagi 1000 sedangkan untuk berat berlebih dicari dengan variable parsel di modulus 1000. selanjutnya untuk mencari harga maka berat utama dikali dengan harga per kilonya yaitu 10000, sedangkan untuk biaya tambahan ini menggunakan if else karena di bagian ini terdapat 3 kondisi yang berbeda dengan action yang berbeda juga yaitu;</p>
<ol>
<li align="justify"> jika berat utama nya lebih dari 10 kg maka biaya tambahan(biaya berat sisa) digratiskan.</li>
<li align="justify"> jika berat utama kurang dari 10 kg dan berat sisa tidak kurang dari 500 maka harga 1 gr dari brt sisa adalah 5 rupiah.</li>
<li align="justify"> jika berat sisa kurang dari 500g maka  harga 1 gr dari berat sisa adalah 15 rupiah.</li>
</ol>
<p align="justify">selanjutnya untuk total hanya menambahkan saja biaya utama dengan biaya tambahan. diakhir total biaya akan ditampilkan pada user.</p>