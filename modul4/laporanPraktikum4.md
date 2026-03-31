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
<p align = "justify">procedur factorial disini digunakan untuk n! sesuai yang ada dirumus, nilai ini kemudian digunakan procedure combination dan permutation dengan mengirimkan alamat memori untuk menyimpan hasil dari faktorial. nilai faktorial tersebut akan dikirimkan ke alamat memori yang dikasih apabila n>=r bernilai true dan hasil dari permutation atau combination juga akan dikirim ke alamat memori yang dikasih saat procedure combination atau permutation ini dipanggil, jika false maka nilai yang akan dikirim ke pemanggil procedure permutation dan combination adalah 0.</p>
<p align = "justify">pada function utama atau func main user diminta menginputkan nilai untuk dihitung combination dan permutation lalu nilai tersebut dimasukan pada procedur combination dan permutation sebagai parameter dan mengirimkan alamat memori untuk menyimpan hasil dari perhitungan permutation dan combination. nantinya hasil tersebut akan ditampilkan pada user.</p>

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
<p align="justify">Program ini bertujuan untuk menentukan pemenang dari sebuah kompetisi soal berdasarkan jumlah soal yang berhasil dijawab dan waktu yang digunakan. Terdapat satu fungsi eksternal yang dipanggil yaitu <code>hitungSkor</code>.</p>

<ol>
<li>Fungsi <code>hitungSkor(&banyaksoal, &waktu_terdikit)</code> dipanggil untuk setiap peserta dan mengisi dua nilai yaitu <code>banyaksoal</code> (jumlah soal yang berhasil dijawab) dan <code>waktu_terdikit</code> (total waktu yang digunakan), keduanya dikirim sebagai pointer sehingga nilai aslinya dapat diubah langsung oleh fungsi tersebut.</li>
</ol>

<p align="justify">Pada <code>func main</code>, program terus menerima input nama peserta secara berulang menggunakan perulangan <code>for</code> tanpa kondisi (infinite loop). Jika nama yang dimasukkan adalah <code>"Selesai"</code>, maka perulangan dihentikan. Untuk setiap peserta yang valid, program memanggil <code>hitungSkor</code> untuk mendapatkan skor mereka. Peserta pertama langsung ditetapkan sebagai pemenang sementara menggunakan variabel <code>cek</code>. Untuk peserta berikutnya, program membandingkan skornya dengan pemenang sementara — peserta baru akan menggantikan pemenang sementara jika ia menjawab soal lebih banyak, atau jika jumlah soalnya sama namun waktu yang digunakan lebih sedikit. Setelah semua peserta selesai diinput, program mencetak nama pemenang beserta jumlah soal dan waktu terbaiknya.</p>

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
<p align="justify">Program ini menampilkan dari deret bilangan yang memiliki polanya jika genap <code>1/2 dari suku sebelumnya</code>sedangkan ganjil <code>3*(suku sebelum)+1</code></p>
<p align="justify">pada func main hanya terdapat variable n sebagai tempat menyimpan nilai suku awal yang diinput oleh user. setelah menginputkan nilai awal terdapat pemanggilan procedure <code>deret ()</code> dengan parameter nya adalah nilai yang terdapat pada variable n. </p>
<p align="justify">saat procedure deret dipanggil pada func main, maka procedure tersebut akan menjalankan code yang terdapat di dalam procedure tersebut, saat procedure dijalankan ia akan menampilkan nilai variable n, selanjutnya terdapat perulangan dengan kondisi <code>n!=1 && n>0</code> yang berarti nilai variable n tidak sama dengan 1 dan harus lebih dari 0, selanjutnya jika kondisi itu bernilai true maka looping akan berjalan.   </p>
<p align="justify">di dalam looping terdapat if untuk pengecekan apakah nilai n merupakan nilai ganjil atau genap, jika nilai n genap akan dihitung sesuai rumus yaitu <code>n/2</code>, sedangkan saat dicek ternyata angka ganjil maka akan menggunakan rumus <code>3*n+1</code> yang hasil dari perhitungan akan disimpan pada varible n itu kembali. setelah if selesai nilai itu akan ditampilkan ke layar user. </p>