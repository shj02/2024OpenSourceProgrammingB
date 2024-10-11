package main

import (
	"fmt"
)

func main() {

	var f float64
	var i int
	var b bool
	var s string

	fmt.Println(f, b, s, i) // 0 value(어떤 값을 할당하지 않았을 때의 값) => 0 false 공백 0
	fmt.Printf("%f%t%s%d\n", f, b, s, i)

	i = 3
	f = 2.7
	fmt.Print("\n\n", f > float64(i), "\n") // 2.7 > 3.0 -> false
	fmt.Print(f <= float64(i))              // 2.7 <= 3.0 -> true
}
