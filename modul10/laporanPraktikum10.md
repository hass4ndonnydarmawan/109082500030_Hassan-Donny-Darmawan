# <h1 align="center">Laporan Praktikum Modul 10 - ... </h1>
<p align="center">Hassan Donny Darmawan - 109082500030</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main
import "fmt"

const N = 1000
func hitungMinMax(berat[N] float64,n int, min *float64, max *float64){
*min = berat[0]
*max= berat[0]
for  i:= 1;  i<n; i++ {
	if berat[i]<*min {
		*min= berat[i]
	}

	if berat[i]>*max{
		*max= berat[i]
	}
}
}

func main (){
	var berat[N] float64
	var max, min float64
	var n int
	fmt.Scan(&n)

	for i:= 0; i<n; i++{
		fmt.Scan(&berat[i])
	}
	
	hitungMinMax(berat, n, &min,&max)
	fmt.Printf("%.2f %.2f\n",min, max )
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul10/output/outputsoal1.png)
[penjelasan]
<p align="justify">Program ini bertujuan untuk mencari nilai minimum dan maksimum dari sekumpulan data berat anak kelinci yang dimasukkan oleh pengguna. Program menggunakan sebuah array statis berkapasitas 1000 elemen untuk menampung seluruh data berat yang bertipe bilangan riil.</p>

<ol>

<li>Program mendefinisikan sebuah konstanta <i>N</i> dengan nilai 1000 sebagai batas kapasitas maksimal array dan mendeklarasikan variabel <i>berat</i> untuk menyimpan data berat kelinci dalam bentuk array berukuran <i>N</i>.</li>

<li>Fungsi <i>hitungMinMax(berat, n, min, max)</i> digunakan untuk mencari nilai terkecil dan terbesar di dalam array. Fungsi ini menerima array berat, jumlah data aktual (<i>n</i>), serta pointer ke variabel <i>min</i> dan <i>max</i> agar perubahan nilai di dalam fungsi dapat langsung mengubah variabel di fungsi utama (<i>pass by reference</i>).</li>

<li>Di dalam fungsi <i>hitungMinMax</i>, nilai awal <i>*min</i> dan <i>*max</i> diinisialisasi dengan elemen pertama array (<i>berat[0]</i>). Kemudian, program melakukan perulangan dari indeks ke-1 hingga <i>n-1</i> untuk membandingkan setiap data. Jika ditemukan nilai yang lebih kecil dari <i>*min</i>, maka nilai <i>*min</i> diperbarui. Begitu pula jika ditemukan nilai yang lebih besar dari <i>*max</i>, maka nilai <i>*max</i> akan diperbarui.</li>

<li>Pada fungsi <i>main</i>, program membaca input bilangan bulat <i>n</i> yang menyatakan jumlah anak kelinci. Setelah itu, program menggunakan perulangan untuk membaca input berat masing-masing kelinci sebanyak <i>n</i> kali dan menyimpannya ke dalam array.</li>

<li>Program kemudian memanggil fungsi <i>hitungMinMax</i> dengan mengirimkan alamat dari variabel <i>min</i> dan <i>max</i> menggunakan operator <i>&</i>.</li>

<li>Terakhir, program mencetak nilai berat kelinci terkecil dan terbesar dengan format dua angka di belakang koma (<i>%.2f</i>).</li>

</ol>

<p align="justify">Secara keseluruhan, program ini menerapkan konsep pencarian nilai ekstrem (minimum dan maksimum) pada struktur data array linear dengan memanfaatkan pointer untuk mengembalikan dua nilai sekaligus dari sebuah prosedur/fungsi.</p>

<p align="justify">Kritik untuk kode ini adalah penggunaan kapasitas array statis yang dipatok keras pada angka 1000. Jika pengguna memasukkan nilai N yang lebih besar dari 1000, program akan mengalami error <i>out of bounds</i> (index di luar batas). Selain itu, tidak ada pengecekan jika pengguna memasukkan data N sama dengan atau kurang dari 0, yang akan menyebabkan fungsi pencarian mengakses indeks memori yang tidak valid.</p>


### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

const N = 1000

func main() {
	var berat [N]float64
	var x, y int
	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	w := (x + y - 1) / y
	var tot [N]float64

	for i := 0; i < x; i++ {
		idx := i / y
		tot[idx] += berat[i]
	}

	for i := 0; i < w; i++ {
		fmt.Printf("%.2f", tot[i])
		if i < w-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	for i := 0; i < w; i++ {
		nIkan := y
		if i == w-1 && x%y != 0 {
			nIkan = x % y
		}
		rata := tot[i] / float64(nIkan)
		fmt.Printf("%.2f", rata)
		if i < w-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul10/output/outputsoal2.png)
[penjelasan]
<p align="justify">Program ini bertujuan untuk mensimulasikan pembagian sejumlah ikan ke dalam beberapa wadah berdasarkan kapasitas tertentu, lalu menghitung total berat serta rata-rata berat ikan yang ada di setiap wadah tersebut.</p>

<ol>

<li>Program menyiapkan array <i>berat</i> dengan kapasitas statis 1000 untuk menyimpan data berat dari setiap ikan yang diinput oleh pengguna.</li>

<li>Pada fungsi <i>main</i>, program membaca dua nilai input yaitu <i>x</i> (total jumlah ikan) dan <i>y</i> (kapasitas maksimal ikan dalam satu wadah).</li>

<li>Program melakukan perulangan sebanyak <i>x</i> kali untuk membaca satu per satu berat ikan dari baris kedua input dan menyimpannya ke dalam array <i>berat</i>.</li>

<li>Program menghitung jumlah wadah yang dibutuhkan menggunakan rumus pembulatan ke atas secara integer, yaitu <i>(x + y - 1) / y</i>, dan hasilnya disimpan dalam variabel <i>w</i>. Array baru bernama <i>tot</i> dideklarasikan untuk menyimpan total berat ikan pada masing-masing wadah.</li>

<li>Melalui perulangan, program menentukan posisi wadah untuk setiap ikan menggunakan operasi pembagian bilangan bulat <i>idx := i / y</i>. Berat ikan ke-<i>i</i> kemudian diakumulasikan ke dalam wadah dengan indeks <i>idx</i> tersebut.</li>

<li>Program mencetak total berat ikan untuk setiap wadah (dari wadah ke-0 hingga ke-<i>w-1</i>) pada baris pertama keluaran, dipisahkan oleh spasi.</li>

<li>Pada baris kedua keluaran, program menghitung rata-rata berat ikan per wadah. Program melakukan pengecekan khusus untuk wadah terakhir: jika total ikan tidak habis dibagi kapasitas wadah (<i>x % y != 0</i>), maka jumlah ikan di wadah terakhir adalah sisa pembagiannya (<i>x % y</i>), sedangkan wadah lainnya berisi penuh sebanyak <i>y</i> ikan. Nilai rata-rata didapat dari pembagian total berat wadah dengan jumlah ikan di wadah tersebut, lalu dicetak ke layar.</li>

</ol>

<p align="justify">Secara keseluruhan, program ini memadukan manipulasi indeks array dengan logika matematika (pembagian dan modulus) untuk mengelompokkan data linear ke dalam beberapa segmen kelompok (wadah).</p>

<p align="justify">Kekurangan fatal dari kode ini adalah penggunaan konstanta N=1000 pada array <i>tot</i> yang ukurannya dipaksakan sama dengan array berat. Jika jumlah wadah <i>w</i> secara tidak sengaja melebihi 1000 akibat input nilai pembagian yang tidak divalidasi, program akan langsung crash. Logika penentuan jumlah ikan di wadah terakhir juga ditulis secara berulang di dalam loop, yang sebenarnya bisa dioptimasi agar kodenya menjadi lebih bersih.</p>




### 3. [Soal]
#### soal3.go

```go
package main 
import "fmt"

type balita [100] float64

func hitungminmax(bb balita, n int, min *float64, max *float64){
	*min= bb[0]
	*max= bb[0]

	for  i:= 1;  i<n; i++ {
		if bb[i]<*min {
			*min= bb[i]
		}

		if bb[i]>*max{
			*max= bb[i]
		}
	}
}

func rerata (bb balita, n int)float64{
	total := 0.0
	for i:=0; i<n;i++{
		total += bb[i]
	}
	return total/float64(n)
}

func main(){
	var bb balita
	var n int
	var min, max float64
	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i:= 0; i<n;i++{
		fmt.Print("masukan berat balita ke - ",i+1," : ")
		fmt.Scan(&bb[i])
	}
	hitungminmax(bb,n,&min,&max)
	fmt.Printf("Berat balita minimum: %.2f kg \n",min)
	fmt.Printf("Berat balita maximum: %.2f kg \n",max)
	fmt.Printf("Rerata berat balita: %.2f kg", rerata(bb,n))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul10/output/outputsoal3.png)
[penjelasan]
<p align="justify">Program ini dibuat untuk membantu petugas Posyandu dalam mencatat data berat badan sejumlah balita, menampungnya ke dalam tipe data bentukan (array terstruktur), kemudian menyajikan informasi berupa berat badan terkecil, terbesar, serta rata-rata dari seluruh balita tersebut.</p>

<ol>

<li>Program mendefinisikan sebuah tipe data baru (alias) bernama <i>balita</i>, yang merupakan array bertipe <i>float64</i> dengan kapasitas statis sebanyak 100 elemen.</li>

<li>Fungsi <i>hitungminmax(bb, n, min, max)</i> bertugas mencari nilai berat badan ekstrem (paling kecil dan paling besar). Fungsi ini menerima array dengan tipe data <i>balita</i>, jumlah data yang valid (<i>n</i>), serta pointer ke variabel penampung hasil (<i>min</i> dan <i>max</i>). Proses pencarian dilakukan dengan membandingkan elemen awal ke seluruh elemen array berikutnya secara sekuensial.</li>

<li>Fungsi <i>rerata(bb, n)</i> dideklarasikan sebagai fungsi yang mengembalikan nilai riil (<i>float64</i>). Fungsi ini melakukan perulangan untuk menjumlahkan seluruh berat balita yang ada, kemudian membagi total tersebut dengan jumlah balita (<i>n</i>) untuk mendapatkan nilai rata-rata.</li>

<li>Di dalam fungsi <i>main</i>, program berinteraksi dengan pengguna secara interaktif dengan menampilkan panduan teks cetak (<i>prompting</i>) sebelum membaca input jumlah data (<i>n</i>).</li>

<li>Menggunakan struktur perulangan, program meminta petugas memasukkan berat badan balita satu per satu dari balita ke-1 hingga ke-<i>n</i>, lalu menyimpannya ke dalam variabel array <i>bb</i> yang bertipe <i>balita</i>.</li>

<li>Program mengeksekusi prosedur <i>hitungminmax</i> untuk mendapatkan berat terkecil dan terbesar, lalu mencetak hasilnya ke layar beserta hasil dari fungsi <i>rerata</i> dengan format dua angka desimal di belakang koma lengkap dengan satuan "kg".</li>

</ol>

<p align="justify">Secara keseluruhan, program ini sudah menerapkan modularitas yang baik dengan memisahkan subprogram berdasarkan tugas spesifiknya (pencarian nilai ekstrem dan penghitungan rata-rata) serta mengimplementasikan tipe data bentukan buatan sendiri.</p>

<p align="justify">Namun, program ini masih memiliki kelemahan dalam aspek interaksi dan validasi. Tipe data array <i>balita</i> dikunci pada angka 100, sehingga jika petugas menginput jumlah balita lebih dari 100, program akan rusak saat runtime. Program juga belum mengantisipasi jika petugas melakukan salah input berupa teks/huruf saat meminta input berat badan balita, yang bisa menyebabkan program langsung berhenti bekerja tanpa menampilkan pesan error yang informatif.</p>