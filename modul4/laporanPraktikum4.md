# <h1 align="center">Laporan Praktikum Modul 4 - ... </h1>
<p align="center">Hassan Donny Darmawan - 109082500030</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main
import "fmt"

func faktorial (n int, hasil *int){
	*hasil = 1
	for i:= 1; i<=n;  i++{
		*hasil *= i
	}
}
func permutation (n, r int, hasil *int){
	var fn, fnr int
	if n>=r {
		faktorial(n,&fn)
		faktorial(n-r,&fnr)

		*hasil = fn/fnr
	}else {
		*hasil = 0
	}
}

func combination (n,r int, hasil *int){
	var fr, fnr, fn int
	if n>=r{
	faktorial(n,&fn)
	faktorial(r,&fr)
	faktorial(n-r,&fnr)

	*hasil = fn/(fr*fnr)
	}else {
		*hasil = 0
	}
}

func main (){
	var a,b,c,d,p,co int
	fmt.Scan(&a,&b,&c,&d)
	permutation(a,c,&p)
	combination (a,c,&co)
	fmt.Println(p,co)

	permutation(b,d,&p)
	combination (b,d,&co)
	fmt.Println(p,co)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul4/output/output_soal1_modul4.png)
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
import "fmt"

func hitungSkor (soal, waktu_m *int){
	*waktu_m = 0
	*soal = 0
	var waktu int
	for i:= 1; i<= 8; i++{
		fmt.Scan(&waktu)

		if waktu<301 {
			*soal ++
			*waktu_m += waktu
		}
	}
}

func main (){
	var nama, pemenang string
	var soal, waktu int
	cek := true
	for {
		fmt.Scan(&nama)
		if nama == "Selesai"{
			break
		}
		var banyaksoal, waktu_terdikit int
		hitungSkor(&banyaksoal,&waktu_terdikit)

		if cek{
			soal = banyaksoal
			waktu = waktu_terdikit
			pemenang = nama
			cek = false
		}else if banyaksoal>soal || (soal== banyaksoal && waktu_terdikit<waktu) {
			soal = banyaksoal
			waktu = waktu_terdikit
			pemenang = nama
		}
	}

	fmt.Println(pemenang, soal, waktu)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul4/output/output_soal2_modul4.png)
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
import "fmt"

func deret(n int){
	fmt.Printf("%d ",n)
	for n!=1 && n>0{
		if n%2==0 {
			n = n/2
		} else {
			n = 3*n + 1
		}
		fmt.Printf("%d ",n)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	deret(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul4/output/output_soal3_modul4.png)
[penjelasan]
<p align="justify">Program ini menentukan posisi sebuah titik terhadap dua lingkaran.func jarak untuk  menghitung jarak antara kedua titik sedangkan func didalam mengecek apakah titik berada di dalam lingkaran dengan membandingkan jarak ke pusat dengan radius.</p>
<p align="justify">pada func main, program membaca dua lingkaran dan satu titik, lalu mengecek apakah titik berada di dalam masing-masing lingkaran. Hasilnya digunakan untuk menentukan apakah titik berada di dalam kedua lingkaran atau salah satu atau di luar keduanya yang kemudian akan ditampilkan.</p>