package main

import (
	"fmt"
	"math"
)

func fogoh (n int)int {
	h := n+1
	return int(float64(math.Pow(float64(h)-float64(2),2)))
}

func gohof (n int)int{
	f:= math.Pow(float64(n),float64(2))
	return int(f+1)-2
}

func hofog (n int)int{
	g := n-2
	return int((math.Pow(float64(g),float64(2)))+1)
}

func main (){
	var a,b,c int
	fmt.Scan(&a,&b,&c)
	fmt.Printf("%d\n", fogoh(a))
	fmt.Printf("%d\n", gohof(b))
	fmt.Printf("%d\n", hofog(c))
}
