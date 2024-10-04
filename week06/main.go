package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 13
	var f float64 = 12.9 //f := 12.9

	fmt.Printf("value i  : %d, value f : %f\n", i, f)
	// fmt.Printf("%d * %f = %f", i, f, i*f) mismatched types int and float64 : 타입이 안 맞음
	// fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f) 타입 변환
	fmt.Printf("%d * %f = %d\n", i, f, i*int(f))      // 타입 변환
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i)) // 원본 타입은 바뀌지 않음!
}
