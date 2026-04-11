# <h1 align="center">Laporan Praktikum Modul 4 - ... </h1>
<p align="center">Hassan Donny Darmawan - 109082500030</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main 
import "fmt"

func main () {
	var n int
	fmt.Scan(&n)

	for i:=0; i<=n; i++ {
		fmt.Printf("%d ",fibonaci(i))
	}
}

func fibonaci(n int)int{
	if n <= 1 {
		return n
	}
	return fibonaci(n-1) + fibonaci(n-2)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul5/Output/soal1_modul5.png)
[penjelasan]
<p align="justify">Program ini bertujuan untuk menampilkan deret bilangan Fibonacci dari indeks ke-0 hingga ke-n. Bilangan Fibonacci adalah suatu barisan bilangan dimana setiap bilangan merupakan hasil penjumlahan dua bilangan sebelumnya, dimulai dari 0 dan 1.</p>

<p align="justify">Fungsi <code>fibonaci(n int) int</code> merupakan fungsi rekursif yang menghitung nilai Fibonacci pada indeks ke-n. Fungsi ini memiliki base case yaitu apabila nilai n kurang dari atau sama dengan 1, maka fungsi akan langsung mengembalikan nilai n itu sendiri (0 jika n=0, dan 1 jika n=1). Apabila nilai n lebih dari 1, fungsi akan memanggil dirinya sendiri secara rekursif dengan mengembalikan hasil penjumlahan <code>fibonaci(n-1) + fibonaci(n-2)</code> hingga mencapai base case tersebut.</p>

<p align="justify">Pada fungsi utama atau <code>func main</code>, user diminta untuk menginputkan sebuah bilangan bulat n. Nilai n tersebut kemudian digunakan sebagai batas perulangan <code>for</code> yang berjalan dari indeks 0 hingga n (inklusif). Pada setiap iterasi perulangan, program akan memanggil fungsi <code>fibonaci(i)</code> dan langsung mencetak hasilnya secara berurutan, sehingga menghasilkan tampilan deret bilangan Fibonacci lengkap dari F(0) sampai F(n).</p>

### 2. [Soal]
#### soal2.go

```go
package main
import "fmt"

func bintang (n int){
	if n!=0 {
		fmt.Print("*")
		bintang(n-1)
	}
}

func main (){
	var n int
	fmt.Scan(&n)

	for i:= 1;i<=n;i++{
		bintang(i)
		fmt.Println()
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul5/Output/soal2_modul5.png)
[penjelasan]
<p align="justify">Program ini bertujuan untuk menampilkan pola segitiga bintang yang tersusun dari karakter <code>*</code>, dimana jumlah baris dan banyaknya bintang pada setiap baris ditentukan oleh input dari user.</p>

<p align="justify">Fungsi <code>bintang(n int)</code> merupakan fungsi rekursif yang bertugas mencetak karakter <code>*</code> sebanyak n kali. Fungsi ini memiliki base case yaitu apabila nilai n sama dengan 0, maka fungsi akan berhenti dan tidak mencetak apapun. Apabila nilai n tidak sama dengan 0, fungsi akan mencetak satu karakter <code>*</code> kemudian memanggil dirinya sendiri secara rekursif dengan nilai <code>n-1</code>, proses ini terus berulang hingga nilai n mencapai 0.</p>

<p align="justify">Pada fungsi utama atau <code>func main</code>, user diminta untuk menginputkan sebuah bilangan bulat n yang menentukan jumlah baris segitiga yang akan ditampilkan. Nilai n tersebut digunakan sebagai batas perulangan <code>for</code> yang berjalan dari i=1 hingga i=n. Pada setiap iterasi, program memanggil fungsi <code>bintang(i)</code> untuk mencetak bintang sebanyak i buah, kemudian <code>fmt.Println()</code> dipanggil untuk pindah ke baris berikutnya. Hasilnya adalah pola segitiga bintang yang bertambah satu bintang di setiap barisnya, dari baris pertama dengan 1 bintang hingga baris terakhir dengan n bintang.</p>

### 3. [Soal]
#### soal3.go

```go
package main
import "fmt"

func pembagian (n int, i int){
	if n%i == 0{
		fmt.Print(i," ")
	}
	if n>=i{
		pembagian(n,i+1)
	}
}

func main (){
	var n int
	fmt.Scan(&n)
	pembagian(n,1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul5/Output/soal3_modul5.png)
<p align="justify">Program ini bertujuan untuk menampilkan semua faktor pembagi dari suatu bilangan bulat n yang diinputkan oleh user. Faktor pembagi adalah bilangan yang dapat membagi habis bilangan n tanpa menyisakan sisa.</p>

<p align="justify">Fungsi <code>pembagian(n int, i int)</code> merupakan fungsi rekursif yang bertugas mencari dan mencetak semua faktor pembagi dari bilangan n. Fungsi ini bekerja dengan dua kondisi pengecekan. Kondisi pertama yaitu apabila <code>n%i == 0</code> bernilai true, artinya bilangan i dapat membagi habis n, maka nilai i akan dicetak sebagai salah satu faktor pembagi. Kondisi kedua yaitu apabila <code>n >= i</code> bernilai true, maka fungsi akan memanggil dirinya sendiri secara rekursif dengan menaikkan nilai i menjadi <code>i+1</code> untuk memeriksa bilangan berikutnya. Proses rekursif ini akan terus berjalan hingga nilai i melebihi nilai n, yang menjadi base case berhentinya fungsi.</p>

<p align="justify">Pada fungsi utama atau <code>func main</code>, user diminta untuk menginputkan sebuah bilangan bulat n. Nilai n tersebut kemudian dikirimkan ke fungsi <code>pembagian(n, 1)</code> dengan nilai awal i=1, karena pencarian faktor pembagi dimulai dari angka 1. Hasilnya adalah seluruh faktor pembagi dari bilangan n yang ditampilkan secara berurutan dari yang terkecil hingga terbesar.</p>

### 4. [Soal]
#### soal4.go

```go
package main
import "fmt"

func baris (n int){
	if n == 1{
		fmt.Print(1, " ")
		return
	}
	fmt.Print(n," ")
	baris(n-1)
	fmt.Print(n," ")
}

func main (){
	var n int
	fmt.Scan(&n)
	baris(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul5/Output/soal4_modul5.png)
[penjelasan]
<p align="justify">Program ini bertujuan untuk menampilkan pola bilangan yang tersusun secara simetris, dimana bilangan dari n turun hingga 1 kemudian naik kembali hingga n, kecuali angka 1 yang hanya muncul satu kali di tengah sebagai titik balik.</p>

<p align="justify">Fungsi <code>baris(n int)</code> merupakan fungsi rekursif yang bertugas mencetak pola bilangan simetris tersebut. Fungsi ini memiliki base case yaitu apabila nilai n sama dengan 1, maka fungsi akan mencetak angka 1 diikuti spasi kemudian langsung berhenti dengan perintah <code>return</code>. Apabila nilai n lebih dari 1, fungsi akan mencetak nilai n terlebih dahulu, kemudian memanggil dirinya sendiri secara rekursif dengan nilai <code>n-1</code>, dan setelah pemanggilan rekursif selesai barulah fungsi mencetak nilai n sekali lagi. Mekanisme inilah yang menciptakan efek simetris, karena pencetakan nilai n dilakukan dua kali yaitu sebelum dan sesudah pemanggilan rekursif, sedangkan nilai 1 hanya dicetak satu kali di tengah.</p>

<p align="justify">Pada fungsi utama atau <code>func main</code>, user diminta untuk menginputkan sebuah bilangan bulat n. Nilai n tersebut kemudian dikirimkan ke fungsi <code>baris(n)</code> untuk diproses. Sebagai contoh apabila user menginputkan nilai 4, maka output yang dihasilkan adalah <code>4 3 2 1 2 3 4</code>, dimana bilangan turun dari 4 hingga 1 lalu naik kembali dari 2 hingga 4 secara simetris.</p>

### 5. [Soal]
#### soal5.go

```go
package main 
import "fmt"

func ganjil (n int, a int){
	if a<=n{
		fmt.Print(a," ")
		ganjil(n, a+2)
	}
}

func main (){
	var n int
	fmt.Scan(&n)
	ganjil(n,1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul5/Output/soal5_modul5.png)
[penjelasan]
<p align="justify">Program ini bertujuan untuk menampilkan semua bilangan ganjil yang dimulai dari 1 hingga tidak melebihi nilai n yang diinputkan oleh user. Bilangan ganjil adalah bilangan yang tidak habis dibagi 2, seperti 1, 3, 5, 7, dan seterusnya.</p>

<p align="justify">Fungsi <code>ganjil(n int, a int)</code> merupakan fungsi rekursif yang bertugas mencetak bilangan ganjil secara berurutan. Fungsi ini memiliki base case yaitu apabila nilai a melebihi nilai n, maka fungsi akan berhenti dan tidak mencetak apapun. Apabila nilai <code>a <= n</code> bernilai true, fungsi akan mencetak nilai a yang merupakan bilangan ganjil saat itu, kemudian memanggil dirinya sendiri secara rekursif dengan menambahkan nilai a sebesar 2 menjadi <code>a+2</code>. Penambahan sebesar 2 inilah yang memastikan setiap pemanggilan rekursif selalu menghasilkan bilangan ganjil berikutnya, karena bilangan ganjil selalu memiliki selisih 2 antara satu dengan yang lainnya.</p>

<p align="justify">Pada fungsi utama atau <code>func main</code>, user diminta untuk menginputkan sebuah bilangan bulat n sebagai batas atas pencetakan bilangan ganjil. Nilai n tersebut kemudian dikirimkan ke fungsi <code>ganjil(n, 1)</code> dengan nilai awal a=1, karena bilangan ganjil positif pertama dimulai dari angka 1. Sebagai contoh apabila user menginputkan nilai 10, maka output yang dihasilkan adalah <code>1 3 5 7 9</code>, yaitu seluruh bilangan ganjil dari 1 hingga nilai yang tidak melebihi 10.</p>

### 6. [Soal]
#### soal6.go

```go
package main 
import "fmt"

func pangkat (n, a  int )int{
	if a == 0{
		return 1
	}
	return n * pangkat(n, a-1)
}

func main (){
	var n,a int
	fmt.Scan(&n,&a)
	fmt.Print(pangkat(n,a))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul5/Output/soal6_modul5.png)
[penjelasan]
<p align="justify">Program ini bertujuan untuk menghitung hasil perpangkatan dari suatu bilangan n yang dipangkatkan dengan a, dimana kedua nilai tersebut diinputkan oleh user. Perpangkatan sendiri merupakan operasi matematika yang mengalikan bilangan n dengan dirinya sendiri sebanyak a kali.</p>

<p align="justify">Fungsi <code>pangkat(n, a int) int</code> merupakan fungsi rekursif yang bertugas menghitung hasil perpangkatan dari n pangkat a. Fungsi ini memiliki base case yaitu apabila nilai a sama dengan 0, maka fungsi akan mengembalikan nilai 1, hal ini sesuai dengan sifat matematika bahwa bilangan apapun yang dipangkatkan 0 hasilnya selalu 1. Apabila nilai a tidak sama dengan 0, fungsi akan mengembalikan hasil perkalian n dengan pemanggilan rekursif <code>pangkat(n, a-1)</code>, dimana nilai a dikurangi 1 pada setiap pemanggilan rekursif hingga mencapai base case. Mekanisme inilah yang secara bertahap mengalikan n dengan dirinya sendiri sebanyak a kali untuk menghasilkan nilai perpangkatan yang benar.</p>

<p align="justify">Pada fungsi utama atau <code>func main</code>, user diminta untuk menginputkan dua buah bilangan bulat yaitu n sebagai bilangan yang akan dipangkatkan dan a sebagai nilai pangkatnya. Kedua nilai tersebut kemudian dikirimkan ke fungsi <code>pangkat(n, a)</code> dan hasilnya langsung dicetak menggunakan <code>fmt.Print</code>. Sebagai contoh apabila user menginputkan nilai n=2 dan a=5, maka output yang dihasilkan adalah <code>32</code>, karena 2 dipangkatkan 5 menghasilkan nilai 32.</p>