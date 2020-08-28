package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	const (
		A = iota
		B = iota
		C
		D = 6
		E = iota
	)
	fmt.Println(A, B, C, D, E)
	// time.Sleep(100)
	end := time.Now()
	fmt.Println(end.Sub(start))

}
