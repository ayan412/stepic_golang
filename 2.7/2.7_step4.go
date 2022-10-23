/*
Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.

Входные данные

Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и строка содержит только арабские цифры.

Выходные данные

Выведите максимальную цифру, которая встречается во введенной строке.

Sample Input:

1112221112
Sample Output:

2
*/
package main

import (
	"fmt"
)

func main() {
	var i string
	fmt.Scan(&i)
	byteSlyce := []byte(i)
	mx := byteSlyce[0]
	for i := 1; i <= len(byteSlyce)-1; i++ {
		if byteSlyce[i] >= mx {
			mx = byteSlyce[i]
		}
	}
	fmt.Print(string(mx))
}
