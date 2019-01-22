package main

import (
	"fmt"
)

type Cc float64
type Ff float64

const (
	Zero Cc = -273.15
	Free Cc = 0
	Boil Cc = 100
)

func CToF(c Cc) Ff {
	return Ff(c*9/5 + 32)
}

func FToC(f Ff) Cc {
	return Cc((f - 32) / 9 * 5)
}

func (c Cc) String() string {
	return fmt.Sprintf("%g°C", c)
}

func main() {
	fmt.Printf("%g\n", Boil-Free) // "100" °C
	boilF := CToF(Boil)
	fmt.Printf("%g\n", boilF-CToF(Free)) // "180" °F

	// 下面三行注释 compile error: type mismatch
	// var c1 Cc;
	// var f1 Ff;
	// fmt.Println(c1 == f1)

	c2 := FToC(212.0)
	fmt.Println(c2.String()) // "100°C"
	fmt.Printf("%v\n", c2)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c2)   // "100°C"
	fmt.Println(c2)          // "100°C"
	fmt.Printf("%g\n", c2)   // "100"; does not call String
	fmt.Println(float64(c2)) // "100"; does not call String
}
