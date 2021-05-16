package main

import "fmt"

func main() {

	hash := map[string]string{
		"aaaa": "4a",
		"bbbb": "4b",
	}
	if value, ok := hash["aaaa"]; ok {
		fmt.Println(value) //4a
	} else {
		fmt.Println("Not found!")
	}

	if value, ok := hash["aaa"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("Not found!") //Not found!
	}
}
