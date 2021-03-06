/*
Дано трехзначное число. Найдите сумму его цифр.

Формат входных данных
На вход дается трехзначное число.

Формат выходных данных
Выведите одно целое число - сумму цифр введенного числа.

Sample Input:

745
Sample Output:

16
*/
package main

import "fmt"

func main() {
	var num, sum int
	fmt.Scan(&num)

	for num > 0 {
		sum += num % 10
		num = num / 10
	}
	fmt.Println(sum)

}
