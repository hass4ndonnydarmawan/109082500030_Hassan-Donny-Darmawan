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