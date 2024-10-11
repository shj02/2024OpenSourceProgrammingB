package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Printf("오늘은 %d년 %d월 %d일\n", now.Year(), int(now.Month()), now.Day())
	fmt.Printf("지금은 %d시 %d분 %d초", now.Hour(), now.Minute(), now.Second())
}
