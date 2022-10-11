package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	//strRune := []rune(str)
	for i := 0; i < len(str); i++ {
		if strings.Count(str, string(str[i])) == 1 {
			fmt.Print(string(str[i]))
		}
		//fmt.Println(strings.Count(str, string(str[i])))
	}
}
