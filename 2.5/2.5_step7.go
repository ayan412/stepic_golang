package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Print("Enter a text: ")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	textRune := []rune(text)
	if unicode.IsUpper(textRune[0]) && (textRune[len(textRune)-2] == '.') {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
