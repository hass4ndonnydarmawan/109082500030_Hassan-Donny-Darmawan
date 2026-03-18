package main
import "fmt"

func factorial(n int)int {
	hasil := 1
	for i:= 1; i<=n;{
		hasil*=i
		i++
	}
	return hasil
}


func permutation (n,r int)int {
	if n>=r {
		var hasil int
		hasil = factorial(n)/factorial(n-r)
		return hasil
	}
	return 0
}


func combination (n,r int)int {
	if n>=r {
		var hasil int
		hasil = factorial(n)/(factorial(r) * factorial(n-r))
		return hasil
	}
	return 0
}

func main () {
	var a,b,c,d int
	fmt.Scan(&a,&b,&c,&d)
	fmt.Printf("%d %d\n",permutation(a,c),combination(a,c))
	fmt.Printf("%d %d\n",permutation(b,d),combination(b,d))

}