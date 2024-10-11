package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("점수 입력 :")
	in := bufio.NewReader(os.Stdin)
	score, err := in.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	score = strings.TrimSpace(score)                // TrimSpace() 공백, 띄어쓰기, 줄바꿈 제거(python의 strip과 유사)
	realScore, _ := strconv.ParseInt(score, 10, 32) // 정수형 32비트 타입으로 형변환(string, 진법, 비트), := 선언 단축
	if realScore >= 60 {
		fmt.Println("A")
	} else {
		fmt.Println("BCDF")
	}
}
