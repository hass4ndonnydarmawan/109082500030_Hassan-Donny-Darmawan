package main 
import "fmt"

func ganjil (n int, a int){
	if a<=n{
		fmt.Print(a," ")
		ganjil(n, a+2)
	}
}

func main (){
	var n int
	fmt.Scan(&n)
	ganjil(n,1)
}