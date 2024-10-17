package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("점수 입력 : ")
	in := bufio.NewReader(os.Stdin)
	score, err := in.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	// TrimSpace() 공백, 띄어쓰기, 줄바꿈 제거(python의 strip과 유사)
	score = strings.TrimSpace(score)

	// 실수형 64비트 타입으로 형변환
	// realScore, _ := strconv.ParseFloat(score, 64)

	// 10진수 정수형 32비트 타입으로 형변환(string, 진법, 비트)
	realScore, _ := strconv.ParseInt(score, 10, 32)

	// 16진수 정수형 32비트 타입으로 형변환
	// realScore, _ := strconv.ParseInt(score, 16, 32)

	var grade string
	if realScore >= 90 {
		grade = "A"
	} else {
		grade = "BCDF"
	}

	fmt.Printf("%d점은 %s등급입니다.\n", realScore, grade)

}
