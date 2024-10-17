package main

import (
	"fmt"
	"reflect"
)

func main() {

	// i := 55
	// fmt.Printf("i is %s", i) 출력 오류 발생

	c1 := 'Z'
	c2 := '김'

	fmt.Println(c1, c2)                                 // 10진수 출력
	fmt.Printf("%x\n", c2)                              // 유니코드 '김'에 대한 16진수 출력(http://home.unicode.org)
	fmt.Println(reflect.TypeOf(c1), reflect.TypeOf(c2)) //int32임!!!!!! string아님!!!
}
