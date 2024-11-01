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

	count := 0
	i := 1
	for i <= n {
		if n%i == 0 {
			count = count + 1 // count++
		}
		i++
	}
	if count == 2 {
		fmt.Printf("%d는(은) 소수입니다", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다", n)
	}
}
