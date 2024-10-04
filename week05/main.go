package main

import (
	"fmt" // c언어에서의 #include <stdio.h>와 같음
)

func main() {
	var f float64
	var i int
	var b bool
	var s string

	fmt.Println(f, b, s, i)
	fmt.Printf("%f%t%s%d\n", f, b, s, i)
}
