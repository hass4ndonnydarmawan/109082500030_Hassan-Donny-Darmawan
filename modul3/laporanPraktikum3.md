# <h1 align="center">Laporan Praktikum Modul 3 - ... </h1>
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
<p align="justify">program ini bertujuan untuk mengitung permutasi dan combination, dalam menghitung ini program menggunakan rumus</p>
<ol>
<li>P(n,r) = n! / (n−r)!</li>
<li>C(n,r) = n! / (r!(n−r)!)</li>
</ol> 
<p align = "justify">function factorial disini digunakan untuk n! sesuai yang ada dirumus, nilai ini kemudian digunakan func combination dan permutation sesuai rumus masing masing, dengan syarat n >= r jika tidak memenuhi maka hasil nya 0.</p>
<p align = "justify">pada function utama atau func main user diminta menginputkan nilai untuk dihitung combination dan permutation nya dan akan ditampilkan kembali hasil perhitunganya setelah semua program selesai</p>

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
<p align="justify">Program ini bertujuan untuk menghitung hasil dari komposisi beberapa fungsi matematika berdasarkan input yang diberikan oleh user. terdapat 3 fungsi komposisi yaitu fogoh,gohof dan hofog.</p>
<ol>
<li>Fungsi fogoh(n) merepresentasikan f(g(h(n))), dari beberapa tahap hasil akhir yang didapatkan (n − 1)².</li>
<li>Fungsi gohof(n) merepresentasikan g(h(f(n))). dari beberapa tahap hasil akhir yang akan didapatkan n² − 1.</li>
<li>Fungsi hofog(n) merepresentasikan h(f(g(n))), hasil akhirnya adalah (n − 2)² + 1.</li>
</ol>
<p align ="justify">pada func main program menerima 3 inputan dalam bilangan bulat, masing masing inputan akan digunakan sebagai parameter ketiga fungsi, dan hasil akhir yang didapat dari perhitungan akan ditampilkan kepada layar user.</p>



### 3. [Soal]
#### soal3.go

```go
package main 
import (
	"fmt"
	"math"
)

func jarak (a,b,c,d float64)float64 {
	return math.Sqrt(math.Pow((a-c),2) + math.Pow((b-d),2))
}

func didalam (cx,cy,r,x,y float64) bool {
	return jarak(cx,cy,x,y) <= r
}

func main (){
	var (
		cx1,cy1,cx2,cy2,r1,r2 float64
		x,y float64
	)
	fmt.Scan(&cx1,&cy1,&r1)
	fmt.Scan(&cx2,&cy2,&r2)
	fmt.Scan(&x,&y)

	d1 := didalam(cx1,cy1,r1,x,y)
	d2 := didalam(cx2,cy2,r2,x,y)

	if d1 && d2 {
		fmt.Print("titik didalam lingkaran 1 dan 2")
	}else if d1 {
		fmt.Print("titik didalam lingkaran 1")
	}else if d2 {
		fmt.Print("titik didalam lingkaran 2")
	}else {
		fmt.Print("titik diluar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul3/output/output-modul3-soal3.png)
[penjelasan]
<p align="justify">Program ini menentukan posisi sebuah titik terhadap dua lingkaran.func jarak untuk  menghitung jarak antara kedua titik sedangkan func didalam mengecek apakah titik berada di dalam lingkaran dengan membandingkan jarak ke pusat dengan radius.</p>
<p align="justify">pada func main, program membaca dua lingkaran dan satu titik, lalu mengecek apakah titik berada di dalam masing-masing lingkaran. Hasilnya digunakan untuk menentukan apakah titik berada di dalam kedua lingkaran atau salah satu atau di luar keduanya yang kemudian akan ditampilkan.</p>
