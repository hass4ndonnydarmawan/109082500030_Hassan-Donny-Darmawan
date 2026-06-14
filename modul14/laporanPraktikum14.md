# <h1 align="center">Laporan Praktikum Modul 14 - ... </h1>
<p align="center">Hassan Donny Darmawan - 109082500030</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

const NMAX = 1000000

var arr [NMAX]int

func selectionSort(n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		temp := arr[i]
		arr[i] = arr[idxMin]
		arr[idxMin] = temp
	}
}

func main() {
	var n, m int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)

		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		selectionSort(m)

		for j := 0; j < m; j++ {
			fmt.Print(arr[j])
			if j < m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul14/output/soal1.png)
[penjelasan]
<p align="justify">Program ini bertujuan untuk mengurutkan sejumlah data bilangan bulat menggunakan algoritma <i>Selection Sort</i>. Program menerima beberapa kumpulan data (test case), kemudian mengurutkan setiap kumpulan data dari nilai terkecil ke terbesar dan menampilkan hasil pengurutannya.</p>

<ol>

<li>Program mendeklarasikan konstanta <i>NMAX</i> bernilai 1000000 yang digunakan sebagai kapasitas maksimum array. Selain itu, dideklarasikan array global <i>arr</i> bertipe integer untuk menyimpan data yang akan diurutkan.</li>

<li>Program mendefinisikan fungsi <i>selectionSort(n int)</i> yang digunakan untuk mengurutkan data dalam array <i>arr</i> sebanyak <i>n</i> elemen menggunakan metode <i>Selection Sort</i>.</li>

<li>Pada fungsi <i>selectionSort</i>, perulangan luar menggunakan variabel <i>i</i> dari indeks pertama hingga indeks ke-<i>n-2</i>. Setiap iterasi bertujuan menentukan nilai terkecil yang akan ditempatkan pada posisi ke-<i>i</i>.</li>

<li>Variabel <i>idxMin</i> diinisialisasi dengan nilai <i>i</i> sebagai indeks sementara dari elemen terkecil yang ditemukan pada bagian array yang belum terurut.</li>

<li>Perulangan dalam menggunakan variabel <i>j</i> untuk menelusuri elemen setelah indeks <i>i</i>. Jika ditemukan elemen yang lebih kecil daripada <i>arr[idxMin]</i>, maka nilai <i>idxMin</i> diperbarui sehingga selalu menyimpan posisi elemen terkecil.</li>

<li>Setelah seluruh elemen pada bagian yang belum terurut diperiksa, program menukar (<i>swap</i>) nilai pada indeks <i>i</i> dengan nilai pada indeks <i>idxMin</i>. Proses ini menempatkan elemen terkecil ke posisi yang seharusnya.</li>

<li>Pada fungsi <i>main</i>, program mendeklarasikan variabel <i>n</i> untuk menyimpan jumlah test case dan variabel <i>m</i> untuk menyimpan jumlah elemen pada setiap test case.</li>

<li>Program membaca nilai <i>n</i> yang menunjukkan berapa banyak kumpulan data yang akan diproses.</li>

<li>Melalui perulangan sebanyak <i>n</i> kali, program membaca nilai <i>m</i> sebagai jumlah bilangan pada test case yang sedang diproses.</li>

<li>Selanjutnya, program membaca <i>m</i> buah bilangan dan menyimpannya ke dalam array <i>arr</i> mulai dari indeks 0 hingga indeks <i>m-1</i>.</li>

<li>Setelah seluruh data pada test case masuk, fungsi <i>selectionSort(m)</i> dipanggil untuk mengurutkan elemen-elemen tersebut secara menaik (<i>ascending</i>).</li>

<li>Hasil pengurutan kemudian ditampilkan menggunakan perulangan. Setiap elemen dicetak dan dipisahkan dengan spasi agar format keluaran lebih rapi.</li>

<li>Setelah seluruh elemen pada satu test case ditampilkan, program mencetak baris baru dan melanjutkan proses ke test case berikutnya hingga seluruh data selesai diproses.</li>

</ol>

<p align="justify">Secara keseluruhan, program ini menerapkan konsep array, perulangan bersarang (<i>nested loop</i>), pertukaran data (<i>swap</i>), serta algoritma <i>Selection Sort</i> untuk mengurutkan data bilangan bulat. Setiap test case diproses secara terpisah sehingga menghasilkan keluaran berupa urutan angka dari yang terkecil hingga terbesar.</p>

<p align="justify">Kritik untuk kode ini adalah penggunaan algoritma <i>Selection Sort</i> kurang efisien untuk jumlah data yang besar karena memiliki kompleksitas waktu sebesar O(n²). Padahal program menyediakan kapasitas array hingga 1.000.000 elemen, sehingga proses pengurutan dapat menjadi sangat lambat jika data yang diolah mendekati kapasitas maksimum. Selain itu, array <i>arr</i> dideklarasikan secara global sehingga seluruh test case menggunakan memori yang sama. Program juga tidak melakukan validasi apakah nilai <i>m</i> melebihi kapasitas array, sehingga terdapat risiko <i>out of bounds error</i> apabila jumlah data yang dimasukkan lebih besar dari <i>NMAX</i>. Untuk data berukuran besar, penggunaan algoritma yang lebih efisien seperti <i>Merge Sort</i> atau <i>Quick Sort</i> akan memberikan performa yang jauh lebih baik.</p>


### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

const NMAX = 1000000

var arrGanjil [NMAX]int
var arrGenap [NMAX]int

func sortGanjil(n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arrGanjil[j] < arrGanjil[idxMin] {
				idxMin = j
			}
		}
		temp := arrGanjil[i]
		arrGanjil[i] = arrGanjil[idxMin]
		arrGanjil[idxMin] = temp
	}
}

func sortGenap(n int) {
	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if arrGenap[j] > arrGenap[idxMax] {
				idxMax = j
			}
		}
		temp := arrGenap[i]
		arrGenap[i] = arrGenap[idxMax]
		arrGenap[idxMax] = temp
	}
}

func main() {
	var n, m, num int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)

		var nGanjil int = 0
		var nGenap int = 0

		for j := 0; j < m; j++ {
			fmt.Scan(&num)
			if num%2 != 0 {
				arrGanjil[nGanjil] = num
				nGanjil++
			} else {
				arrGenap[nGenap] = num
				nGenap++
			}
		}

		sortGanjil(nGanjil)
		sortGenap(nGenap)

		var pertama int = 1

		for j := 0; j < nGanjil; j++ {
			if pertama == 1 {
				fmt.Print(arrGanjil[j])
				pertama = 0
			} else {
				fmt.Print(" ", arrGanjil[j])
			}
		}

		for j := 0; j < nGenap; j++ {
			if pertama == 1 {
				fmt.Print(arrGenap[j])
				pertama = 0
			} else {
				fmt.Print(" ", arrGenap[j])
			}
		}
		fmt.Println()
	}
}	
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul14/output/soal2.png)
[penjelasan]
<p align="justify">Program ini bertujuan untuk mengelompokkan bilangan berdasarkan paritasnya (ganjil dan genap), kemudian mengurutkan bilangan ganjil secara menaik (<i>ascending</i>) dan bilangan genap secara menurun (<i>descending</i>). Setelah proses pengelompokan dan pengurutan selesai, seluruh bilangan ganjil ditampilkan terlebih dahulu, diikuti oleh seluruh bilangan genap.</p>

<ol>

<li>Program mendeklarasikan konstanta <i>NMAX</i> bernilai 1000000 sebagai kapasitas maksimum array. Selain itu, program mendeklarasikan dua array global, yaitu <i>arrGanjil</i> untuk menyimpan bilangan ganjil dan <i>arrGenap</i> untuk menyimpan bilangan genap.</li>

<li>Program mendefinisikan fungsi <i>sortGanjil(n int)</i> yang digunakan untuk mengurutkan isi array <i>arrGanjil</i> menggunakan algoritma <i>Selection Sort</i> secara menaik (<i>ascending</i>).</li>

<li>Pada fungsi <i>sortGanjil</i>, setiap iterasi mencari elemen terkecil pada bagian array yang belum terurut. Posisi elemen terkecil tersebut disimpan dalam variabel <i>idxMin</i>.</li>

<li>Setelah elemen terkecil ditemukan, program menukar (<i>swap</i>) elemen tersebut dengan elemen pada posisi saat ini sehingga bagian array yang telah diproses selalu berada dalam keadaan terurut.</li>

<li>Program juga mendefinisikan fungsi <i>sortGenap(n int)</i> yang digunakan untuk mengurutkan isi array <i>arrGenap</i> menggunakan algoritma <i>Selection Sort</i> secara menurun (<i>descending</i>).</li>

<li>Pada fungsi <i>sortGenap</i>, setiap iterasi mencari elemen terbesar pada bagian array yang belum terurut. Posisi elemen terbesar tersebut disimpan dalam variabel <i>idxMax</i>.</li>

<li>Setelah elemen terbesar ditemukan, program melakukan pertukaran data sehingga elemen terbesar ditempatkan pada posisi yang sesuai. Dengan cara ini, array genap akan tersusun dari nilai terbesar ke nilai terkecil.</li>

<li>Pada fungsi <i>main</i>, program mendeklarasikan variabel <i>n</i> untuk menyimpan jumlah test case, <i>m</i> untuk menyimpan jumlah data pada setiap test case, dan <i>num</i> untuk menyimpan bilangan yang sedang dibaca.</li>

<li>Program membaca nilai <i>n</i> yang menunjukkan banyaknya kumpulan data yang akan diproses.</li>

<li>Untuk setiap test case, program membaca nilai <i>m</i> sebagai jumlah bilangan yang akan dimasukkan. Selain itu, program menginisialisasi variabel <i>nGanjil</i> dan <i>nGenap</i> dengan nilai 0 untuk menghitung jumlah data ganjil dan genap.</li>

<li>Program kemudian membaca <i>m</i> buah bilangan satu per satu. Jika bilangan memiliki sisa bagi tidak sama dengan nol saat dibagi dua (<i>num % 2 != 0</i>), bilangan tersebut dimasukkan ke array <i>arrGanjil</i>. Sebaliknya, jika bilangan habis dibagi dua, bilangan dimasukkan ke array <i>arrGenap</i>.</li>

<li>Setiap kali sebuah bilangan dimasukkan ke array ganjil atau genap, penghitung jumlah elemen yang sesuai (<i>nGanjil</i> atau <i>nGenap</i>) akan ditambah satu.</li>

<li>Setelah seluruh data selesai dibaca, program memanggil fungsi <i>sortGanjil(nGanjil)</i> untuk mengurutkan seluruh bilangan ganjil secara menaik dan fungsi <i>sortGenap(nGenap)</i> untuk mengurutkan seluruh bilangan genap secara menurun.</li>

<li>Program menggunakan variabel <i>pertama</i> sebagai penanda agar elemen pertama dapat dicetak tanpa spasi di depan, sedangkan elemen berikutnya dicetak dengan awalan spasi.</li>

<li>Seluruh bilangan ganjil yang telah terurut dicetak terlebih dahulu menggunakan perulangan dari indeks 0 hingga <i>nGanjil - 1</i>.</li>

<li>Setelah itu, seluruh bilangan genap yang telah terurut dicetak menggunakan perulangan dari indeks 0 hingga <i>nGenap - 1</i>, sehingga hasil akhir menampilkan bilangan ganjil terurut naik diikuti bilangan genap terurut turun.</li>

<li>Setelah satu test case selesai ditampilkan, program mencetak baris baru dan melanjutkan proses ke test case berikutnya hingga seluruh data selesai diproses.</li>

</ol>

<p align="justify">Secara keseluruhan, program ini menerapkan konsep array, percabangan, perulangan bersarang (<i>nested loop</i>), pengelompokan data berdasarkan paritas, serta algoritma <i>Selection Sort</i>. Data dipisahkan menjadi kelompok ganjil dan genap, kemudian masing-masing kelompok diurutkan dengan aturan yang berbeda sebelum digabungkan kembali pada saat ditampilkan.</p>

<p align="justify">Kritik untuk kode ini adalah penggunaan algoritma <i>Selection Sort</i> pada kedua fungsi pengurutan memiliki kompleksitas waktu O(n²), sehingga kurang efisien untuk jumlah data yang besar. Hal ini menjadi kurang sesuai karena kapasitas array yang disediakan mencapai 1.000.000 elemen. Selain itu, array <i>arrGanjil</i> dan <i>arrGenap</i> bersifat global sehingga data dari test case sebelumnya tetap berada di memori, meskipun tidak memengaruhi hasil karena jumlah elemen dikontrol oleh variabel <i>nGanjil</i> dan <i>nGenap</i>. Program juga tidak melakukan validasi terhadap nilai <i>m</i>, sehingga jika jumlah data yang dimasukkan melebihi kapasitas array, program berpotensi mengalami <i>out of bounds error</i>. Untuk menangani data berukuran besar, algoritma yang lebih efisien seperti <i>Merge Sort</i> atau <i>Quick Sort</i> dapat digunakan untuk meningkatkan performa pengurutan.</p>

### 3. [Soal]
#### soal3.go

```go
package main

import "fmt"

const NMAX = 1000000

var arr [NMAX]int

func insertionSort(n int) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var n int = 0
	var num int

	fmt.Scan(&num)
	for num != -5313 {
		if num == 0 {
			if n > 0 {
				insertionSort(n)
				if n%2 != 0 {
					fmt.Println(arr[n/2])
				} else {
					fmt.Println((arr[n/2-1] + arr[n/2]) / 2)
				}
			}
		} else {
			arr[n] = num
			n++
		}
		fmt.Scan(&num)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul14/output/soal3.png)
[penjelasan]
<p align="justify">Program ini bertujuan untuk menghitung nilai median dari sekumpulan bilangan bulat. Data dimasukkan secara berurutan dan disimpan ke dalam array. Ketika pengguna memasukkan angka <i>0</i>, program akan mengurutkan seluruh data yang telah tersimpan menggunakan algoritma <i>Insertion Sort</i> dan menampilkan nilai median. Proses akan terus berlanjut hingga pengguna memasukkan nilai <i>-5313</i> sebagai tanda akhir program.</p>

<ol>

<li>Program mendeklarasikan konstanta <i>NMAX</i> bernilai 1000000 yang digunakan sebagai kapasitas maksimum array. Selain itu, dideklarasikan array global <i>arr</i> untuk menyimpan seluruh data bilangan yang dimasukkan pengguna.</li>

<li>Program mendefinisikan fungsi <i>insertionSort(n int)</i> yang digunakan untuk mengurutkan elemen array sebanyak <i>n</i> data menggunakan algoritma <i>Insertion Sort</i>.</li>

<li>Pada fungsi <i>insertionSort</i>, perulangan dimulai dari indeks 1 karena elemen pertama dianggap telah berada pada posisi yang benar.</li>

<li>Pada setiap iterasi, nilai <i>arr[i]</i> disimpan ke dalam variabel <i>key</i>. Nilai ini akan disisipkan ke posisi yang sesuai pada bagian array yang sudah terurut.</li>

<li>Variabel <i>j</i> diinisialisasi dengan nilai <i>i - 1</i> untuk membandingkan <i>key</i> dengan elemen-elemen sebelumnya.</li>

<li>Selama nilai <i>j</i> masih valid dan elemen <i>arr[j]</i> lebih besar daripada <i>key</i>, elemen tersebut digeser satu posisi ke kanan. Proses ini berlanjut hingga ditemukan posisi yang tepat untuk menyisipkan <i>key</i>.</li>

<li>Setelah posisi yang sesuai ditemukan, nilai <i>key</i> ditempatkan pada indeks <i>j + 1</i>. Dengan cara ini, bagian array yang telah diproses selalu berada dalam keadaan terurut.</li>

<li>Pada fungsi <i>main</i>, program mendeklarasikan variabel <i>n</i> untuk menghitung jumlah data yang tersimpan dan variabel <i>num</i> untuk menyimpan nilai input yang sedang dibaca.</li>

<li>Program membaca sebuah bilangan dari masukan dan menyimpannya ke dalam variabel <i>num</i>.</li>

<li>Selama nilai yang dimasukkan bukan <i>-5313</i>, program akan terus melakukan proses pembacaan data menggunakan perulangan <i>for</i>.</li>

<li>Jika nilai yang dibaca bukan <i>0</i>, maka nilai tersebut disimpan ke dalam array <i>arr</i> pada indeks ke-<i>n</i>, kemudian variabel <i>n</i> ditambah satu untuk mencatat jumlah data yang telah masuk.</li>

<li>Jika nilai yang dibaca adalah <i>0</i>, program menganggap bahwa pengguna meminta perhitungan median terhadap seluruh data yang telah tersimpan.</li>

<li>Sebelum menghitung median, program memeriksa apakah jumlah data lebih dari nol. Jika tidak ada data yang tersimpan, proses perhitungan median tidak dilakukan.</li>

<li>Apabila terdapat data, fungsi <i>insertionSort(n)</i> dipanggil untuk mengurutkan seluruh elemen array dari nilai terkecil ke terbesar.</li>

<li>Setelah data terurut, program menentukan median berdasarkan jumlah elemen yang ada. Jika jumlah data ganjil, median berada tepat di tengah array, yaitu pada indeks <i>n/2</i>.</li>

<li>Jika jumlah data genap, median dihitung dengan mengambil rata-rata dua elemen tengah, yaitu <i>arr[n/2 - 1]</i> dan <i>arr[n/2]</i>, kemudian hasilnya ditampilkan.</li>

<li>Setelah median dicetak, program kembali membaca input berikutnya. Data yang sudah tersimpan tetap berada di dalam array sehingga perhitungan median berikutnya akan menggunakan seluruh data yang telah dimasukkan sebelumnya.</li>

<li>Proses berakhir ketika pengguna memasukkan nilai <i>-5313</i> sebagai penanda akhir input.</li>

</ol>

<p align="justify">Secara keseluruhan, program ini menerapkan konsep array, perulangan, percabangan, algoritma <i>Insertion Sort</i>, dan perhitungan median. Data disimpan secara kumulatif, kemudian diurutkan setiap kali pengguna memasukkan angka <i>0</i> untuk memperoleh nilai tengah dari seluruh data yang telah masuk.</p>

<p align="justify">Kritik untuk kode ini adalah setiap kali pengguna memasukkan angka <i>0</i>, seluruh data diurutkan kembali menggunakan <i>Insertion Sort</i> meskipun sebagian besar data mungkin sudah terurut dari proses sebelumnya. Akibatnya, program menjadi kurang efisien karena memiliki kompleksitas waktu O(n²) untuk setiap proses pengurutan. Selain itu, median untuk jumlah data genap dihitung menggunakan pembagian bilangan bulat sehingga nilai pecahan akan dibulatkan ke bawah. Sebagai contoh, median dari data 1, 2, 3, dan 4 akan menghasilkan 2, bukan 2,5. Program juga tidak melakukan validasi apakah jumlah data telah melebihi kapasitas array <i>NMAX</i>, sehingga terdapat risiko <i>out of bounds error</i> apabila data yang dimasukkan terlalu banyak.</p>

### 4. [Soal]
#### soal4.go

```go
package main

import "fmt"

const NMAX = 1000

type tabInt [NMAX]int

func insertionSort(arr *tabInt, n int) {
	for i := 1; i < n; i++ {
		key := (*arr)[i]
		j := i - 1
		for j >= 0 && (*arr)[j] > key {
			(*arr)[j+1] = (*arr)[j]
			j--
		}
		(*arr)[j+1] = key
	}
}

func cekJarak(arr tabInt, n int) {
	if n > 1 {
		diff := arr[1] - arr[0]
		status := 1

		for i := 1; i < n-1; i++ {
			if arr[i+1]-arr[i] != diff {
				status = 0
			}
		}

		if status == 1 {
			fmt.Printf("Data berjarak %d\n", diff)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}

func main() {
	var arr tabInt
	var n int
	var num int

	fmt.Scan(&num)
	for num >= 0 {
		arr[n] = num
		n++
		fmt.Scan(&num)
	}

	insertionSort(&arr, n)

	for i := 0; i < n; i++ {
		fmt.Print(arr[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	cekJarak(arr, n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul14/output/soal4.png)
[penjelasan]
<p align="justify">Program ini bertujuan untuk mengurutkan sekumpulan bilangan bulat menggunakan algoritma <i>Insertion Sort</i>, kemudian memeriksa apakah seluruh data yang telah diurutkan memiliki selisih (jarak) yang tetap antar elemen berurutan. Jika selisihnya tetap, program akan menampilkan besar jaraknya. Jika tidak, program akan memberikan informasi bahwa data tidak memiliki jarak yang tetap.</p>

<ol>

<li>Program mendeklarasikan konstanta <i>NMAX</i> bernilai 1000 yang digunakan sebagai kapasitas maksimum array.</li>

<li>Program mendefinisikan tipe data <i>tabInt</i> sebagai array bertipe integer dengan ukuran maksimum <i>NMAX</i>. Tipe ini digunakan untuk mempermudah pengelolaan data pada program.</li>

<li>Program mendeklarasikan fungsi <i>insertionSort(arr *tabInt, n int)</i> yang digunakan untuk mengurutkan data secara menaik (<i>ascending</i>) menggunakan algoritma <i>Insertion Sort</i>.</li>

<li>Pada fungsi <i>insertionSort</i>, perulangan dimulai dari indeks 1 karena elemen pertama dianggap telah berada pada posisi yang benar.</li>

<li>Setiap elemen yang sedang diproses disimpan ke dalam variabel <i>key</i>. Variabel ini digunakan sebagai nilai yang akan disisipkan ke posisi yang sesuai pada bagian array yang telah terurut.</li>

<li>Variabel <i>j</i> digunakan untuk menelusuri elemen-elemen sebelumnya. Selama elemen sebelumnya lebih besar daripada <i>key</i>, elemen tersebut digeser satu posisi ke kanan.</li>

<li>Setelah posisi yang sesuai ditemukan, nilai <i>key</i> ditempatkan pada indeks <i>j+1</i>. Dengan cara ini, bagian array yang telah diproses selalu berada dalam keadaan terurut.</li>

<li>Program juga mendefinisikan fungsi <i>cekJarak(arr tabInt, n int)</i> yang digunakan untuk memeriksa apakah selisih antar elemen yang berurutan memiliki nilai yang sama.</li>

<li>Pada fungsi <i>cekJarak</i>, program terlebih dahulu memastikan bahwa jumlah data lebih dari satu elemen karena minimal diperlukan dua data untuk menghitung selisih.</li>

<li>Program menghitung selisih awal antara dua elemen pertama, yaitu <i>arr[1] - arr[0]</i>, dan menyimpannya ke dalam variabel <i>diff</i>. Nilai ini dijadikan acuan untuk pemeriksaan berikutnya.</li>

<li>Variabel <i>status</i> diinisialisasi dengan nilai 1 sebagai penanda bahwa data sementara dianggap memiliki jarak yang tetap.</li>

<li>Program melakukan perulangan untuk membandingkan selisih setiap pasangan elemen yang berurutan dengan nilai <i>diff</i>. Jika ditemukan selisih yang berbeda, maka nilai <i>status</i> diubah menjadi 0.</li>

<li>Setelah seluruh data diperiksa, program mengevaluasi nilai <i>status</i>. Jika nilainya tetap 1, program menampilkan pesan bahwa data memiliki jarak tetap beserta besar jaraknya.</li>

<li>Jika nilai <i>status</i> bernilai 0, program menampilkan pesan bahwa data memiliki jarak yang tidak tetap.</li>

<li>Pada fungsi <i>main</i>, program mendeklarasikan variabel <i>arr</i> bertipe <i>tabInt</i> untuk menyimpan data, variabel <i>n</i> untuk menghitung jumlah data yang masuk, dan variabel <i>num</i> untuk menyimpan nilai input sementara.</li>

<li>Program membaca data satu per satu dari masukan. Selama nilai yang dimasukkan bernilai nol atau positif (<i>num >= 0</i>), data akan disimpan ke dalam array dan jumlah data <i>n</i> akan bertambah.</li>

<li>Proses input berhenti ketika pengguna memasukkan bilangan negatif. Bilangan negatif tersebut tidak disimpan ke dalam array.</li>

<li>Setelah seluruh data selesai dibaca, program memanggil fungsi <i>insertionSort</i> untuk mengurutkan data dari nilai terkecil ke terbesar.</li>

<li>Program kemudian menampilkan seluruh data yang telah terurut dalam satu baris dengan setiap elemen dipisahkan oleh spasi.</li>

<li>Terakhir, fungsi <i>cekJarak</i> dipanggil untuk memeriksa apakah data yang telah diurutkan memiliki selisih tetap antar elemen dan menampilkan hasil pemeriksaannya.</li>

</ol>

<p align="justify">Secara keseluruhan, program ini menerapkan konsep array, tipe data bentukan (<i>user-defined type</i>), perulangan, percabangan, algoritma <i>Insertion Sort</i>, dan pemeriksaan pola selisih antar elemen. Setelah data diurutkan, program menentukan apakah data tersebut membentuk suatu barisan dengan beda yang tetap, mirip dengan konsep barisan aritmetika.</p>

<p align="justify">Kritik untuk kode ini adalah fungsi <i>cekJarak</i> tetap melanjutkan pemeriksaan hingga akhir meskipun sudah ditemukan selisih yang berbeda. Proses tersebut dapat dibuat lebih efisien dengan menghentikan perulangan segera setelah ketidaksesuaian ditemukan. Selain itu, program tidak memberikan keluaran khusus apabila jumlah data kurang dari dua elemen, sehingga pengguna tidak memperoleh informasi mengenai kondisi tersebut. Program juga tidak melakukan validasi apakah jumlah data yang dimasukkan telah melebihi kapasitas array <i>NMAX</i>, sehingga terdapat risiko <i>out of bounds error</i> jika data yang dimasukkan terlalu banyak. Dari sisi performa, penggunaan <i>Insertion Sort</i> dengan kompleksitas waktu O(n²) juga kurang optimal untuk data berukuran besar dibandingkan algoritma pengurutan yang lebih efisien seperti <i>Merge Sort</i> atau <i>Quick Sort</i>.</p>

### 5. [Soal]
#### soal5.go

```go
package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax + 1]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i <= n; i++ {
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n > 0 {
		maxRating := pustaka[1].rating
		for i := 2; i <= n; i++ {
			if pustaka[i].rating > maxRating {
				maxRating = pustaka[i].rating
			}
		}

		for i := 1; i <= n; i++ {
			if pustaka[i].rating == maxRating {
				fmt.Printf("%s, %s, %s, %d\n", pustaka[i].judul, pustaka[i].penulis, pustaka[i].penerbit, pustaka[i].tahun)
			}
		}
	}
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 2; i <= n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 1 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	limit := 5
	if n < 5 {
		limit = n
	}
	for i := 1; i <= limit; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	kiri := 1
	kanan := n
	found := -1

	for kiri <= kanan && found == -1 {
		med := (kiri + kanan) / 2
		if pustaka[med].rating == r {
			found = med
		} else if pustaka[med].rating < r {
			kanan = med - 1
		} else {
			kiri = med + 1
		}
	}

	if found != -1 {
		fmt.Printf("%s, %s, %s, %d, %d, %d\n", pustaka[found].judul, pustaka[found].penulis, pustaka[found].penerbit, pustaka[found].tahun, pustaka[found].eksemplar, pustaka[found].rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var n, targetRating int
	var pustaka DaftarBuku

	fmt.Scan(&n)
	DaftarkanBuku(&pustaka, n)
	fmt.Scan(&targetRating)

	CetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)
	CariBuku(pustaka, n, targetRating)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul14/output/soal5.png)
[penjelasan]
<p align="justify">Program ini bertujuan untuk mengelola data buku pada sebuah perpustakaan sederhana. Program menerima sejumlah data buku, kemudian menampilkan buku dengan rating tertinggi, mengurutkan buku berdasarkan rating secara menurun, menampilkan lima buku teratas setelah pengurutan, serta mencari buku berdasarkan rating tertentu menggunakan algoritma <i>Binary Search</i>.</p>

<ol>

<li>Program mendeklarasikan konstanta <i>nMax</i> bernilai 7919 yang digunakan sebagai kapasitas maksimum penyimpanan data buku.</li>

<li>Program mendefinisikan tipe data <i>Buku</i> yang berisi beberapa atribut, yaitu <i>id</i>, <i>judul</i>, <i>penulis</i>, <i>penerbit</i>, <i>eksemplar</i>, <i>tahun</i>, dan <i>rating</i>. Struktur ini digunakan untuk menyimpan informasi lengkap setiap buku.</li>

<li>Program juga mendefinisikan tipe data <i>DaftarBuku</i> berupa array yang berisi kumpulan data bertipe <i>Buku</i>.</li>

<li>Fungsi <i>DaftarkanBuku(pustaka *DaftarBuku, n int)</i> digunakan untuk membaca data buku dari masukan dan menyimpannya ke dalam array <i>pustaka</i>.</li>

<li>Pada fungsi tersebut, program melakukan perulangan dari indeks 1 sampai <i>n</i>. Pada setiap iterasi, program membaca seluruh atribut buku dan menyimpannya ke elemen array yang sesuai.</li>

<li>Fungsi <i>CetakTerfavorit(pustaka DaftarBuku, n int)</i> digunakan untuk mencari dan menampilkan buku yang memiliki rating tertinggi.</li>

<li>Pada fungsi ini, program terlebih dahulu menginisialisasi variabel <i>maxRating</i> dengan rating buku pertama. Selanjutnya dilakukan penelusuran seluruh data buku untuk mencari nilai rating terbesar.</li>

<li>Setelah rating tertinggi ditemukan, program melakukan perulangan kembali untuk menampilkan semua buku yang memiliki rating sama dengan <i>maxRating</i>. Dengan demikian, jika terdapat lebih dari satu buku dengan rating tertinggi, seluruhnya akan ditampilkan.</li>

<li>Fungsi <i>UrutBuku(pustaka *DaftarBuku, n int)</i> digunakan untuk mengurutkan data buku berdasarkan rating secara menurun (<i>descending</i>) menggunakan algoritma <i>Insertion Sort</i>.</li>

<li>Pada setiap iterasi pengurutan, sebuah data buku disimpan ke variabel <i>key</i>. Program kemudian menggeser seluruh buku yang memiliki rating lebih kecil ke posisi berikutnya hingga ditemukan posisi yang sesuai untuk <i>key</i>.</li>

<li>Setelah proses penggeseran selesai, data buku pada variabel <i>key</i> ditempatkan pada posisi yang tepat sehingga array tetap terurut berdasarkan rating tertinggi ke terendah.</li>

<li>Fungsi <i>Cetak5Terbaru(pustaka DaftarBuku, n int)</i> digunakan untuk menampilkan lima buku pertama setelah proses pengurutan selesai.</li>

<li>Variabel <i>limit</i> diinisialisasi dengan nilai 5. Jika jumlah buku kurang dari lima, maka nilai <i>limit</i> disesuaikan dengan jumlah buku yang tersedia.</li>

<li>Program kemudian mencetak judul buku dari indeks 1 hingga <i>limit</i>. Karena data telah diurutkan berdasarkan rating tertinggi, buku-buku yang ditampilkan merupakan lima buku dengan rating tertinggi.</li>

<li>Fungsi <i>CariBuku(pustaka DaftarBuku, n int, r int)</i> digunakan untuk mencari buku berdasarkan rating menggunakan algoritma <i>Binary Search</i>.</li>

<li>Pada fungsi ini, variabel <i>kiri</i> dan <i>kanan</i> digunakan untuk menandai batas pencarian. Variabel <i>found</i> digunakan sebagai penanda apakah data ditemukan atau tidak.</li>

<li>Program menghitung indeks tengah (<i>med</i>) dan membandingkan rating buku pada posisi tersebut dengan rating yang dicari.</li>

<li>Jika rating sesuai dengan target, indeks buku disimpan ke dalam variabel <i>found</i>. Jika rating buku lebih kecil dari target, pencarian dilanjutkan ke bagian kiri array karena data telah diurutkan secara menurun. Sebaliknya, jika rating buku lebih besar dari target, pencarian dilanjutkan ke bagian kanan array.</li>

<li>Apabila data ditemukan, program menampilkan informasi lengkap buku yang meliputi judul, penulis, penerbit, tahun terbit, jumlah eksemplar, dan rating.</li>

<li>Jika tidak ditemukan buku dengan rating yang dicari, program menampilkan pesan "Tidak ada buku dengan rating seperti itu".</li>

<li>Pada fungsi <i>main</i>, program mendeklarasikan variabel <i>n</i> untuk jumlah buku, <i>targetRating</i> untuk rating yang akan dicari, dan variabel <i>pustaka</i> untuk menyimpan seluruh data buku.</li>

<li>Program membaca jumlah buku, memanggil fungsi <i>DaftarkanBuku</i> untuk mengisi data perpustakaan, kemudian membaca nilai rating yang akan dicari.</li>

<li>Selanjutnya program memanggil fungsi <i>CetakTerfavorit</i>, <i>UrutBuku</i>, <i>Cetak5Terbaru</i>, dan <i>CariBuku</i> secara berurutan untuk menghasilkan seluruh keluaran yang diperlukan.</li>

</ol>

<p align="justify">Secara keseluruhan, program ini menerapkan konsep struktur data (<i>struct</i>), array, prosedur, perulangan, algoritma <i>Insertion Sort</i>, dan algoritma <i>Binary Search</i>. Data buku diolah untuk memperoleh informasi buku dengan rating tertinggi, mengurutkan koleksi berdasarkan rating, menampilkan lima buku terbaik, serta mencari buku berdasarkan nilai rating tertentu.</p>

<p align="justify">Kritik untuk kode ini adalah nama fungsi <i>Cetak5Terbaru</i> kurang sesuai dengan implementasinya. Fungsi tersebut sebenarnya menampilkan lima buku dengan rating tertinggi setelah proses pengurutan, bukan lima buku terbaru berdasarkan tahun terbit. Selain itu, fungsi <i>CariBuku</i> hanya menampilkan satu buku meskipun mungkin terdapat beberapa buku dengan rating yang sama. Penggunaan <i>Insertion Sort</i> juga memiliki kompleksitas waktu O(n²), sehingga kurang efisien jika jumlah buku mendekati kapasitas maksimum. Program juga tidak melakukan validasi terhadap nilai <i>n</i> sehingga terdapat risiko <i>out of bounds error</i> apabila jumlah buku yang dimasukkan melebihi kapasitas array. Dari sisi pencarian, penggunaan <i>Binary Search</i> sudah tepat karena data telah diurutkan terlebih dahulu, sehingga proses pencarian memiliki kompleksitas waktu O(log n).</p>