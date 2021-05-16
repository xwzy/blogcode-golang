package main

import "fmt"

func main() {
	a := []byte{97, 98}
	b := "你好"
	b1 := []rune(b)

	fmt.Println(a)  //[97 98]
	fmt.Println(b)  //你好
	fmt.Println(b1) //[20320 22909]
}
