package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// около миши молоко
func main() {

	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	//fmt.Print(strings.ReplaceAll(str, " ", "")) эту констр-ю использовать в создании срезе рун
	strRune := []rune(strings.ReplaceAll(str, " ", ""))
	//fmt.Println("в срезе рун без пробелов:", strRune)
	a := len(strRune) - 2 // минус 2 т.к. индексация с нуля - число длины испол-ся в кач-ве индекса
	//fmt.Println(a)
	var sum int
	for i := 0; i <= a/2; i++ {
		if strRune[i] == strRune[a-i] {
			sum += i
			continue
		} else {
			fmt.Println("Нет")
			break
		}
	}
	if sum != 0 {
		fmt.Println("Палиндром")
	}
}
