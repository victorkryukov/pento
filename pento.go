package main

import "fmt"

const Num = 8          // Number of different pento figures, with all mirros and rotations
var Pentos [Num]Figure // All different pento figures

func init() {
	p := Figure{Point{0, 0}, Point{0, 1}, Point{1, 1}, Point{0, 2}, Point{0, 3}}
	Pentos[0] = p
	Pentos[1] = p.Rotate()
	Pentos[2] = p.Rotate().Rotate()
	Pentos[3] = p.Rotate().Rotate().Rotate()
	for i := 0; i <= 3; i++ {
		Pentos[i+4] = Pentos[i].Mirror()
	}
}

func main() {
	for _, f := range Pentos {
		fmt.Println(f)
	}
}
