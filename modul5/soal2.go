package main
import "fmt"

func bintang (n int){
	if n!=0 {
		fmt.Print("*")
		bintang(n-1)
	}
}

func main (){
	var n int
	fmt.Scan(&n)

	for i:= 1;i<=n;i++{
		bintang(i)
		fmt.Println()
	}
}