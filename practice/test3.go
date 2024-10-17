package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 13
	var f float64 = 12.9

	fmt.Printf("value i : %d, value f : %f\n", i, f)
	fmt.Printf("%d * %f = %d\n", i, f, i*int(f))      //f의 타입을 int로 변환
	fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f)  //i의 타입을 float64로 변환
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f)) //원본 타입은 바뀌지 않음
}
