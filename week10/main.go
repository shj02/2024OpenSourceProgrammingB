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
	} else if n == 2 {
		isPrime = true
	} else if n%2 == 0 {
		isPrime = false
	} else {
		i := 3
		// for i <= int(math.Sqrt(float64(n))) {
		for i*i <= n {
			if n%i == 0 {
				isPrime = false
				break
			}
			fmt.Printf("%d ", i)
			i = i + 2
		}
	}

	if isPrime {
		fmt.Printf("%d는(은) 소수입니다", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다", n)
	}
}
