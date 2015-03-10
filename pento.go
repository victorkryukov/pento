package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var Pentos []Figure // All different pento figures

func init() {
	p := Figure{Point{0, 0}, Point{0, 1}, Point{1, 1}, Point{0, 2}, Point{0, 3}}
	Pentos = make([]Figure, 8)
	Pentos[0] = p
	Pentos[1] = p.Rotate()
	Pentos[2] = p.Rotate().Rotate()
	Pentos[3] = p.Rotate().Rotate().Rotate()
	for i := 0; i <= 3; i++ {
		Pentos[i+4] = Pentos[i].Mirror()
	}
	for i := range Pentos {
		Pentos[i] = Pentos[i].Recenter()
	}
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Cannot convert '%s' to int\n", s)
	}
	return i
}

func main() {
	b := NewBoard(atoi(os.Args[1]), atoi(os.Args[2]))
	printBoard := func(b *Board) {
		fmt.Println(b)
	}
	b.Fill(Pentos, printBoard)
}
