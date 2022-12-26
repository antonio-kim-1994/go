package main

import "fmt"

var g int = 10 // 패키지 전역변수

func main() {
	var m int = 20

	{
		var s int = 50
		fmt.Println(m, s, g)
	}

	// m = s + 20 => s가 변수의 범위를 벗어남
}
