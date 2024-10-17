package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	// f := 1.92 -> 64바이트
	var f float64 = 1.92
	i := 55

	fmt.Println(strings.Title("kim inha"))
	fmt.Printf("%f %f\n", f, math.Ceil(f))
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))
}
