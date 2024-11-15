package main

import (
	"fmt"
	"time"
)

func main() {
	// var dates [3]time.Time
	// dates[0] = time.Unix(0, 0)
	// dates[2] = time.Unix(1708012346, 0)
	// fmt.Println(dates[0], dates[1], dates[2]) // unix time, zero value, 2024-02-16 ...

	// var dates [3]time.Time = [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1708012346, 0)}
	// fmt.Println(dates[0], dates[1], dates[2])

	// dates := [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1708012346, 0)}
	// fmt.Println(dates[0], dates[1], dates[2])

	// dates := [3]time.Time{
	// 	time.Unix(0, 0),
	// 	time.Unix(1, 0),
	// 	time.Unix(1708012346, 0), // need comma
	// }
	// fmt.Println(dates[0], dates[1], dates[2]) // unix time, zero value, 2024-02-16 ...

	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(1708012346, 0)} // 중괄호랑 이어져 있으면 콤마 안 써도 됨!
	// fmt.Println(dates[0], dates[1], dates[2]) // unix time, zero value, 2024-02-16 ...
	// fmt.Println(dates)                        // array
	// fmt.Printf("%#v\n", dates)                // array literal

	// for i, date := range dates { // = python 'for in'
	// 	fmt.Println(i, date)
	// }

	for _, date := range dates { // = python 'for in', i를 _로 바꾸어서 인덱스는 출력 안되게 함
		fmt.Println(date)
	}
}
