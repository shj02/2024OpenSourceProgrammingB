package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // seed로 고정시킴(39를 하면 6만 나옴), 유닉스 시간
	answer := rand.Intn(6) + 1   // dice 1~6까지 랜덤 숫자 생성
	fmt.Println(answer)
}
