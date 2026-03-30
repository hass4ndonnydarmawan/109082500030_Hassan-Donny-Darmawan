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