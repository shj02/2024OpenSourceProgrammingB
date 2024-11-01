package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("정수 입력 :")
	in := bufio.NewReader(os.Stdin)
	number, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	number = strings.TrimSpace(number)
	n, _ := strconv.Atoi(number)
	if err != nil {
		fmt.Println(err)
	}

	var isPrime bool = true
	if n <= 1 {
		isPrime = false
	} else {
		i := 2
		for i < n {
			if n%i == 0 {
				isPrime = false
				break // 1과 자기자신을 제외한 첫 번째 약수가 발견되면 반복문 종료
			}
			fmt.Printf("%d ", i) // 반복 횟수 확인용 코드
			i++
		}
	}

	if isPrime {
		fmt.Printf("%d는(은) 소수입니다", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다", n)
	}
}
