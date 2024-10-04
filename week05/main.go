package main

import (
	"fmt" // c언어에서의 #include <stdio.h>와 같음
	"math"
	"reflect"
	"strings"
)

func main() {
	// f := 1.92 이대로 출력하면 64바이트로 나옴
	var f float64 = 1.92 // float64=double(c언어)

	i := "55" //문자열로 변경
	// var i int
	// i = 55
	// var i int = 55

	fmt.Println(strings.Title("kim inha")) // Title 함수 -> 첫글자 대문자 변경
	fmt.Printf("%f %f\n", f, math.Ceil(f))
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))
	fmt.Println("f is", f)
	fmt.Println("i is", i)
	fmt.Print("i is ", i, "\n") // 띄어쓰기
	fmt.Println("i is", i)
	// fmt.Printf("i is %d\n", i)
	fmt.Printf("i is %s", i) // %s 로 문자열 타입
}
