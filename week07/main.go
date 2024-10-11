package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Printf("이름 입력 :")
	in := bufio.NewReader(os.Stdin)
	name, err := in.ReadString('\n')

	if err != nil { // 에러가 있다면
		log.Fatal(err) // 프로그램 종료
	} else {
		fmt.Println(name)
	}
}
