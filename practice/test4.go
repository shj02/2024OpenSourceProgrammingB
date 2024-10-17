package main

import (
	"fmt"
)

func main() {
	var f float64
	var i int
	var b bool
	var s string
	// 지정 안하면 각 0 0 False (공백) 으로 나옴

	fmt.Println(f, i, b, s)
	fmt.Printf("%f %d %t %s\n", f, i, b, s)
}
