package main
import "fmt"

func baris (n int){
	if n == 1{
		fmt.Print(1, " ")
		return
	}
	fmt.Print(n," ")
	baris(n-1)
	fmt.Print(n," ")
}

func main (){
	var n int
	fmt.Scan(&n)
	baris(n)
}