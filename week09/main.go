package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // seed로 고정시킴(39를 하면 6만 나옴), 유닉스 시간
	answer := rand.Intn(6) + 1   // dice 1~6까지 랜덤 숫자 생성
	fmt.Println(answer)

	fmt.Printf("숫자 입력 :")
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	// guess, err := strconv.ParseInt(input, 10, 32) 더 디테일하게 조정이 가능함
	guess, err := strconv.Atoi(input) // string을 int로 변환하고 간단함
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(guess)

	if answer == guess {
		fmt.Println("정답입니다!")
	} else if answer > guess {
		fmt.Println("입력하신 값은 정답보다 작은 수 입니다. LOW")
	} else { // answer < guess
		fmt.Println("입력하신 값은 정답보다 큰 수 입니다. HIGH")
	}
}
