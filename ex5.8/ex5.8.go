package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// ":=" : var stdin = bufio.NewReader(os.Stdin), 선언대입문
	stdin := bufio.NewReader(os.Stdin) //stdin : standard input

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b) // n, err : Scanln의 출력값, n은 결과값의 갯수
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n') // err를 발생시킨 stdin의 값을 메모리에서 비운다. '\n'이 나올 때까지 읽어오기
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b) // n, err : Scanln의 출력값
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n') // err를 발생시킨 stdin의 값을 메모리에서 비운다. '\n'이 나올 때까지 읽어오기
	} else {
		fmt.Println(n, a, b)
	}
}
