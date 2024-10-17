package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Printf("오늘은 %d년 %d월 %d일\n", now.Year(), int(now.Month()), now.Day())
	fmt.Printf("지금은 %d시 %d분 %d초", now.Hour(), now.Minute(), now.Second())

	var army string = "우리는 !가와 !민에 충성을 다하는 대한민! 육!이다."
	armyFixed := strings.NewReplacer("!", "국")
	fmt.Println(army)
	fmt.Println(armyFixed.Replace(army))
}
