# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">Hassan Donny Darmawan - 109082500030</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main
import "fmt"
func main() {
var (
satu, dua, tiga string
temp string
)
fmt.Print("Masukan input string: ")
fmt.Scanln(&satu)
fmt.Print("Masukan input string: ")
fmt.Scanln(&dua)
fmt.Print("Masukan input string: ")
fmt.Scanln(&tiga)
fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
temp = satu
satu = dua
dua = tiga
tiga = temp
fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul1/output/output_soal1.png)
[penjelasan]
<p align="justify">program ini bertujuan untuk melakukan pertukaran nilai variable. proses pertukaran ini menggunakan variable tambahan yaitu variable temp yang berfungsi sebagai tempat penyimpan nilai sementara agar saat nilai awal sedang melakukan pertukaran, nilai tersebut tidak hilang melainkan disimpan pada varible temp. lalu saat sudah selesai prosesnya maka user akan mendapatkan nilai variable yang memiliki urutan baru</p> 


### 2. [Soal]
#### soal2.go

```go
package main
import "fmt"
func main() {

}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hass4ndonnydarmawan/109082500030_Hassan-Donny-Darmawan/blob/main/modul1/output_soal1.png)
[penjelasan]s

