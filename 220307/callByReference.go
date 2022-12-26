package main

import "fmt"

func inc(i *int) {
	*i = *i + 1
	fmt.Println(*i)
}

func main() {
	i := 10
	inc(&i)
	fmt.Println(i)
}
