package main

import "fmt"

func main() {
	str := make([]int, 5, 8)
	fmt.Println(len(str), cap(str), str) // 5 8 [0 0 0 0 0]
}
