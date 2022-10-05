package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	strRune := []rune(str)
	for index, value := range strRune {
		if index%2 != 0 {
			fmt.Print(string(value))
		}

	}
	//fmt.Println(str)
}
