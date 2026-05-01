package main 
import (
	"fmt"
	"math"
)

type pusat struct {
	cx,cy,r float64
}
type input [2]pusat

func jarak (a,b,c,d float64)float64 {
	return math.Sqrt(math.Pow((a-c),2) + math.Pow((b-d),2))
}

func didalam (cx,cy,r,x,y float64) bool {
	return jarak(cx,cy,x,y) <= r
}

func main (){
	var (
		masukan input
		x,y float64
	)

	fmt.Scan(&masukan[0].cx, &masukan[0].cy, &masukan[0].r)
    fmt.Scan(&masukan[1].cx, &masukan[1].cy, &masukan[1].r)
    fmt.Scan(&x, &y)

	d1 := didalam(masukan[0].cx,masukan[0].cy,masukan[0].r,x,y)
	d2 := didalam(masukan[1].cx,masukan[1].cy,masukan[1].r,x,y)

	if d1 && d2 {
		fmt.Print("titik didalam lingkaran 1 dan 2")
	}else if d1 {
		fmt.Print("titik didalam lingkaran 1")
	}else if d2 {
		fmt.Print("titik didalam lingkaran 2")
	}else {
		fmt.Print("titik diluar lingkaran 1 dan 2")
	}
}