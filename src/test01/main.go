package main

import "fmt"

func main() {
	a := []byte{97, 98}
	b := "你好"
	b1 := []rune(b)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(b1)
}
