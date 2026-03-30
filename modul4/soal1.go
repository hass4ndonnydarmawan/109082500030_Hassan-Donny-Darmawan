package main
import "fmt"

func faktorial (n int, hasil *int){
	*hasil = 1
	for i:= 1; i<=n;  i++{
		*hasil *= i
	}
}
func permutation (n, r int, hasil *int){
	var fn, fnr int
	if n>=r {
		faktorial(n,&fn)
		faktorial(n-r,&fnr)

		*hasil = fn/fnr
	}else {
		*hasil = 0
	}
}

func combination (n,r int, hasil *int){
	var fr, fnr, fn int
	if n>=r{
	faktorial(n,&fn)
	faktorial(r,&fr)
	faktorial(n-r,&fnr)

	*hasil = fn/(fr*fnr)
	}else {
		*hasil = 0
	}
}

func main (){
	var a,b,c,d,p,co int
	fmt.Scan(&a,&b,&c,&d)
	permutation(a,c,&p)
	combination (a,c,&co)
	fmt.Println(p,co)

	permutation(b,d,&p)
	combination (b,d,&co)
	fmt.Println(p,co)
}