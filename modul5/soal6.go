package main 
import "fmt"

func pangkat (n, a  int )int{
	if a == 0{
		return 1
	}
	return n * pangkat(n, a-1)
}

func main (){
	var n,a int
	fmt.Scan(&n,&a)
	fmt.Print(pangkat(n,a))
}