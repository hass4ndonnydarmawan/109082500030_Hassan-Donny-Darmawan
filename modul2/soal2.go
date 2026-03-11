package main
import "fmt"

func main (){
var w1,w2,w3,w4 string
status := true

for i:=1;i<=5;{
	fmt.Print("Percobaan ",i,": ")
	fmt.Scanln(&w1,&w2,&w3,&w4)
	if w1 != "merah" || w2!= "kuning" || w3!="hijau" || w4!="ungu"{
		status= false
	}
	i++
}
fmt.Println("Berhasil:",status)
}