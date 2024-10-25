package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {
	// fmt.Printf("%d\n", rand.Intn(6)+1)
	r := fmt.Sprintf("%d\n", rand.Intn(6)+1) // 출력 안함! %10.3f : 10칸 잡고 소수 3째자리까지 출력
	fmt.Println(reflect.TypeOf(r))
	fmt.Printf("%T\n", 2.1) // 타입 반환(reflect.TypeOf랑 같음)

	i := 1
	for i <= 10 {
		fmt.Printf("%2d점\n", i)
		i = i + 1
	}
}
