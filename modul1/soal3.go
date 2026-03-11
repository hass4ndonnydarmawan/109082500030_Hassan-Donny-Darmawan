package main 
import "fmt"

func main (){
var kg_parsel, sisa, parsel int
var biayautama, biayatambahan, total float64

fmt.Print("Berat parsel(gram): ")
fmt.Scanln(&parsel)
kg_parsel = parsel / 1000
sisa = parsel % 1000

fmt.Println("Detail berat:", kg_parsel, "kg +", sisa, "gr")
biayautama = float64(kg_parsel)*10000
if kg_parsel > 10{
	biayatambahan = 0
} else if sisa >=500 {
	biayatambahan = float64(sisa)*5
}else {
	biayatambahan = float64(sisa)*15
}

fmt.Println("Detail biaya: Rp.",biayautama, "+ Rp.",biayatambahan)
total = biayautama + biayatambahan
fmt.Println("Total biaya: Rp.",total)
}