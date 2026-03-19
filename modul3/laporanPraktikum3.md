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
<p align="justify">Program ini digunakan untuk membantu user apakah percobaan nya berhasil atau tidak, pada awal program ini berjalan user dimintakan input 4 warna hasil dari reaksi nantinya program ini akan mengecek apakah warna yang diinputkan user pada setiap percobaan nya sesuai urutan atau tidak, jika ada satu saja percobaan yang tidak urut warna nya maka program yang ada didalam if akan berjalan yaitu mengubah nilai status yang awalnya true menjadi false. diakhir status nya akan ditampilkan pada layar user.</p>



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
<p align="justify">program ini untuk menghitung biaya pengiriman parsel, pada awal program user diminta menginputkan berat suatu parsel dalam gram lalu disimpan pada variable parsel, lalu program akan mencari berat utama (kg) dan berat yang lebih (sisa). </p>