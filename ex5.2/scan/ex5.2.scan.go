package main

import "fmt"

func main() {
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	// & : 메모리 주소값, &a : a의 메모리 주소
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}
}
