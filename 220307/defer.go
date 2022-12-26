package main

import "fmt"

func main() {
	f1()
}

func f1() {
	fmt.Println("f1 - start")
	// 함수가 종료되기 전까지 특정 구문의 실행을 지연시켰다가, 함수가 종료되기 직전에 구문 수행
	// 보통 특정 리소스의 사용을 해제하는 코드를 defer 구문으로 작성
	defer f2()
	fmt.Println("f1 - end")
}

func f2() {
	fmt.Println("f2 - deffered")
}
