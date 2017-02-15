package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	z := float64(2)
	zsebelum := float64(1)
	for math.Abs(z-zsebelum) > 0.001 {
		zbaru:=zsebelum-(((zsebelum*zsebelum)-x)/(2*zsebelum))
		zsebelum=z
		z=zbaru
		fmt.Println(z)
	}
	return z
	}

func main() {
	fmt.Println(Sqrt(100))
}
