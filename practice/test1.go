package main

import "fmt"

func main() {
	//var i int
	//i = 55

	//var i int = 55

	var f float32 = 1.92
	i := 55

	fmt.Println("f is", f) //기본출력, 줄바꿈함
	fmt.Println("i is", i)
	fmt.Print("i is ", i, "\n") //기본출력, 줄바꿈 안함, 문자열로 변환 후 출력
	fmt.Printf("i is %d\n", i)  //형식화된 출력, 줄바꿈 안함

	fmt.Println("-----")

	fmt.Println("i is ", i, "~")
	fmt.Println("i is %d", i) //적용안됨!

	fmt.Print("1")
	fmt.Print("2") //12
	fmt.Print("줄바꿈\n")
	fmt.Print("%d", i, "\n") //적용안됨!

	fmt.Printf("i is ", i)
	fmt.Printf(" 적용안됨!\n")
	fmt.Printf("i is %d", i)

	// 정리
	// 줄바꿈 : Println
	// %d 적용 : Printf
	// "i is", i 적용 : Println
}
