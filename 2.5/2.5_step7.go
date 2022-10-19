/*
На вход подается строка. Нужно определить, является ли она правильной или нет. Правильная строка начинается с заглавной буквы и заканчивается точкой. Если строка правильная - вывести Right иначе - вывести Wrong

Маленькая подсказка: fmt.Scan() считывает строку до первого пробела, вы можете считать полностью строку за раз с помощью bufio:

text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

Sample Input:

Быть или не быть.
Sample Output:

Right
*/

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
