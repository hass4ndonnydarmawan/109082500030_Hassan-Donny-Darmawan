# <h1 align="center">Laporan Praktikum Modul 12 - ... </h1>
<p align="center">Hassan Donny Darmawan - 109082500030</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func sequential(suara []int, target int) int {
	for i := 0; i < 20; i++ {
		if i+1 == target {
			return suara[i]
		}
	}
	return -1
}

func main() {
	var n int
	total := 0
	sah := 0
	data := [100]int{}
	suara := [20]int{}

	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		data[total] = n
		total++
	}

	for i := 0; i < total; i++ {
		if data[i] >= 1 && data[i] <= 20 {
			sah++
			suara[data[i]-1]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", total)
	fmt.Printf("Suara sah: %d\n", sah)

	for i := 1; i <= 20; i++ {
		hasil := sequential(suara[:], i)
		if hasil > 0 {
			fmt.Printf("%d: %d\n", i, hasil)
		}
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul12/output/output1.png)
[penjelasan]
<p align="justify">Program ini bertujuan untuk menghitung hasil pemungutan suara dari 20 kandidat. Program menerima masukan berupa nomor kandidat yang dipilih oleh pemilih, kemudian menghitung jumlah suara masuk, jumlah suara sah, dan total suara yang diperoleh masing-masing kandidat.</p>

<ol>

<li>Program mendeklarasikan fungsi <i>sequential(suara []int, target int)</i> yang digunakan untuk mencari jumlah suara suatu kandidat berdasarkan nomor urut kandidat. Fungsi ini melakukan pencarian secara sekuensial (linear search) dari kandidat nomor 1 sampai 20. Jika nomor kandidat yang dicari sesuai dengan nilai <i>target</i>, fungsi akan mengembalikan jumlah suara kandidat tersebut. Jika tidak ditemukan, fungsi mengembalikan nilai <i>-1</i>.</li>

<li>Pada fungsi <i>main</i>, program mendeklarasikan beberapa variabel, yaitu <i>n</i> untuk menyimpan input sementara, <i>total</i> untuk menghitung jumlah seluruh suara yang masuk, <i>sah</i> untuk menghitung jumlah suara sah, array <i>data</i> berukuran 100 elemen untuk menyimpan seluruh suara yang masuk, dan array <i>suara</i> berukuran 20 elemen untuk menyimpan jumlah suara setiap kandidat.</li>

<li>Program kemudian membaca data suara menggunakan perulangan tak hingga (<i>for</i>). Setiap nomor kandidat yang dimasukkan pengguna akan disimpan ke dalam array <i>data</i>. Proses input berhenti ketika pengguna memasukkan angka <i>0</i> sebagai tanda akhir data.</li>

<li>Setelah seluruh data masuk, program melakukan perulangan pada array <i>data</i> untuk memeriksa validitas setiap suara. Suara dianggap sah apabila nomor kandidat berada pada rentang 1 sampai 20.</li>

<li>Jika suara valid, variabel <i>sah</i> ditambah satu dan elemen array <i>suara[data[i]-1]</i> dinaikkan satu. Pengurangan satu dilakukan karena indeks array dimulai dari 0, sedangkan nomor kandidat dimulai dari 1.</li>

<li>Program kemudian menampilkan jumlah suara masuk yang tersimpan pada variabel <i>total</i> dan jumlah suara sah yang tersimpan pada variabel <i>sah</i>.</li>

<li>Untuk menampilkan hasil perolehan setiap kandidat, program melakukan perulangan dari kandidat nomor 1 sampai 20. Pada setiap iterasi, fungsi <i>sequential</i> dipanggil untuk memperoleh jumlah suara kandidat yang sedang diproses.</li>

<li>Jika jumlah suara yang diperoleh kandidat lebih dari 0, program akan mencetak nomor kandidat beserta jumlah suaranya. Kandidat yang tidak memperoleh suara tidak akan ditampilkan.</li>

</ol>

<p align="justify">Secara keseluruhan, program ini menerapkan konsep array, validasi data, pencacahan frekuensi (frequency counting), dan pencarian sekuensial (<i>sequential search</i>) untuk mengolah hasil pemungutan suara sederhana. Setiap suara yang sah dihitung dan dikelompokkan berdasarkan nomor kandidat sehingga menghasilkan rekapitulasi jumlah suara per kandidat.</p>

<p align="justify">Kritik untuk kode ini adalah fungsi <i>sequential</i> sebenarnya tidak melakukan pencarian terhadap isi array, melainkan langsung mengakses indeks berdasarkan nomor kandidat yang dicari. Akibatnya, penggunaan <i>sequential search</i> menjadi kurang efektif karena hanya menambah proses yang sebenarnya dapat digantikan dengan akses langsung ke array. Selain itu, array <i>data</i> memiliki kapasitas tetap sebesar 100 elemen tanpa pengecekan batas. Jika jumlah suara yang dimasukkan melebihi 100 data, program dapat mengalami <i>out of bounds error</i>. Program juga tidak memberikan informasi mengenai jumlah suara tidak sah sehingga sebagian data input dapat terabaikan tanpa diketahui pengguna.</p>

### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

func cariKetua(suara []int) int {
	pos := 0
	for i := 1; i < 20; i++ {
		if suara[i] > suara[pos] {
			pos = i
		}
	}
	return pos
}

func cariWakil(suara []int, posKetua int) int {
	pos := -1
	for i := 0; i < 20; i++ {
		if i == posKetua {
			continue
		}
		if pos == -1 || suara[i] > suara[pos] {
			pos = i
		}
	}
	return pos
}

func main() {
	var n int
	total := 0
	sah := 0
	suara := [20]int{}

	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		total++
		if n >= 1 && n <= 20 {
			sah++
			suara[n-1]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", total)
	fmt.Printf("Suara sah: %d\n", sah)

	ketua := cariKetua(suara[:])
	wakil := cariWakil(suara[:], ketua)

	fmt.Printf("Ketua RT: %d\n", ketua+1)
	fmt.Printf("Wakil ketua: %d\n", wakil+1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul12/output/output2.png)
[penjelasan]
<p align="justify">Program ini bertujuan untuk menentukan Ketua RT dan Wakil Ketua RT berdasarkan hasil pemungutan suara dari 20 kandidat. Program menerima masukan berupa nomor kandidat yang dipilih oleh pemilih, kemudian menghitung jumlah suara yang diterima setiap kandidat dan menentukan kandidat dengan suara terbanyak sebagai Ketua RT serta kandidat dengan suara terbanyak kedua sebagai Wakil Ketua RT.</p>

<ol>

<li>Program mendefinisikan fungsi <i>cariKetua(suara []int)</i> yang digunakan untuk mencari kandidat dengan jumlah suara terbanyak. Fungsi ini menggunakan variabel <i>pos</i> untuk menyimpan indeks kandidat yang sementara memiliki suara tertinggi.</li>

<li>Di dalam fungsi <i>cariKetua</i>, program melakukan perulangan dari indeks ke-1 hingga ke-19. Jika ditemukan kandidat dengan jumlah suara lebih besar daripada kandidat yang saat ini disimpan pada <i>pos</i>, maka nilai <i>pos</i> diperbarui. Setelah perulangan selesai, fungsi mengembalikan indeks kandidat dengan suara terbanyak.</li>

<li>Program juga mendefinisikan fungsi <i>cariWakil(suara []int, posKetua int)</i> yang digunakan untuk mencari kandidat dengan suara terbanyak kedua. Parameter <i>posKetua</i> berisi indeks kandidat yang telah terpilih sebagai Ketua RT sehingga tidak ikut dibandingkan kembali.</li>

<li>Di dalam fungsi <i>cariWakil</i>, variabel <i>pos</i> diinisialisasi dengan nilai <i>-1</i>. Program kemudian melakukan perulangan pada seluruh kandidat. Jika indeks yang sedang diperiksa sama dengan indeks Ketua RT, maka iterasi dilewati menggunakan perintah <i>continue</i>.</li>

<li>Untuk kandidat selain Ketua RT, program membandingkan jumlah suaranya dengan kandidat yang saat ini tersimpan pada <i>pos</i>. Jika jumlah suaranya lebih besar, maka <i>pos</i> diperbarui. Setelah seluruh kandidat diperiksa, fungsi mengembalikan indeks kandidat dengan suara terbanyak kedua sebagai Wakil Ketua RT.</li>

<li>Pada fungsi <i>main</i>, program mendeklarasikan variabel <i>n</i> untuk menyimpan input suara, <i>total</i> untuk menghitung jumlah seluruh suara yang masuk, <i>sah</i> untuk menghitung jumlah suara sah, dan array <i>suara</i> berukuran 20 elemen untuk menyimpan jumlah suara masing-masing kandidat.</li>

<li>Program membaca data suara menggunakan perulangan tak hingga (<i>for</i>). Input berhenti ketika pengguna memasukkan angka <i>0</i>. Setiap suara yang masuk akan menambah nilai <i>total</i>.</li>

<li>Program melakukan validasi terhadap setiap suara. Jika nomor kandidat berada pada rentang 1 sampai 20, maka suara dianggap sah, variabel <i>sah</i> ditambah satu, dan elemen array yang sesuai pada <i>suara</i> akan ditambah satu untuk mencatat perolehan suara kandidat tersebut.</li>

<li>Setelah seluruh data selesai dibaca, program menampilkan jumlah suara masuk dan jumlah suara sah.</li>

<li>Program kemudian memanggil fungsi <i>cariKetua</i> untuk memperoleh indeks kandidat dengan suara terbanyak dan menyimpannya ke dalam variabel <i>ketua</i>. Selanjutnya, fungsi <i>cariWakil</i> dipanggil untuk memperoleh indeks kandidat dengan suara terbanyak kedua dan menyimpannya ke dalam variabel <i>wakil</i>.</li>

<li>Terakhir, program menampilkan nomor kandidat yang terpilih sebagai Ketua RT dan Wakil Ketua RT. Karena indeks array dimulai dari 0 sedangkan nomor kandidat dimulai dari 1, maka program menambahkan nilai 1 pada indeks sebelum ditampilkan.</li>

</ol>

<p align="justify">Secara keseluruhan, program ini menerapkan konsep array, pencacahan frekuensi (<i>frequency counting</i>), pencarian nilai maksimum, dan pencarian nilai maksimum kedua untuk menentukan peringkat kandidat berdasarkan jumlah suara yang diperoleh. Metode ini cukup efisien karena hanya memerlukan beberapa kali penelusuran linear pada array berukuran tetap.</p>

<p align="justify">Kritik untuk kode ini adalah program belum menangani kondisi seri (<i>tie</i>) apabila terdapat dua atau lebih kandidat yang memperoleh jumlah suara yang sama. Dalam kondisi tersebut, program akan memilih kandidat yang muncul lebih dahulu pada array. Selain itu, program tetap akan menentukan Ketua RT dan Wakil Ketua RT meskipun tidak ada suara sah yang masuk, sehingga hasil yang ditampilkan dapat menjadi kurang bermakna. Program juga tidak menampilkan jumlah suara yang diperoleh oleh Ketua RT maupun Wakil Ketua RT sehingga informasi hasil pemilihan masih terbatas.</p>

### 3. [Soal]
#### soal3.go

```go
package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	kiri := 0
	kanan := n - 1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if data[tengah] == k {
			return tengah
		} else if data[tengah] < k {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	isiArray(n)
	hasil := posisi(n, k)
	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul12/output/output3.png)
[penjelasan]
<p align="justify">Program ini bertujuan untuk mencari posisi suatu nilai pada sebuah array bilangan bulat yang telah terurut menggunakan algoritma <i>Binary Search</i>. Jika nilai yang dicari ditemukan, program akan menampilkan indeks posisinya. Sebaliknya, jika nilai tersebut tidak terdapat di dalam array, program akan menampilkan pesan "TIDAK ADA".</p>

<ol>

<li>Program mendefinisikan konstanta <i>NMAX</i> dengan nilai 1.000.000 yang digunakan sebagai kapasitas maksimum array. Selanjutnya, program mendeklarasikan array global <i>data</i> bertipe <i>int</i> dengan ukuran <i>NMAX</i> untuk menyimpan data yang akan dicari.</li>

<li>Fungsi <i>isiArray(n int)</i> digunakan untuk mengisi array <i>data</i> sebanyak <i>n</i> elemen. Fungsi ini melakukan perulangan dari indeks 0 hingga <i>n-1</i> dan membaca setiap nilai yang dimasukkan pengguna menggunakan <i>fmt.Scan()</i>.</li>

<li>Program mendefinisikan fungsi <i>posisi(n, k int)</i> yang bertugas mencari nilai <i>k</i> di dalam array menggunakan metode <i>Binary Search</i>. Fungsi menerima dua parameter, yaitu jumlah elemen array (<i>n</i>) dan nilai yang akan dicari (<i>k</i>).</li>

<li>Di dalam fungsi <i>posisi</i>, variabel <i>kiri</i> diinisialisasi dengan nilai 0 sebagai batas awal pencarian, sedangkan variabel <i>kanan</i> diinisialisasi dengan nilai <i>n-1</i> sebagai batas akhir pencarian.</li>

<li>Program melakukan perulangan selama nilai <i>kiri</i> masih lebih kecil atau sama dengan <i>kanan</i>. Pada setiap iterasi, program menghitung indeks tengah menggunakan rumus <i>(kiri + kanan) / 2</i>.</li>

<li>Jika nilai pada indeks tengah sama dengan nilai yang dicari (<i>data[tengah] == k</i>), maka fungsi langsung mengembalikan indeks tersebut sebagai hasil pencarian.</li>

<li>Jika nilai pada indeks tengah lebih kecil daripada nilai yang dicari, maka pencarian dilanjutkan ke bagian kanan array dengan memperbarui nilai <i>kiri</i> menjadi <i>tengah + 1</i>.</li>

<li>Sebaliknya, jika nilai pada indeks tengah lebih besar daripada nilai yang dicari, maka pencarian dilanjutkan ke bagian kiri array dengan memperbarui nilai <i>kanan</i> menjadi <i>tengah - 1</i>.</li>

<li>Apabila seluruh proses pencarian selesai dan nilai yang dicari tidak ditemukan, fungsi akan mengembalikan nilai <i>-1</i> sebagai penanda bahwa data tidak ada di dalam array.</li>

<li>Pada fungsi <i>main</i>, program membaca input berupa jumlah data <i>n</i> dan nilai yang akan dicari <i>k</i>. Setelah itu, program memanggil fungsi <i>isiArray</i> untuk mengisi array dengan data yang diberikan pengguna.</li>

<li>Selanjutnya, program memanggil fungsi <i>posisi</i> dan menyimpan hasilnya ke dalam variabel <i>hasil</i>.</li>

<li>Jika nilai <i>hasil</i> sama dengan <i>-1</i>, program menampilkan pesan "TIDAK ADA". Jika tidak, program menampilkan indeks tempat nilai tersebut ditemukan di dalam array.</li>

</ol>

<p align="justify">Secara keseluruhan, program ini menerapkan algoritma <i>Binary Search</i> untuk melakukan pencarian data pada array terurut. Dibandingkan dengan <i>Sequential Search</i>, metode ini jauh lebih efisien karena setiap langkah selalu mempersempit ruang pencarian menjadi setengah bagian. Kompleksitas waktunya adalah <i>O(log n)</i>, sehingga sangat cocok digunakan untuk jumlah data yang besar.</p>

<p align="justify">Kritik untuk kode ini adalah program mengasumsikan bahwa data yang dimasukkan sudah dalam keadaan terurut menaik. Jika array tidak terurut, hasil pencarian dapat menjadi salah atau data yang sebenarnya ada tidak dapat ditemukan. Selain itu, tidak terdapat validasi terhadap nilai <i>n</i>. Jika pengguna memasukkan nilai <i>n</i> yang lebih besar dari <i>NMAX</i> atau kurang dari 1, program dapat mengalami perilaku yang tidak diharapkan. Program juga hanya mengembalikan satu posisi data yang ditemukan. Jika terdapat beberapa elemen dengan nilai yang sama, posisi yang dikembalikan belum tentu merupakan kemunculan pertama ataupun terakhir dari nilai tersebut.</p>