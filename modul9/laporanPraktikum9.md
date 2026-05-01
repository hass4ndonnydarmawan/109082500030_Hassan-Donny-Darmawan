# <h1 align="center">Laporan Praktikum Modul 9 - ... </h1>
<p align="center">Hassan Donny Darmawan - 109082500030</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main 
import (
	"fmt"
	"math"
)

type pusat struct {
	cx,cy,r float64
}
type input [2]pusat

func jarak (a,b,c,d float64)float64 {
	return math.Sqrt(math.Pow((a-c),2) + math.Pow((b-d),2))
}

func didalam (cx,cy,r,x,y float64) bool {
	return jarak(cx,cy,x,y) <= r
}

func main (){
	var (
		masukan input
		x,y float64
	)

	fmt.Scan(&masukan[0].cx, &masukan[0].cy, &masukan[0].r)
    fmt.Scan(&masukan[1].cx, &masukan[1].cy, &masukan[1].r)
    fmt.Scan(&x, &y)

	d1 := didalam(masukan[0].cx,masukan[0].cy,masukan[0].r,x,y)
	d2 := didalam(masukan[1].cx,masukan[1].cy,masukan[1].r,x,y)

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
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul9/output/outputno1.png)
[penjelasan]
<p align="justify">Program ini bertujuan untuk menentukan posisi sebuah titik terhadap dua lingkaran berdasarkan input yang diberikan oleh pengguna. Program akan mengecek apakah titik tersebut berada di dalam salah satu lingkaran, kedua lingkaran, atau di luar keduanya.</p>

<ol>

<li>Program mendefinisikan tipe data <i>pusat</i> yang berisi koordinat titik pusat lingkaran (<i>cx</i>, <i>cy</i>) serta jari-jari <i>r</i>. Selain itu, dibuat tipe <i>input</i> berupa array yang menampung dua buah lingkaran.</li>

<li>Fungsi <i>jarak(a, b, c, d)</i> digunakan untuk menghitung jarak antara dua titik pada bidang koordinat menggunakan rumus Euclidean, yaitu akar dari jumlah kuadrat selisih koordinat.</li>

<li>Fungsi <i>didalam(cx, cy, r, x, y)</i> berfungsi untuk mengecek apakah titik (<i>x</i>, <i>y</i>) berada di dalam lingkaran dengan pusat (<i>cx</i>, <i>cy</i>) dan jari-jari <i>r</i>. Kondisi terpenuhi jika jarak titik ke pusat lingkaran kurang dari atau sama dengan jari-jari.</li>

<li>Pada fungsi <i>main</i>, program menerima input dua lingkaran (masing-masing terdiri dari koordinat pusat dan jari-jari), serta satu titik yang akan diuji posisinya.</li>

<li>Program kemudian memanggil fungsi <i>didalam</i> untuk masing-masing lingkaran, dan hasilnya disimpan dalam variabel <i>d1</i> dan <i>d2</i>.</li>

<li>Berdasarkan nilai <i>d1</i> dan <i>d2</i>, program menentukan posisi titik:
	<ul>
	<li>Jika kedua nilai bernilai benar, maka titik berada di dalam kedua lingkaran.</li>
	<li>Jika hanya <i>d1</i> yang benar, maka titik berada di dalam lingkaran pertama.</li>
	<li>Jika hanya <i>d2</i> yang benar, maka titik berada di dalam lingkaran kedua.</li>
	<li>Jika keduanya salah, maka titik berada di luar kedua lingkaran.</li>
	</ul>
</li>

</ol>

<p align="justify">Secara keseluruhan, program ini menggabungkan konsep geometri (jarak antar titik) dengan logika percabangan untuk menentukan posisi relatif sebuah titik terhadap dua lingkaran.</p>

<p align="justify">Masalahnya, kamu masih bermain di level dasar tanpa pengamanan apa pun. Tidak ada validasi input (misalnya jari-jari negatif), tidak ada penanganan error input, dan struktur datanya juga terlalu kaku (hanya 2 lingkaran). Kalau ini mau naik level, kamu harus mulai berpikir fleksibel: gunakan slice, tambahkan validasi, dan pisahkan logika dari input/output supaya bisa diuji. Sekarang ini masih sekadar program latihan, belum siap dipakai di dunia nyata.</p>


### 2. [Soal]
#### soal2.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul9/output/outputno2.png)
[penjelasan]
<p align="justify">Program ini bertujuan untuk mengelola dan melakukan berbagai operasi pada sebuah array bilangan bulat yang diinput oleh pengguna. Operasi yang dilakukan meliputi penampilan elemen, penghapusan data, serta perhitungan statistik sederhana seperti rata-rata dan standar deviasi.</p>
<ol>
<li>Program mendeklarasikan sebuah array dengan kapasitas maksimal 100 elemen serta variabel <i>n</i> untuk menyimpan jumlah elemen yang akan digunakan. Pengguna diminta memasukkan jumlah elemen, lalu program akan memvalidasi agar nilainya tidak melebihi kapasitas atau kurang dari sama dengan nol.</li>
<li>Setelah itu, pengguna diminta mengisi array dengan <i>n</i> bilangan bulat. Data yang dimasukkan akan disimpan ke dalam array.</li>
<li>Program menampilkan seluruh isi array sesuai dengan jumlah elemen yang dimasukkan.</li>
<li>Program menampilkan elemen dengan indeks ganjil, yaitu elemen pada posisi 1, 3, 5, dan seterusnya.</li>
<li>Program menampilkan elemen dengan indeks genap, yaitu elemen pada posisi 0, 2, 4, dan seterusnya.</li>
<li>Program meminta input nilai <i>x</i>, kemudian menampilkan elemen yang berada pada indeks kelipatan <i>x</i>. Jika <i>x</i> bernilai nol, maka proses ini tidak dijalankan untuk menghindari pembagian dengan nol.</li>
<li>Program meminta indeks yang ingin dihapus, lalu menggeser elemen di sebelah kanan indeks tersebut ke kiri untuk menutupi posisi yang dihapus. Setelah itu, jumlah elemen <i>n</i> dikurangi satu.</li>
<li>Program menghitung rata-rata dari seluruh elemen array dengan menjumlahkan semua nilai lalu membaginya dengan jumlah elemen.</li>
<li>Program menghitung standar deviasi (simpangan baku) menggunakan rumus akar dari rata-rata kuadrat selisih setiap elemen terhadap nilai rata-rata.</li>
<li>Program meminta sebuah bilangan untuk dicari frekuensinya, kemudian menghitung berapa kali bilangan tersebut muncul di dalam array.</li>
</ol>
<p align="justify">Pada fungsi utama (<i>main</i>), seluruh proses dilakukan secara berurutan mulai dari input data, pengolahan array, hingga perhitungan statistik, dan hasilnya ditampilkan langsung ke layar pengguna.</p>
<p align="justify">Namun, terdapat kelemahan pada kode ini yaitu belum adanya validasi terhadap indeks yang dihapus (bisa menyebabkan error jika indeks di luar batas), serta penggunaan <i>strings.Repeat</i> tanpa melakukan import package <i>strings</i>. Hal ini perlu diperbaiki agar program lebih aman dan dapat dijalankan tanpa error.</p>




### 3. [Soal]
#### soal3.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul9/output/outputno3.png)
[penjelasan]
<p align="justify">Program ini bertujuan untuk mencatat hasil beberapa pertandingan antara dua klub, lalu menampilkan pemenang dari setiap pertandingan berdasarkan skor yang dimasukkan oleh pengguna.</p>
<ol>
<li>Program mendeklarasikan dua variabel bertipe string, yaitu <i>klubA</i> dan <i>klubB</i>, untuk menyimpan nama kedua klub yang akan bertanding. Selain itu, digunakan slice <i>hasilPertandingan</i> untuk menyimpan hasil setiap pertandingan.</li>
<li>Pengguna diminta memasukkan nama kedua klub. Nilai ini akan digunakan sebagai referensi dalam menentukan pemenang pada setiap pertandingan.</li>
<li>Program menggunakan perulangan tanpa batas (<i>infinite loop</i>) untuk menerima input skor pertandingan secara terus-menerus. Variabel <i>pertandinganKe</i> digunakan untuk menandai urutan pertandingan.</li>
<li>Pada setiap iterasi, pengguna diminta memasukkan skor untuk klub A dan klub B. Jika salah satu skor bernilai negatif, maka perulangan akan dihentikan menggunakan <i>break</i>. Ini menjadi sinyal bahwa input pertandingan telah selesai.</li>
<li>Program membandingkan skor kedua klub:
	<ul>
	<li>Jika skor klub A lebih besar, maka klub A dinyatakan sebagai pemenang.</li>
	<li>Jika skor klub B lebih besar, maka klub B menjadi pemenang.</li>
	<li>Jika kedua skor sama, maka hasil pertandingan dinyatakan <i>Draw</i>.</li>
	</ul>
	Hasil ini kemudian disimpan ke dalam slice <i>hasilPertandingan</i>.</li>
<li>Setelah perulangan selesai, program menampilkan seluruh hasil pertandingan yang telah disimpan dalam slice, lengkap dengan nomor urutnya.</li>
<li>Terakhir, program mencetak pesan bahwa seluruh pertandingan telah selesai diproses.</li>
</ol>
<p align="justify">Secara fungsi, program ini sudah berjalan, tetapi jangan cepat puas—ini masih mentah. Kamu tidak melakukan validasi input sama sekali (misalnya input non-angka akan langsung merusak alur program), tidak ada pembatasan jumlah pertandingan (user bisa memasukkan data tanpa kontrol), dan tidak ada ringkasan hasil (misalnya total kemenangan masing-masing klub). Lebih parah lagi, kamu hanya menyimpan pemenang tanpa menyimpan skor—artinya kamu kehilangan data penting yang bisa dipakai untuk analisis lebih lanjut.</p>
<p align="justify">Kalau kamu serius ingin naik level, ubah cara berpikirmu: jangan hanya membuat program “yang jalan”, tapi program yang tahan terhadap kesalahan input, punya struktur data yang lebih kaya (misalnya struct untuk menyimpan skor dan pemenang), dan mampu menghasilkan insight (statistik kemenangan, jumlah draw, dll). Sekarang ini kamu masih di tahap mencatat, belum menganalisis.</p>


### 4. [Soal]
#### soal4.go

```go
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0
	for {
		fmt.Scanf("%c", &ch)
		if ch == '.' {
			break
		}
		if ch != '\n' && ch != ' ' {
			t[*n] = ch
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks : ")
	isiArray(&tab, &m)

	fmt.Print("Reverse teks : ")
	balikanArray(&tab, m)
	cetakArray(tab, m)

	balikanArray(&tab, m)

	fmt.Print("Palindrom ? ")
	fmt.Println(palindrom(tab, m))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul9/output/outputno4.png)
[penjelasan]
<p align="justify">Program ini bertujuan untuk membaca sekumpulan karakter dari input pengguna, menampilkan hasil pembalikan urutan karakter tersebut, serta menentukan apakah teks yang dimasukkan merupakan palindrom atau bukan.</p>
<ol>
<li>Program mendeklarasikan konstanta <i>NMAX</i> sebagai kapasitas maksimum array, serta tipe data <i>tabel</i> berupa array karakter (<i>rune</i>) dengan ukuran tetap.</li>
<li>Fungsi <i>isiArray</i> digunakan untuk membaca karakter satu per satu dari input pengguna hingga ditemukan tanda titik ('.') sebagai penanda akhir input. Karakter spasi dan newline diabaikan agar tidak mempengaruhi isi array. Setiap karakter valid disimpan ke dalam array dan jumlah elemen <i>n</i> diperbarui.</li>
<li>Fungsi <i>cetakArray</i> bertugas menampilkan isi array ke layar sesuai jumlah elemen yang telah diinput, dengan mencetak setiap karakter secara berurutan.</li>
<li>Fungsi <i>balikanArray</i> digunakan untuk membalik urutan elemen dalam array. Proses dilakukan dengan menukar elemen dari depan dan belakang secara bertahap hingga mencapai tengah array.</li>
<li>Fungsi <i>palindrom</i> berfungsi untuk memeriksa apakah array membentuk palindrom. Pemeriksaan dilakukan dengan membandingkan elemen dari awal dan akhir secara berpasangan. Jika semua pasangan sama, maka fungsi mengembalikan nilai <i>true</i>, jika tidak maka <i>false</i>.</li>
<li>Pada fungsi <i>main</i>, program meminta pengguna memasukkan teks, lalu memanggil fungsi <i>isiArray</i> untuk menyimpan karakter ke dalam array.</li>
<li>Program kemudian membalik isi array menggunakan <i>balikanArray</i> dan menampilkan hasilnya sebagai teks terbalik.</li>
<li>Setelah itu, array dikembalikan ke kondisi semula dengan memanggil kembali <i>balikanArray</i>, sehingga data asli tetap digunakan untuk proses berikutnya.</li>
<li>Program memanggil fungsi <i>palindrom</i> untuk mengecek apakah teks yang dimasukkan merupakan palindrom, lalu menampilkan hasilnya ke layar.</li>
</ol>
<p align="justify">Pada fungsi utama (<i>main</i>), seluruh proses dilakukan secara berurutan mulai dari input karakter, pembalikan array, hingga pengecekan palindrom, dan hasilnya ditampilkan langsung kepada pengguna.</p>
<p align="justify">Meskipun program ini sudah berjalan dengan baik, masih terdapat kelemahan yaitu tidak adanya pembatasan jumlah input agar tidak melebihi kapasitas array, serta tidak adanya validasi tambahan terhadap karakter selain spasi dan newline. Jika input terlalu panjang atau tidak terkontrol, program berpotensi mengalami kesalahan. Untuk pengembangan lebih lanjut, sebaiknya ditambahkan validasi batas array dan penanganan input yang lebih ketat.</p>
