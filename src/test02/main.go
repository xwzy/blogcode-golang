package main

import "fmt"

func main() {
	str := make([]int, 5, 8)
	fmt.Println(len(str), cap(str), str) // 5 8 [0 0 0 0 0]

	str2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	str3 := str2[5:8]
	fmt.Println(len(str2), cap(str2)) //15 15
	fmt.Println(len(str3), cap(str3)) //3 10
	fmt.Println(str3)                 //[6 7 8]
}
