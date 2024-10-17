package main

import "fmt"

func main() {
	i := 3
	f := 2.7

	fmt.Print(f > float64(i), "\n") //false 타입 맞춰줘야 함
	fmt.Print(f <= float64(i))      //true
}
