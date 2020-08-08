package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool         = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main(){
	fmt.Print("Type: %T Value: %v\n" , ToBe , ToBe , "\n")
	fmt.Print("Type: %T Value: %v\n" , MaxInt ,MaxInt, "\n")
	fmt.Print("Type: %T Value: %v\n" , z , z , "\n")
}
