package main

import "fmt" // c언어에서의 #include <stdio.h>와 같음

func main() {
	//var i int
	//i = 55
	//var i int = 55
	var f float32 = 1.92 //float64=double(c언어)
	i := 55              //단축어

	fmt.Println("f is", f)
	fmt.Println("i is", i)
	fmt.Print("i is ", i, "\n") //띄어쓰기
	fmt.Println("i is", i)
	fmt.Printf("i is %d\n", i)
}
