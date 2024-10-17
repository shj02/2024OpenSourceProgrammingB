package main

import (
	"bufio" // 버퍼링된 I/O를 제공, 입력 스트림을 더 효율적으로 읽을 수 있게 해줌
	"fmt"
	"log"
	"os" // 운영 체제 관련 함수 제공
)

func main() {
	fmt.Printf("이름 입력 : ")

	// 표준입력(os.Stdin)을 읽기 위해 리더 생성
	in := bufio.NewReader(os.Stdin)

	// 사용자가 Enter를 누를 때까지 입력을 받아서 반환함, 만약 입력 도중 오류가 발생하면 err에 그 오류가 저장되고, 그렇지 않으면 nil임
	name, err := in.ReadString('\n')

	//Println 또는 Print만 가능 (Printf 불가능)
	fmt.Println(name)

	// 에러 값 출력
	fmt.Println(err)

	if err != nil { // 에러가 있다면
		log.Fatal(err) // 프로그램 종료
	} else {
		fmt.Println(name)
	}
}
