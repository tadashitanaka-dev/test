package main

import (
	"fmt"
	"math"
)



func Sqrt(x float64) float64 {
	z := 1.0
	//z := float64(1)
	for i := 1 ; i < 100 ;i++{
	z -= (z*z - x) / (2*z)
	fmt.Println(z,math.Sqrt(2),z - math.Sqrt(2))
	//fmt.Println(z - math.Sqrt(2))
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}


