/*
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования. Длина пароля должна быть не менее 5 символов, он должен содержать только арабские цифры и буквы латинского алфавита. На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

Sample Input:

fdsghdfgjsdDD1
Sample Output:

Ok
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter a pwd: ")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	textRune := []rune(text)
	textRune1 := textRune[:len(textRune)-1]
	var wordPrint []string
	if len(textRune1) >= 5 {
		if strings.ContainsAny(text, "0123456789") {
			for _, v := range textRune1 {
				if (v >= '0' && v <= '9') || (v >= 'A' && v <= 'Z' || v >= 'a' && v <= 'z') {
					wordPrint = append(wordPrint, string(v))
				} else {
					fmt.Println("Wrong not-latin password")
					break
				}
			}
		} else {
			fmt.Println("Password without numbers")
		}
	} else {
		fmt.Println("Password's length less 5 symbols")
	}
	if len(textRune1) == len(wordPrint) { //не смог по другому вывести ОК
		fmt.Println("Ok")
	}
}
