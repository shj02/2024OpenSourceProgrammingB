package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool {
	// var isPrime bool = true
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		i := 3
		for i*i <= n {
			if n%i == 0 {
				// isPrime = false
				// break
				return false
			}
			fmt.Printf("%d ", i)
			i = i + 2
		}
	}
	return true
}

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

	if isPrime(n) {
		fmt.Printf("%d는(은) 소수입니다", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다", n)
	}
}
