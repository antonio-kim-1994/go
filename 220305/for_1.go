package main

import "fmt"

func main() {
	sum := 0
	// 초기화 구문; 조건식; 후속 작업 정의
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
