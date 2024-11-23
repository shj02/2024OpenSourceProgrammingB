package main

import (
	"fmt"
	"os"
	"reflect"
)

// func test(strs string) {
// func test(strs ...string, i int) {   // error
func test(i int, strs ...string) { /// ... : 가변매개변수
	fmt.Println(i, strs, reflect.TypeOf(strs))
}

func main() {
	// fmt.Println(os.Args, len(os.Args), reflect.TypeOf(os.Args))
	slices := os.Args[1:] // 0번방: 실행경로
	fmt.Println(slices, slices[2])
	test(1, "123")
	test(-99, "123", "abc")
	test(55)
	test(0, "123", "abc", "123a")
}
